package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type UserLocationRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type UserLocationHandler struct {
	UserLocationUsecase *usecase.UserLocationUsecase
}

func RegisterUserLocationHandler(e *echo.Echo, userLocationUsecase *usecase.UserLocationUsecase) {
	h := &UserLocationHandler{UserLocationUsecase: userLocationUsecase}
	group := e.Group("/api/v1/user-locations")
	group.POST("", h.UploadLocation)
}

func (h *UserLocationHandler) UploadLocation(c echo.Context) error {
	// TODO: 上传用户定位
	return c.JSON(http.StatusCreated, map[string]string{"message": "上传成功"})
}
