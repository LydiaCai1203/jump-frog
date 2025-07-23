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
	group.POST("/list", h.List)
	group.POST("/:routeId", h.Detail)
}

func (h *RouteHandler) List(c echo.Context) error {
	var req domain.RouteListRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	// TODO: 调用 usecase 获取数据
	resp := domain.RouteListResponse{
		List:  []domain.RouteDetail{},
		Total: 0,
		Page:  req.Page,
		Limit: req.Limit,
	}
	return h.NewResponseWithData(c, resp)
}

func (h *RouteHandler) Detail(c echo.Context) error {
	var req domain.RouteDetailRequestV2
	if err := c.Bind(&req); err != nil || req.RouteID == "" {
		return h.NewResponseWithError(c, "route_id is required", err)
	}
	// TODO: 调用 usecase 获取数据
	resp := domain.RouteDetail{}
	return h.NewResponseWithData(c, resp)
}
