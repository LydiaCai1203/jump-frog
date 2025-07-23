package v1

import (
	"framework/domain"
	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type UserTripHandler struct {
	BaseHandler
	UserTripUsecase *usecase.UserTripUsecase
}

func RegisterUserTripHandler(e *echo.Echo, userTripUsecase *usecase.UserTripUsecase) {
	h := &UserTripHandler{UserTripUsecase: userTripUsecase}
	group := e.Group("/api/v1/user-trips")
	group.POST("", h.List)
	group.POST("/detail", h.Detail)
}

func (h *UserTripHandler) List(c echo.Context) error {
	var req domain.UserTripChooseRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	// TODO: 调用 usecase 获取行程列表
	resp := []domain.TripDetail{}
	return h.NewResponseWithData(c, resp)
}

func (h *UserTripHandler) Detail(c echo.Context) error {
	var req domain.UserTripDetailRequestV2
	if err := c.Bind(&req); err != nil || req.TripID == "" {
		return h.NewResponseWithError(c, "trip_id is required", err)
	}
	// TODO: 调用 usecase 获取行程详情
	resp := domain.TripDetail{}
	return h.NewResponseWithData(c, resp)
}
