package v1

import (
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
	return nil
}

func (h *RouteHandler) Detail(c echo.Context) error {
	var req struct {
		RouteID string `json:"route_id"`
	}
	if err := c.Bind(&req); err != nil || req.RouteID == "" {
		return h.NewResponseWithError(c, "route_id is required", err)
	}
	return nil
}
