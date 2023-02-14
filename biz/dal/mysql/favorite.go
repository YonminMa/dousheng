package mysql

import (
	"context"
	"dousheng/pkg/constants"
	"gorm.io/gorm"
)

type FavoriteRaw struct {
	gorm.Model
	UserId  int64
	VideoId int64
}

func (FavoriteRaw) TableName() string {
	return constants.FavoriteTableName
}

// QueryFavoriteByVideoIds 根据当前用户id和视频id获取点赞信息
func QueryFavoriteByVideoIds(ctx context.Context, currentId int64, videoIds []int64) (map[int64]*FavoriteRaw, error) {
	var favorites []*FavoriteRaw
	err := DB.WithContext(ctx).
		Where("user_id = ? AND video_id IN ?", currentId, videoIds).
		Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	favoriteMap := make(map[int64]*FavoriteRaw)
	for _, favorite := range favorites {
		favoriteMap[favorite.VideoId] = favorite
	}
	return favoriteMap, nil
}
