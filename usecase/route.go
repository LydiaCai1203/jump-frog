package usecase

type RouteUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewRouteUsecase() *RouteUsecase {
	return &RouteUsecase{}
}

func (u *RouteUsecase) ListRoutes() (interface{}, error) {
	// TODO: 实现获取所有热门路线逻辑
	return nil, nil
}

func (u *RouteUsecase) GetRoute(routeId string) (interface{}, error) {
	// TODO: 实现获取单条路线详情逻辑
	return nil, nil
}
