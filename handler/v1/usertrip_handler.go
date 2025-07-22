package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type CreateUserTripRequest struct {
	RouteID   string `json:"route_id"`
	StartDate string `json:"start_date"`
}

type UserTripHandler struct {
	UserTripUsecase *usecase.UserTripUsecase
}

func RegisterUserTripHandler(e *echo.Echo, userTripUsecase *usecase.UserTripUsecase) {
	h := &UserTripHandler{UserTripUsecase: userTripUsecase}
	group := e.Group("/api/v1/user-trips")
	group.GET("", h.ListUserTrips)
	group.POST("", h.CreateUserTrip)
	group.GET("/:tripId", h.GetUserTrip)
}

func (h *UserTripHandler) ListUserTrips(c echo.Context) error {
	// TODO: 获取我的所有行程
	return c.JSON(http.StatusOK, map[string]string{"message": "行程列表"})
}

func (h *UserTripHandler) CreateUserTrip(c echo.Context) error {
	// TODO: 创建新行程
	return c.JSON(http.StatusCreated, map[string]string{"message": "创建成功"})
}

func (h *UserTripHandler) GetUserTrip(c echo.Context) error {
	// TODO: 获取单个行程详情
	return c.JSON(http.StatusOK, map[string]string{"message": "行程详情"})
}
