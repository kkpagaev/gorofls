package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BookGroup(e *echo.Group) {
	g := e.Group("/books")

	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello users")
	})
}
