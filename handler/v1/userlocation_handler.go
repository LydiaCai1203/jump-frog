package v1

import (
	"framework/domain"
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
	group.POST("/upload", h.Upload)
}

func (h *UserLocationHandler) Upload(c echo.Context) error {
	var req domain.UserLocationRequest
	if err := c.Bind(&req); err != nil || req.Latitude == 0 || req.Longitude == 0 {
		return h.NewResponseWithError(c, "latitude and longitude are required", err)
	}
	// TODO: 调用 usecase 上传定位
	resp := domain.UserLocationResponse{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Status:    "uploaded",
	}
	return h.NewResponseWithData(c, resp)
}
