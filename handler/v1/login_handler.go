package v1

import (
	"time"

	"framework/domain"
	"framework/repo"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	BaseHandler
	UserRepo *repo.UserRepo
}

func NewLoginHandler(e *echo.Echo, userRepo *repo.UserRepo) {
	h := &LoginHandler{UserRepo: userRepo}
	group := e.Group("/api/v1")
	group.POST("/login", h.Login)
}

func (h *LoginHandler) Login(c echo.Context) error {
	var req domain.LoginRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	if req.Username == "" || req.Password == "" {
		return h.NewResponseWithError(c, "username and password are required", nil)
	}
	user, err := h.UserRepo.FindByUsername(req.Username)
	if err != nil || user.Password != req.Password {
		return h.NewResponseWithError(c, "invalid credentials", err)
	}
	token, err := generateJWT(user.Username)
	if err != nil {
		return h.NewResponseWithError(c, "token error", err)
	}
	return h.NewResponseWithData(c, &domain.LoginResponse{Token: token})
}

func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte("your_secret_key"))
}
