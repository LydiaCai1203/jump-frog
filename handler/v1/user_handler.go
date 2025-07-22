package v1

import (
	"framework/domain"
	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	BaseHandler
	UserUsecase *usecase.UserUsecase
}

func RegisterUserHandler(e *echo.Echo, userUsecase *usecase.UserUsecase) {
	h := &UserHandler{UserUsecase: userUsecase}
	group := e.Group("/api/v1/users")
	group.POST("/me", h.Me)
}

func (h *UserHandler) Me(c echo.Context) error {
	// 假设获取用户信息
	user := &domain.User{
		ID:       1,
		Username: "demo",
		Password: "",
	}
	return h.NewResponseWithData(c, user)
}
