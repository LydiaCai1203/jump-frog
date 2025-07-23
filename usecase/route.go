package usecase

import (
	"framework/domain"
	"framework/repo"
)

type RouteUsecase struct {
	routeRepo *repo.RouteRepo
}

func NewRouteUsecase(routeRepo *repo.RouteRepo) *RouteUsecase {
	return &RouteUsecase{
		routeRepo: routeRepo,
	}
}

// RouteList 获取路线列表
func (u *RouteUsecase) RouteList(req domain.RouteListRequest) (domain.RouteListResponse, error) {
	total, routes, err := u.routeRepo.RouteList(req.Page, req.Limit)
	if err != nil {
		return domain.RouteListResponse{}, err
	}

	var list []domain.RouteDetail
	for _, route := range routes {
		list = append(list, domain.RouteDetail{
			ID:          route.ID,
			Name:        route.Name,
			Description: route.Description,
		})
	}

	return domain.RouteListResponse{
		List:  list,
		Total: int(total),
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

// RouteDetail 获取路线详情
func (u *RouteUsecase) RouteDetail(req domain.RouteDetailRequestV2) (domain.RouteDetail, error) {
	route, err := u.routeRepo.RouteDetail(req.RouteID)
	if err != nil {
		return domain.RouteDetail{}, err
	}

	return domain.RouteDetail{
		ID:          route.ID,
		Name:        route.Name,
		Description: route.Description,
	}, nil
}
