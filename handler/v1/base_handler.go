package v1

import (
	"net/http"

	"framework/domain"

	"github.com/labstack/echo/v4"
)

type BaseHandler struct{}

// NewResponseWithData 返回统一成功响应
func (h *BaseHandler) NewResponseWithData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, domain.Response[any]{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// NewResponseWithError 返回统一错误响应
func (h *BaseHandler) NewResponseWithError(c echo.Context, msg string, err error) error {
	return c.JSON(http.StatusOK, domain.Response[any]{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Data:    nil,
	})
}
