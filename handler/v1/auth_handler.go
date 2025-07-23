package v1

import (
	"fmt"

	"framework/domain"
	"framework/usecase"
	"framework/utils"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	BaseHandler
	AuthUsecase *usecase.AuthUsecase
}

// NewAuthHandler 创建 AuthHandler
func NewAuthHandler(e *echo.Echo, authUsecase *usecase.AuthUsecase) *AuthHandler {
	h := &AuthHandler{AuthUsecase: authUsecase}
	group := e.Group("/api/v1/auth")
	group.POST("/register", h.Register)
	group.POST("/login", h.Login)
	return h
}

// Register 注册
func (h *AuthHandler) Register(c echo.Context) error {
	var req domain.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}

	if req.Email == "" && req.Phone == "" {
		return h.NewResponseWithError(c, "email or phone is required", nil)
	}

	if req.Username == "" || req.Password == "" {
		return h.NewResponseWithError(c, "username and password are required", nil)
	}

	userID, err := h.AuthUsecase.Register(req)
	if err != nil {
		return h.NewResponseWithError(c, "register failed", err)
	}

	// generate jwt token
	userID = fmt.Sprintf("JUMP_%s", userID)
	token, err := utils.GenerateJwtToken(userID)
	if err != nil {
		return h.NewResponseWithError(c, "generate jwt token failed", err)
	}

	return h.NewResponseWithData(c, domain.RegisterResponse{
		ID:    userID,
		Token: token,
	})
}

// Login 登录
func (h *AuthHandler) Login(c echo.Context) error {
	var req domain.LoginRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	if req.Username == "" && req.Email == "" && req.Phone == "" {
		return h.NewResponseWithError(c, "username, email or phone is required", nil)
	}

	if req.Password == "" {
		return h.NewResponseWithError(c, "password is required", nil)
	}

	userID, err := h.AuthUsecase.Login(req)
	if err != nil {
		return h.NewResponseWithError(c, "login failed", err)
	}

	// generate jwt token
	userID = fmt.Sprintf("JUMP_%s", userID)
	token, err := utils.GenerateJwtToken(userID)
	if err != nil {
		return h.NewResponseWithError(c, "generate jwt token failed", err)
	}

	return h.NewResponseWithData(c, domain.LoginResponse{
		Token: token,
	})
}
