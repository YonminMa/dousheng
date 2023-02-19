// Code generated by hertz generator.

package feed

import (
	"context"
	"dousheng/biz/dal/mysql"
	feed "dousheng/biz/model/feed"
	"dousheng/biz/mw"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sync"
	"time"
)

// Feed 获取视频 feed 流
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 设置 userId
	var userId int64
	if req.IsSetToken() {
		// 获取 claims 中的 payloads
		// 具体内容在 jwt mw 中进行设置
		claims, err := mw.JwtMiddleware.GetClaimsFromJWT(ctx, c)
		if err != nil {
			return
		}
		userId = int64(claims[mw.IdentityKey].(float64))
	} else {
		userId = -1
	}

	// 设置 latestTime
	var lastTime int64
	if req.IsSetLatestTime() {
		lastTime = req.GetLatestTime()
	} else {
		lastTime = time.Now().Unix()
	}

	// 获取视频数据
	videoRaws, err := mysql.QueryVideoByLatestTime(ctx, lastTime, userId)
	if err != nil {
		return
	}
	// 保存视频 id
	videoIds := make([]int64, 0)
	for _, videoRaw := range videoRaws {
		videoIds = append(videoIds, int64(videoRaw.ID))
	}
	// 保存发布视频用户 id
	userIds := make([]int64, 0)
	for _, videoRaw := range videoRaws {
		userIds = append(userIds, videoRaw.UserId)
	}

	users, err := mysql.QueryUserByIds(ctx, userIds)
	if err != nil {
		return
	}
	userMap := make(map[int64]*mysql.UserRaw)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	// 多协程执行 SQL
	var favoriteMap map[int64]*mysql.FavoriteRaw
	var relationMap map[int64]*mysql.RelationRaw
	var wg sync.WaitGroup
	wg.Add(2)
	//获取当前用户点赞信息
	go func() {
		defer wg.Done()
		favoriteMap, err = mysql.QueryFavoriteByVideoIds(ctx, userId, videoIds)
		if err != nil {
			return
		}
	}()
	//获取当前用户关注信息
	go func() {
		defer wg.Done()
		relationMap, err = mysql.QueryRelationByIds(ctx, userId, userIds)
		if err != nil {
			return
		}
	}()
	wg.Wait()

	// 组装 videoList
	isFavorite := false
	isFollow := false
	videoList := make([]*feed.Video, 0)
	for _, videoRaw := range videoRaws {
		// 获取发布当前视频用户
		videoUser, ok := userMap[videoRaw.UserId]
		if !ok {
			videoUser = &mysql.UserRaw{
				Name:          "用户已注销",
				FollowCount:   0,
				FollowerCount: 0,
			}
			videoUser.ID = 0
		}

		if userId != -1 {
			_, ok = favoriteMap[int64(videoRaw.ID)]
			if ok {
				isFavorite = true
			} else {
				isFavorite = false
			}

			_, ok = relationMap[videoRaw.UserId]
			if ok {
				isFollow = true
			} else {
				isFollow = false
			}
		}

		videoList = append(videoList, &feed.Video{
			ID: int64(videoRaw.ID),
			Author: &feed.User{
				ID:              int64(videoUser.ID),
				Name:            videoUser.Name,
				FollowCount:     videoUser.FollowerCount,
				FollowerCount:   videoUser.FollowCount,
				IsFollow:        isFollow,
				Avatar:          videoUser.Avatar,
				BackgroundImage: videoUser.BackgroundImage,
				Signature:       videoUser.Signature,
				TotalFavorited:  videoUser.TotalFavorited,
				WorkCount:       videoUser.WorkCount,
				FavoriteCount:   videoUser.FavoriteCount,
			},
			PlayURL:       videoRaw.PlayUrl,
			CoverURL:      videoRaw.CoverUrl,
			FavoriteCount: videoRaw.FavoriteCount,
			CommentCount:  videoRaw.CommentCount,
			IsFavorite:    isFavorite,
			Title:         videoRaw.Title,
		})
	}

	resp := feed.FeedResponse{
		StatusCode: feed.Code_Success,
		StatusMsg:  "Success",
		NextTime:   videoRaws[len(videoRaws)-1].CreatedAt.Unix(),
		VideoList:  videoList,
	}

	c.JSON(consts.StatusOK, resp)
}
