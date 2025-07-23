package repo

import "framework/domain"

type RouteRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewRouteRepo() *RouteRepo {
	return &RouteRepo{}
}

// 获取路线列表
func (r *RouteRepo) List(req domain.RouteListRequest) (domain.RouteListResponse, error) {
	// TODO: 实现获取路线列表
	return domain.RouteListResponse{}, nil
}

// 获取路线详情
func (r *RouteRepo) Detail(req domain.RouteDetailRequestV2) (domain.RouteDetail, error) {
	// TODO: 实现获取路线详情
	return domain.RouteDetail{}, nil
}
