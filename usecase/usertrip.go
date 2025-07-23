package usecase

import "framework/domain"

type UserTripUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewUserTripUsecase() *UserTripUsecase {
	return &UserTripUsecase{}
}

// 创建新行程
func (u *UserTripUsecase) Choose(req domain.UserTripChooseRequest) error {
	// TODO: 实现创建新行程逻辑
	return nil
}

// 获取行程详情
func (u *UserTripUsecase) Detail(req domain.UserTripDetailRequestV2) (domain.TripDetail, error) {
	// TODO: 实现获取行程详情逻辑
	return domain.TripDetail{}, nil
}
