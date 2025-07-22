package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type MomentRequest struct {
	MomentID  string   `json:"moment_id,omitempty"`
	Content   string   `json:"content,omitempty"`
	ImageURLs []string `json:"image_urls,omitempty"`
	Location  string   `json:"location,omitempty"`
}

type CommentRequest struct {
	MomentID string `json:"moment_id,omitempty"`
	Content  string `json:"content,omitempty"`
}

type MomentHandler struct {
	MomentUsecase *usecase.MomentUsecase
}

func RegisterMomentHandler(e *echo.Echo, momentUsecase *usecase.MomentUsecase) {
	h := &MomentHandler{MomentUsecase: momentUsecase}
	group := e.Group("/api/v1/moments")
	group.POST("", h.PostMoment)
	group.POST("/detail", h.PostMoment)
	group.POST("/comments", h.PostComment)
}

func (h *MomentHandler) PostMoment(c echo.Context) error {
	var req MomentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	// TODO: 根据 req 字段判断是列表、详情还是发布，调用 usecase
	return c.JSON(http.StatusOK, map[string]string{"message": "动态列表/详情/发布成功"})
}

func (h *MomentHandler) PostComment(c echo.Context) error {
	var req CommentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	// TODO: 发布或获取评论
	return c.JSON(http.StatusOK, map[string]string{"message": "评论列表/发布成功"})
}
