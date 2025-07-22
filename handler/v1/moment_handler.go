package v1

import (
	"net/http"

	"framework/usecase"

	"github.com/labstack/echo/v4"
)

type CreateMomentRequest struct {
	Content   string   `json:"content"`
	ImageURLs []string `json:"image_urls"`
	Location  string   `json:"location"`
}

type CreateCommentRequest struct {
	Content string `json:"content"`
}

type MomentHandler struct {
	MomentUsecase *usecase.MomentUsecase
}

func RegisterMomentHandler(e *echo.Echo, momentUsecase *usecase.MomentUsecase) {
	h := &MomentHandler{MomentUsecase: momentUsecase}
	group := e.Group("/api/v1/moments")
	group.GET("", h.ListMoments)
	group.POST("", h.CreateMoment)
	group.GET("/:momentId", h.GetMoment)
	group.GET("/:momentId/comments", h.ListComments)
	group.POST("/:momentId/comments", h.CreateComment)
}

func (h *MomentHandler) ListMoments(c echo.Context) error {
	// TODO: 获取动态列表
	return c.JSON(http.StatusOK, map[string]string{"message": "动态列表"})
}

func (h *MomentHandler) CreateMoment(c echo.Context) error {
	// TODO: 发布动态
	return c.JSON(http.StatusCreated, map[string]string{"message": "发布成功"})
}

func (h *MomentHandler) GetMoment(c echo.Context) error {
	// TODO: 获取单条动态详情
	return c.JSON(http.StatusOK, map[string]string{"message": "动态详情"})
}

func (h *MomentHandler) ListComments(c echo.Context) error {
	// TODO: 获取动态评论列表
	return c.JSON(http.StatusOK, map[string]string{"message": "评论列表"})
}

func (h *MomentHandler) CreateComment(c echo.Context) error {
	// TODO: 发布评论
	return c.JSON(http.StatusCreated, map[string]string{"message": "评论成功"})
}
