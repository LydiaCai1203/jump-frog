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
	group.POST("", h.PostFollow)
}

func (h *FollowHandler) PostFollow(c echo.Context) error {
	var req FollowRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	// TODO: 关注或取消关注
	return c.JSON(http.StatusOK, map[string]string{"message": "关注/取消关注成功"})
}
