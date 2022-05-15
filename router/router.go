package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// NewRouter creates a new router instance
func NewRouter() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\u001b[32m${time_rfc3339}\u001b[0m " +
			"\u001b[35m${method}:" +
			"${uri}\u001b[0m " +
			"STATUS:${status}," +
			"ERR:${error}" +
			"\n\u001b[32m${time_rfc3339}\u001b[0m " +
			"HOST:${host}," +
			"LATENCY:${latency_human}," +
			"IP:${remote_ip}," +
			"AGENT:${user_agent}\n",
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	origins := []string{
		"http://localhost:8080",
		"http://192.168.1.248:8080",
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string(origins),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Validator = NewValidator()

	return e
}
