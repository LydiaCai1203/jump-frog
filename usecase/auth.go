package usecase

import "framework/domain"

type AuthUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}

// 用户注册
func (u *AuthUsecase) Register(req domain.RegisterRequest) (id string, err error) {
	// TODO: 实现注册逻辑
	return "mock_id", nil
}

// 用户登录
func (u *AuthUsecase) Login(req domain.LoginRequest) (token string, err error) {
	// TODO: 实现登录逻辑
	return "mock_token", nil
}
