package v1

import (
	"framework/domain"
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
	var req domain.FollowRequest
	if err := c.Bind(&req); err != nil || req.FolloweeID == "" {
		return h.NewResponseWithError(c, "followee_id is required", err)
	}
	// TODO: 调用 usecase 关注
	resp := domain.FollowResponse{
		FolloweeID: req.FolloweeID,
		Status:     "followed",
	}
	return h.NewResponseWithData(c, resp)
}
