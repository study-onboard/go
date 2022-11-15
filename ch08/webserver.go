package ch08

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func WebServer() {
	webserver := echo.New()
	webserver.Use(middleware.Logger())
	webserver.Use(middleware.Recover())

	webserver.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Show me the money")
	})

	webserver.Logger.Fatal(webserver.Start(":9090"))
}
