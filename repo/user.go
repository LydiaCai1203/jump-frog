package repo

import (
	"framework/domain"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// 获取当前用户信息
func (r *UserRepo) Info(req domain.UserMeInfoRequest) (domain.UserInfo, error) {
	// TODO: 实现获取当前用户信息
	return domain.UserInfo{}, nil
}

// 更新当前用户信息
func (r *UserRepo) Update(req domain.UserMeUpdateRequest) error {
	// TODO: 实现更新用户信息
	return nil
}
