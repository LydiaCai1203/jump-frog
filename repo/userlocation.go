package repo

import "framework/domain"

type UserLocationRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewUserLocationRepo() *UserLocationRepo {
	return &UserLocationRepo{}
}

// 上传用户定位
func (r *UserLocationRepo) Upload(req domain.UserLocationRequest) (domain.UserLocationResponse, error) {
	// TODO: 实现上传定位
	return domain.UserLocationResponse{}, nil
}
