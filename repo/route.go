package repo

type RouteRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewRouteRepo() *RouteRepo {
	return &RouteRepo{}
}

func (r *RouteRepo) ListRoutes() (interface{}, error) {
	// TODO: 实现获取所有热门路线
	return nil, nil
}

func (r *RouteRepo) GetRoute(routeId string) (interface{}, error) {
	// TODO: 实现获取单条路线详情
	return nil, nil
}
