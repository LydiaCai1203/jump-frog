package repo

import "framework/domain"

type AuthRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}

// 用户注册
func (r *AuthRepo) Register(req domain.RegisterRequest) (id string, err error) {
	// TODO: 实现注册
	return "mock_id", nil
}

// 用户登录校验
func (r *AuthRepo) Login(req domain.LoginRequest) (token string, err error) {
	// TODO: 实现登录校验
	return "mock_token", nil
}
