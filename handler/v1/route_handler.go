package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type RouteHandler struct {
	RouteUsecase *usecase.RouteUsecase
}

func RegisterRouteHandler(e *echo.Echo, routeUsecase *usecase.RouteUsecase) {
	h := &RouteHandler{RouteUsecase: routeUsecase}
	group := e.Group("/api/v1/routes")
	group.GET("", h.ListRoutes)
	group.GET("/:routeId", h.GetRoute)
}

func (h *RouteHandler) ListRoutes(c echo.Context) error {
	// TODO: 获取所有热门路线
	return c.JSON(http.StatusOK, map[string]string{"message": "路线列表"})
}

func (h *RouteHandler) GetRoute(c echo.Context) error {
	// TODO: 获取单条路线详情
	return c.JSON(http.StatusOK, map[string]string{"message": "路线详情"})
}
