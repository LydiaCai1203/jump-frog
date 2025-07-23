package v1

import (
	"framework/domain"
	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type MomentHandler struct {
	BaseHandler
	MomentUsecase *usecase.MomentUsecase
}

func RegisterMomentHandler(e *echo.Echo, momentUsecase *usecase.MomentUsecase) {
	h := &MomentHandler{MomentUsecase: momentUsecase}
	group := e.Group("/api/v1/moments")
	group.POST("/list", h.List)
	group.POST("/post", h.Post)
	group.POST("/:momentId", h.Detail)
	group.POST("/:momentId/comments", h.PostComment)
}

func (h *MomentHandler) List(c echo.Context) error {
	var req domain.MomentListRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	// TODO: 调用 usecase 获取数据
	resp := domain.MomentListResponse{
		List:  []domain.MomentDetail{},
		Total: 0,
		Page:  req.Page,
		Limit: req.Limit,
	}
	return h.NewResponseWithData(c, resp)
}

func (h *MomentHandler) Detail(c echo.Context) error {
	var req domain.MomentDetailRequest
	if err := c.Bind(&req); err != nil || req.MomentID == "" {
		return h.NewResponseWithError(c, "moment_id is required", err)
	}
	// TODO: 调用 usecase 获取数据
	resp := domain.MomentDetail{}
	return h.NewResponseWithData(c, resp)
}

func (h *MomentHandler) Comments(c echo.Context) error {
	var req domain.CommentPostRequest
	if err := c.Bind(&req); err != nil || req.MomentID == "" {
		return h.NewResponseWithError(c, "moment_id is required", err)
	}
	// TODO: 调用 usecase 获取数据
	resp := []domain.CommentPostRequest{}
	return h.NewResponseWithData(c, resp)
}

// 新增 Post 和 PostComment handler 方法
func (h *MomentHandler) Post(c echo.Context) error {
	var req domain.MomentPostRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	// TODO: usecase.Post
	return h.NewResponseWithData(c, nil)
}

func (h *MomentHandler) PostComment(c echo.Context) error {
	var req domain.CommentPostRequest
	if err := c.Bind(&req); err != nil {
		return h.NewResponseWithError(c, "invalid request", err)
	}
	// TODO: usecase.PostComment
	return h.NewResponseWithData(c, nil)
}
