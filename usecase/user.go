package usecase

import "framework/domain"

type UserUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

// 获取当前用户信息
func (u *UserUsecase) Info(req domain.UserMeInfoRequest) (domain.UserInfo, error) {
	// TODO: 实现获取当前用户信息逻辑
	return domain.UserInfo{}, nil
}

// 更新当前用户信息
func (u *UserUsecase) Update(req domain.UserMeUpdateRequest) error {
	// TODO: 实现更新用户信息逻辑
	return nil
}
