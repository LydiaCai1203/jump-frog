package repo

import "framework/domain"

type UserRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
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
