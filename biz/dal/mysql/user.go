package mysql

import (
	"context"
	"dousheng/pkg/constants"
	"gorm.io/gorm"
)

type UserRaw struct {
	gorm.Model
	Name            string
	Password        string
	FollowCount     int64
	FollowerCount   int64
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64
	WorkCount       int64
	FavoriteCount   int64
}

func (UserRaw) TableName() string {
	return constants.UserTableName
}

// QueryUserByName 根据用户名获取用户信息
func QueryUserByName(ctx context.Context, username string) ([]*UserRaw, error) {
	var users []*UserRaw
	err := DB.WithContext(ctx).Where("name = ?", username).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// QueryUserById 根据 user_id 获取用户信息
func QueryUserById(ctx context.Context, userId int64) ([]*UserRaw, error) {
	var users []*UserRaw
	err := DB.WithContext(ctx).Where("id = ?", userId).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// QueryUserByIds 根据 user_ids 获取用户信息
func QueryUserByIds(ctx context.Context, userIds []int64) ([]*UserRaw, error) {
	var users []*UserRaw
	err := DB.WithContext(ctx).Where("id in (?)", userIds).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UserRegister 注册用户
func UserRegister(ctx context.Context, username string, password string) (int64, error) {
	userRaw := UserRaw{
		Name:          username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}

	err := DB.WithContext(ctx).Create(&userRaw).Error
	if err != nil {
		panic(err)
	}
	return int64(userRaw.ID), err
}

// CheckUser 查询用户名和密码是否正确
func CheckUser(username, password string) ([]*UserRaw, error) {
	res := make([]*UserRaw, 0)
	if err := DB.Where("name = ?", username).
		Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
