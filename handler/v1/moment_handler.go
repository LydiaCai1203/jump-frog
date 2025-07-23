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
	group.POST("", h.List)
	group.POST("/detail", h.Detail)
	group.POST("/comments", h.Comments)
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
