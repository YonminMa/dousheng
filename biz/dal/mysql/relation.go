package mysql

import (
	"context"
	"dousheng/pkg/constants"
	"gorm.io/gorm"
)

type RelationRaw struct {
	gorm.Model
	UserId   string
	ToUserId string
}

func (RelationRaw) TableName() string {
	return constants.RelationTableName
}

func CheckIsFollow(ctx context.Context, userId, toUserId string) bool {
	var count int64
	err := DB.WithContext(ctx).
		Model(&RelationRaw{}).
		Where("user_id = ?", userId).
		Where("to_user_id = ?", toUserId).
		Count(&count).Error
	if err != nil {
		panic(err)
	}
	return count != 0
}
