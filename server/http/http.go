package http

import (
	"context"
	"log/slog"
	"net/http"

	"framework/domain"
	"framework/log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	middlewareOtel "go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

	"framework/config"
)

type HTTPServer struct {
	Echo *echo.Echo
}

type echoValidator struct {
	validator *validator.Validate
}

func (v *echoValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewEcho(logger *log.Logger, config *config.Config) *echo.Echo {
	e := echo.New()

	if config.GetString("env") == "local" {
		e.Debug = true
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	// register validator
	e.Validator = &echoValidator{validator: validator.New()}

	if config.GetBool("apm.enabled") {
		e.Use(middlewareOtel.Middleware(config.GetString("apm.service_name")))
	}

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogLatency:  true,
		LogError:    true,
		LogMethod:   true,
		LogRemoteIP: true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// Get the real IP address
			realIP := c.RealIP()
			method := c.Request().Method
			uri := v.URI
			status := v.Status
			latency := v.Latency.Milliseconds()
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("remote_ip", realIP),
					slog.String("method", method),
					slog.String("uri", uri),
					slog.Int("status", status),
					slog.Int("latency", int(latency)),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("remote_ip", realIP),
					slog.String("method", method),
					slog.String("uri", uri),
					slog.Int("status", status),
					slog.Int("latency", int(latency)),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	// 全局中间件：自动包装 handler 返回 (any, error)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 捕获 panic
			defer func() {
				if r := recover(); r != nil {
					_ = c.JSON(200, domain.Response[any]{Code: 500, Message: "internal server error", Data: nil})
				}
			}()
			// 调用 handler
			err := next(c)
			if c.Response().Committed {
				return nil
			}
			// 从 context 取出 handler 返回值
			resp := c.Get("handler_data")
			if err != nil {
				return c.JSON(200, domain.Response[any]{Code: 400, Message: err.Error(), Data: nil})
			}
			return c.JSON(200, domain.Response[any]{Code: 0, Message: "ok", Data: resp})
		}
	})

	return e
}
