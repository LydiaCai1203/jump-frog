package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type UserMeRequest struct {
	Nickname  string `json:"nickname,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Bio       string `json:"bio,omitempty"`
}

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func RegisterUserHandler(e *echo.Echo, userUsecase *usecase.UserUsecase) {
	h := &UserHandler{UserUsecase: userUsecase}
	group := e.Group("/api/v1/users")
	group.POST("/me", h.PostMe)
}

func (h *UserHandler) PostMe(c echo.Context) error {
	var req UserMeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	// TODO: 根据 req 字段判断是获取还是更新，调用 usecase
	return c.JSON(http.StatusOK, map[string]string{"message": "用户信息/更新成功"})
}
