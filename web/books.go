package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookGroup struct{}

func RegisterBookGroup(e *echo.Group, d BookGroup) {
	g := e.Group("/books")

	g.GET("", d.hello)
}

func (b BookGroup) hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello users")
}
