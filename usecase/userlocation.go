package usecase

import "framework/domain"

type UserLocationUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewUserLocationUsecase() *UserLocationUsecase {
	return &UserLocationUsecase{}
}

// 上传用户定位
func (u *UserLocationUsecase) Upload(req domain.UserLocationRequest) (domain.UserLocationResponse, error) {
	// TODO: 实现上传定位逻辑
	return domain.UserLocationResponse{}, nil
}
