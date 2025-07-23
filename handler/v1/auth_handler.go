package v1

import (
	"framework/domain"
	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	BaseHandler
	AuthUsecase *usecase.AuthUsecase
}

func RegisterAuthHandler(e *echo.Echo, authUsecase *usecase.AuthUsecase) {
	h := &AuthHandler{AuthUsecase: authUsecase}
	group := e.Group("/api/v1/auth")
	group.POST("/register", h.Register)
	group.POST("/login", h.Login)
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req domain.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	if req.Username == "" || req.Password == "" {
		return h.NewResponseWithError(c, "username and password are required", nil)
	}
	// TODO: 调用 usecase 注册
	resp := struct {
		ID string `json:"id"`
	}{ID: "mock_id"}
	return h.NewResponseWithData(c, resp)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req domain.LoginRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	if req.Username == "" || req.Password == "" {
		return h.NewResponseWithError(c, "username and password are required", nil)
	}
	// TODO: 调用 usecase 登录
	resp := struct {
		Token string `json:"token"`
	}{Token: "mock_token"}
	return h.NewResponseWithData(c, resp)
}
