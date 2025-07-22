package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	AuthUsecase *usecase.AuthUsecase
}

func RegisterAuthHandler(e *echo.Echo, authUsecase *usecase.AuthUsecase) {
	h := &AuthHandler{AuthUsecase: authUsecase}
	group := e.Group("/api/v1/auth")
	group.POST("/register", h.Register)
	group.POST("/login", h.Login)
}

func (h *AuthHandler) Register(c echo.Context) error {
	// TODO: 参数校验、调用 usecase、返回响应
	return c.JSON(http.StatusCreated, map[string]string{"message": "注册成功"})
}

func (h *AuthHandler) Login(c echo.Context) error {
	// TODO: 参数校验、调用 usecase、返回响应
	return c.JSON(http.StatusOK, map[string]string{"token": "mock_token"})
}
