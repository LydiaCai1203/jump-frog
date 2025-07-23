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
	group.POST("/follow", h.Follow)
	group.POST("/unflow", h.Unfollow)

	e.POST("/api/v1/users/:userId/followers", h.Followers)
	e.POST("/api/v1/users/:userId/followees", h.Followees)
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

// 新增 Unfollow、Followers、Followees handler 方法
func (h *FollowHandler) Unfollow(c echo.Context) error {
	// TODO: usecase.Unfollow
	return h.NewResponseWithData(c, nil)
}

func (h *FollowHandler) Followers(c echo.Context) error {
	// TODO: usecase.Followers
	return h.NewResponseWithData(c, nil)
}

func (h *FollowHandler) Followees(c echo.Context) error {
	// TODO: usecase.Followees
	return h.NewResponseWithData(c, nil)
}
