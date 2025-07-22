package v1

import (
	"framework/domain"
	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type RouteHandler struct {
	BaseHandler
	RouteUsecase *usecase.RouteUsecase
}

func RegisterRouteHandler(e *echo.Echo, routeUsecase *usecase.RouteUsecase) {
	h := &RouteHandler{RouteUsecase: routeUsecase}
	group := e.Group("/api/v1/routes")
	group.POST("", h.List)
	group.POST("/detail", h.Detail)
}

func (h *RouteHandler) List(c echo.Context) error {
	// 假设返回路线列表
	routes := []domain.RouteItem{
		{
			RouteID:   "1",
			Name:      "A-B",
			Days:      3,
			StartCity: "A",
			EndCity:   "B",
			IsHot:     true,
		},
	}
	return h.NewResponseWithData(c, &domain.RouteListData{Routes: routes})
}

func (h *RouteHandler) Detail(c echo.Context) error {
	var req struct {
		RouteID string `json:"route_id"`
	}
	if err := c.Bind(&req); err != nil || req.RouteID == "" {
		return h.NewResponseWithError(c, "route_id is required", err)
	}
	// 假设返回路线详情
	detail := &domain.RouteItem{
		RouteID:   req.RouteID,
		Name:      "A-B",
		Days:      3,
		StartCity: "A",
		EndCity:   "B",
		IsHot:     true,
	}
	return h.NewResponseWithData(c, detail)
}
