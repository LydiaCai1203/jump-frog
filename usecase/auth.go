package usecase

import (
	"framework/domain"
	"framework/repo"
)

type AuthUsecase struct {
	authRepo *repo.AuthRepo
}

func NewAuthUsecase(authRepo *repo.AuthRepo) *AuthUsecase {
	return &AuthUsecase{authRepo: authRepo}
}

// 用户注册
func (u *AuthUsecase) Register(req domain.RegisterRequest) (id string, err error) {
	return u.authRepo.Register(req)
}

// 用户登录
func (u *AuthUsecase) Login(req domain.LoginRequest) (token string, err error) {
	return u.authRepo.Login(req)
}
