package web

import (
	"net/http"

	"github.com/kkpagaev/gorofls/internal"
	"github.com/labstack/echo/v4"
)

type UserGroup struct {
	Users *internal.Users
}

func RegisterUserGroup(g *echo.Group, d UserGroup) {
	g.GET("", d.listUsers)
}

func (u UserGroup) listUsers(c echo.Context) error {
	var query struct {
		page  int32 `query:"page" validate:"required,min=1"`
		limit int32 `query:"limit" validate:"required,min=1,max=100"`
	}
	if err := ValidateRequest(c, &query); err != nil {
		return err
	}
	users, err := u.Users.ListUsers(c.Request().Context(), query.page, query.limit)

	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, users)
	}
}
