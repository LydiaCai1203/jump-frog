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
	group.POST("/list", h.RouteList)
	group.POST("/:routeId", h.RouteDetail)
}

// RouteList 获取路线列表
func (h *RouteHandler) RouteList(c echo.Context) error {
	var req domain.RouteListRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}

	resp, err := h.RouteUsecase.RouteList(req)
	if err != nil {
		return h.NewResponseWithError(c, "failed to get route list", err)
	}

	return h.NewResponseWithData(c, resp)
}

// RouteDetail 获取路线详情
func (h *RouteHandler) RouteDetail(c echo.Context) error {
	var req domain.RouteDetailRequestV2
	if err := c.Bind(&req); err != nil || req.RouteID == "" {
		return h.NewResponseWithError(c, "route_id is required", err)
	}

	resp, err := h.RouteUsecase.RouteDetail(req)
	if err != nil {
		return h.NewResponseWithError(c, "failed to get route detail", err)
	}

	return h.NewResponseWithData(c, resp)
}
