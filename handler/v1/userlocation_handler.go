package v1

import (
	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type UserLocationHandler struct {
	BaseHandler
	UserLocationUsecase *usecase.UserLocationUsecase
}

func RegisterUserLocationHandler(e *echo.Echo, userLocationUsecase *usecase.UserLocationUsecase) {
	h := &UserLocationHandler{UserLocationUsecase: userLocationUsecase}
	group := e.Group("/api/v1/user-locations")
	group.POST("", h.Upload)
}

func (h *UserLocationHandler) Upload(c echo.Context) error {
	var req struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	if err := c.Bind(&req); err != nil || req.Latitude == 0 || req.Longitude == 0 {
		return h.NewResponseWithError(c, "latitude and longitude are required", err)
	}
	// 假设上传成功
	return h.NewResponseWithData(c, map[string]any{
		"latitude":  req.Latitude,
		"longitude": req.Longitude,
		"status":    "uploaded",
	})
}
