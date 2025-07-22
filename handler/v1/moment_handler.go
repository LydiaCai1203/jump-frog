package v1

import (
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
	moments := []map[string]any{
		{"moment_id": "1", "content": "hello", "user_id": "1"},
	}
	return h.NewResponseWithData(c, moments)
}

func (h *MomentHandler) Detail(c echo.Context) error {
	var req struct {
		MomentID string `json:"moment_id"`
	}
	if err := c.Bind(&req); err != nil || req.MomentID == "" {
		return h.NewResponseWithError(c, "moment_id is required", err)
	}
	detail := map[string]any{
		"moment_id": req.MomentID,
		"content":   "hello",
		"user_id":   "1",
	}
	return h.NewResponseWithData(c, detail)
}

func (h *MomentHandler) Comments(c echo.Context) error {
	var req struct {
		MomentID string `json:"moment_id"`
	}
	if err := c.Bind(&req); err != nil || req.MomentID == "" {
		return h.NewResponseWithError(c, "moment_id is required", err)
	}
	comments := []map[string]any{
		{
			"comment_id": "1",
			"moment_id":  req.MomentID,
			"content":    "nice!",
			"user_id":    "2",
		},
	}
	return h.NewResponseWithData(c, comments)
}
