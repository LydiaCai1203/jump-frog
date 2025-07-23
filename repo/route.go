package repo

import (
	"framework/domain"

	"gorm.io/gorm"
)

type RouteRepo struct {
	db *gorm.DB
}

func NewRouteRepo(db *gorm.DB) *RouteRepo {
	return &RouteRepo{db: db}
}

// RouteList 获取路线列表
func (r *RouteRepo) RouteList(page int, limit int) (int64, []*domain.Route, error) {
	var total int64
	err := r.db.Model(&domain.Route{}).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	var routes []*domain.Route
	err = r.db.
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&routes).
		Order("created_at desc").
		Error
	if err != nil {
		return 0, nil, err
	}
	return total, routes, nil
}

// RouteDetail 获取路线详情
func (r *RouteRepo) RouteDetail(routeID string) (*domain.Route, error) {
	var route *domain.Route
	err := r.db.Where("id = ?", routeID).First(&route).Error
	if err != nil {
		return nil, err
	}
	return route, nil
}
