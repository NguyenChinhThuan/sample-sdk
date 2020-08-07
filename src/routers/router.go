package routers

import (
	"github.com/NguyenChinhThuan/sample-sdk/src/handlers"
	"github.com/labstack/echo"
)

// Setup is config router
func Setup(server *echo.Echo) {

	server.GET("/", handlers.Healthcheck)

}
