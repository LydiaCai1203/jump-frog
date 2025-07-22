package v1

import (
	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type FollowHandler struct {
	BaseHandler
	FollowUsecase *usecase.FollowUsecase
}

func RegisterFollowHandler(e *echo.Echo, followUsecase *usecase.FollowUsecase) {
	h := &FollowHandler{FollowUsecase: followUsecase}
	group := e.Group("/api/v1/follows")
	group.POST("", h.Follow)
}

func (h *FollowHandler) Follow(c echo.Context) error {
	var req struct {
		FolloweeID string `json:"followee_id"`
	}
	if err := c.Bind(&req); err != nil || req.FolloweeID == "" {
		return h.NewResponseWithError(c, "followee_id is required", err)
	}
	// 假设关注成功
	return h.NewResponseWithData(c, map[string]any{
		"followee_id": req.FolloweeID,
		"status":      "followed",
	})
}
