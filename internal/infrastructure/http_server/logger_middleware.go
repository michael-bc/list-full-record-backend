package http_server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/logger"
)

const (
	LogFormat = `{` +
		`"time":"${time_custom}",` +
		`"request_id":"${id}",` +
		`"remote_ip":"${remote_ip}",` +
		`"protocol":"${protocol}",` +
		`"user_agent":"${user_agent}",` +
		`"host":"${host}",` +
		`"method":"${method}",` +
		`"uri":"${uri}",` +
		`"user_agent":"${user_agent}",` +
		`"status": ${status},` +
		`"error":"${error}",` +
		`"latency":${latency},` +
		`"latency_human":"${latency_human}",` +
		`"bytes_in":${bytes_in},` +
		`"bytes_out":${bytes_out},` +
		`"referer":${referer}` +
		`}` +
		"\n"

	LogTimeFormat = "2006-01-02T15:04:05Z"
)

func LoggerMiddleware(log logger.Logger) echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          LoggerSkipper(),
		Format:           LogFormat,
		CustomTimeFormat: LogTimeFormat,
		Output:           log.Writer(),
	})
}

func LoggerSkipper() middleware.Skipper {
	return func(c echo.Context) bool {
		switch c.Request().RequestURI {
		case "/":
			return true
		default:
			return middleware.DefaultSkipper(c)
		}
	}
}
