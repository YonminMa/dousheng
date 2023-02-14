package mysql

import (
	"context"
	"dousheng/pkg/constants"
)

// PublishVideoData 发布一条视频
func PublishVideoData(ctx context.Context, videoData *VideoRaw) error {
	if err := DB.WithContext(ctx).Create(&videoData).Error; err != nil {
		return err
	}
	return nil
}

// QueryVideoByUserId 通过 user_id获取视频列表
func QueryVideoByUserId(ctx context.Context, userId int64) ([]*VideoRaw, error) {
	var videos []*VideoRaw
	err := DB.WithContext(ctx).
		Table(constants.VideoTableName).
		Order("updated_at desc").
		Where("user_id = ?", userId).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
