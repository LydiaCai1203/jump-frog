package v1

import (
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
	// 假设返回行程列表
	trips := []map[string]any{
		{
			"trip_id":    "1",
			"route_id":   "1",
			"start_date": "2024-01-01",
			"status":     "ongoing",
		},
	}
	return h.NewResponseWithData(c, trips)
}

func (h *UserTripHandler) Detail(c echo.Context) error {
	var req struct {
		TripID string `json:"trip_id"`
	}
	if err := c.Bind(&req); err != nil || req.TripID == "" {
		return h.NewResponseWithError(c, "trip_id is required", err)
	}
	// 假设返回行程详情
	detail := map[string]any{"trip_id": req.TripID, "route_id": "1", "start_date": "2024-01-01", "status": "ongoing"}
	return h.NewResponseWithData(c, detail)
}
