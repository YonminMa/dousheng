package mysql

import (
	"context"
	"dousheng/pkg/constants"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

// VideoRaw Video 数据库格式
type VideoRaw struct {
	gorm.Model           // 默认添加 id, created_time, updated_time, deleted_time 字段
	UserId        int64  `gorm:"column:user_id;not null;index:idx_userid"`
	Title         string `gorm:"column:title;type:varchar(128);not null"`
	PlayUrl       string `gorm:"column:play_url;varchar(128);not null"`
	CoverUrl      string `gorm:"column:cover_url;varchar(128);not null"`
	FavoriteCount int64  `gorm:"column:favorite_count;default:0"`
	CommentCount  int64  `gorm:"column:comment_count;default:0"`
}

// TableName 映射表名
func (VideoRaw VideoRaw) TableName() string {
	return constants.VideoTableName
}

// QueryVideoByLatestTime 通过 latest update time 查询视频
// 剔除 user_id 为当前用户 id 的视频
func QueryVideoByLatestTime(ctx context.Context, latestTime, userId int64) ([]*VideoRaw, error) {
	var videos []*VideoRaw
	fmt.Println(time.Unix(latestTime, 0).UTC())
	t := time.Unix(latestTime, 0).UTC()
	err := DB.WithContext(ctx).
		Limit(30).
		Order("updated_at desc").
		Where("updated_at < ?", t).
		Where("user_id != ?", userId).
		Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

// FindVideoById 根据 id 查找视频
func FindVideoById(id int64) VideoRaw {
	var video VideoRaw
	err := DB.Where("id = ?", id).First(&video).Error
	if err != nil {
		panic(err)
	}
	return video
}

// FindAllVideos 查询所有视频
func FindAllVideos() []VideoRaw {
	var videos []VideoRaw
	err := DB.Find(&videos).Error
	if err != nil {
		panic(err)
	}
	return videos
}

// CreateVideo 添加一个新视频
func CreateVideo(video VideoRaw) {
	err := DB.Create(video).Error
	if err != nil {
		panic(err)
	}
}

// UpdateVideoById 更新一个视频
func UpdateVideoById(id int64) {
	err := DB.Model(&VideoRaw{}).Where("id = ?", id).Update("author = ?", "sadasd").Error
	if err != nil {
		//panic(err)
		log.Println("update video error", err)
	}
}

// DeleteVideoById 删除一个视频
func DeleteVideoById(id int64) {
	err := DB.Delete(&VideoRaw{}).Where("id = ?", id).Error
	if err != nil {
		panic(err)
	}
}
