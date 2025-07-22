package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type UpdateUserRequest struct {
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
}

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func RegisterUserHandler(e *echo.Echo, userUsecase *usecase.UserUsecase) {
	h := &UserHandler{UserUsecase: userUsecase}
	group := e.Group("/api/v1/users")
	group.GET("/me", h.GetMe)
	group.PUT("/me", h.UpdateMe)
}

func (h *UserHandler) GetMe(c echo.Context) error {
	// TODO: 获取当前用户信息
	return c.JSON(http.StatusOK, map[string]string{"message": "用户信息"})
}

func (h *UserHandler) UpdateMe(c echo.Context) error {
	// TODO: 更新当前用户信息
	return c.JSON(http.StatusOK, map[string]string{"message": "更新成功"})
}
