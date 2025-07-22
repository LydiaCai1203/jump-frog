package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type UserTripRequest struct {
	RouteID   string `json:"route_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	TripID    string `json:"trip_id,omitempty"`
}

type UserTripHandler struct {
	UserTripUsecase *usecase.UserTripUsecase
}

func RegisterUserTripHandler(e *echo.Echo, userTripUsecase *usecase.UserTripUsecase) {
	h := &UserTripHandler{UserTripUsecase: userTripUsecase}
	group := e.Group("/api/v1/user-trips")
	group.POST("", h.PostUserTrip)
	group.POST("/detail", h.PostUserTrip)
}

func (h *UserTripHandler) PostUserTrip(c echo.Context) error {
	var req UserTripRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	// TODO: 根据 req 字段判断是列表、详情还是创建，调用 usecase
	return c.JSON(http.StatusOK, map[string]string{"message": "行程列表/详情/创建成功"})
}
