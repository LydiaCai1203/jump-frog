package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type FollowRequest struct {
	FolloweeID string `json:"followee_id"`
}

type FollowHandler struct {
	FollowUsecase *usecase.FollowUsecase
}

func RegisterFollowHandler(e *echo.Echo, followUsecase *usecase.FollowUsecase) {
	h := &FollowHandler{FollowUsecase: followUsecase}
	group := e.Group("/api/v1/follows")
	group.POST("", h.Follow)
	group.DELETE("", h.Unfollow)
}

func (h *FollowHandler) Follow(c echo.Context) error {
	// TODO: 关注用户
	return c.JSON(http.StatusCreated, map[string]string{"message": "关注成功"})
}

func (h *FollowHandler) Unfollow(c echo.Context) error {
	// TODO: 取消关注
	return c.JSON(http.StatusOK, map[string]string{"message": "取消关注成功"})
}
