package usecase

import "framework/domain"

type RouteUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewRouteUsecase() *RouteUsecase {
	return &RouteUsecase{}
}

// 获取路线列表
func (u *RouteUsecase) List(req domain.RouteListRequest) (domain.RouteListResponse, error) {
	// TODO: 实现获取路线列表逻辑
	return domain.RouteListResponse{}, nil
}

// 获取路线详情
func (u *RouteUsecase) Detail(req domain.RouteDetailRequestV2) (domain.RouteDetail, error) {
	// TODO: 实现获取路线详情逻辑
	return domain.RouteDetail{}, nil
}
