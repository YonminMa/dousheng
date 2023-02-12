package mysql

import (
	"context"
	"dousheng/pkg/constants"
	"gorm.io/gorm"
)

type RelationRaw struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:fk_relation_user"`
	ToUserId int64 `gorm:"column:to_user_id;not null"`
}

func (RelationRaw) TableName() string {
	return constants.RelationTableName
}

// CreateRelation 创建关注
func CreateRelation(ctx context.Context, currentId int64, toUserId int64) error {
	relationRaw := &RelationRaw{
		UserId:   currentId,
		ToUserId: toUserId,
	}
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 增加关注者关注数量
		err := tx.Table(constants.UserTableName).Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			return err
		}
		// 增加被关注者粉丝数量
		err = tx.Table(constants.UserTableName).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			return err
		}
		err = tx.Table(constants.RelationTableName).Create(&relationRaw).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteRelation 取消关注
func DeleteRelation(ctx context.Context, currentId int64, toUserId int64) error {
	var relationRaw *RelationRaw
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 减少取消关注者关注数量
		err := tx.Table(constants.UserTableName).Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
		if err != nil {
			return err
		}
		// 减少被取消关注者粉丝数量
		err = tx.Table(constants.UserTableName).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
		if err != nil {
			return err
		}
		err = tx.Table(constants.RelationTableName).Where("user_id = ? AND to_user_id = ?", currentId, toUserId).Delete(&relationRaw).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// CheckIsFollow 查询是否关注
func CheckIsFollow(ctx context.Context, userId, toUserId int64) bool {
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

// QueryFollowById 根据 user_id 查询该用户关注的用户 id
func QueryFollowById(ctx context.Context, userId int64) ([]*int64, error) {
	var follows []*int64
	err := DB.WithContext(ctx).
		Table(constants.RelationTableName).
		Select("to_user_id").
		Where("user_id = ?", userId).
		Find(&follows).Error
	if err != nil {
		return nil, err
	}
	return follows, nil
}

// QueryFollowerById 根据 to_user_id 查询该用户的粉丝 id
func QueryFollowerById(ctx context.Context, toUserId int64) ([]*int64, error) {
	var followers []*int64
	err := DB.WithContext(ctx).
		Table(constants.RelationTableName).
		Select("user_id").
		Where("to_user_id = ?", toUserId).
		Find(&followers).Error
	if err != nil {
		return nil, err
	}
	return followers, nil
}
