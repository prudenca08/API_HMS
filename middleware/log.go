package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time= ${time_rfc3339_nano}, method= ${method}, uri= ${uri}, status= ${status}\n",
	}))
}
