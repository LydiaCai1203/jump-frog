package v1

import (
	"errors"

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
		return errors.New("invalid request")
	}
	if req.Username == "" || req.Password == "" {
		return errors.New("username and password are required")
	}
	return h.NewResponseWithData(c, &domain.RegisterResponse{
		ID:       1,
		Username: req.Username,
		Nickname: req.Nickname,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req domain.LoginRequest
	if err := c.Bind(&req); err != nil {
		return errors.New("invalid request")
	}
	if req.Username == "" || req.Password == "" {
		return errors.New("username and password are required")
	}
	return h.NewResponseWithData(c, &domain.LoginResponse{
		Token: "mock_token",
	})
}
