package repo

import "framework/domain"

type UserTripRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewUserTripRepo() *UserTripRepo {
	return &UserTripRepo{}
}

// 创建新行程
func (r *UserTripRepo) Choose(req domain.UserTripChooseRequest) error {
	// TODO: 实现创建新行程
	return nil
}

// 获取行程详情
func (r *UserTripRepo) Detail(req domain.UserTripDetailRequestV2) (domain.TripDetail, error) {
	// TODO: 实现获取行程详情
	return domain.TripDetail{}, nil
}
