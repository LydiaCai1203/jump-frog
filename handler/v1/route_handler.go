package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type RouteRequest struct {
	RouteID   string `json:"route_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
}

type RouteHandler struct {
	RouteUsecase *usecase.RouteUsecase
}

func RegisterRouteHandler(e *echo.Echo, routeUsecase *usecase.RouteUsecase) {
	h := &RouteHandler{RouteUsecase: routeUsecase}
	group := e.Group("/api/v1/routes")
	group.POST("", h.PostRoute)
	group.POST("/detail", h.PostRoute)
}

func (h *RouteHandler) PostRoute(c echo.Context) error {
	var req RouteRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	// TODO: 根据 req 字段判断是列表还是详情，调用 usecase
	return c.JSON(http.StatusOK, map[string]string{"message": "路线列表/详情"})
}
