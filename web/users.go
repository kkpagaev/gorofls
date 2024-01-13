package web

import (
	"net/http"

	"github.com/kkpagaev/gorofls/internal"
	"github.com/labstack/echo/v4"
)

type UserGroupDeps struct {
	Users *internal.Users
}

func UserGroup(g *echo.Group, d UserGroupDeps) {
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello users")
	})
}
