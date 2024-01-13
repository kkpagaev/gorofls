package web

import (
	"net/http"

	"github.com/kkpagaev/gorofls/internal"
	"github.com/labstack/echo/v4"
)

type (
	UserGroupDeps struct {
		Users *internal.Users
	}

	UserListQuery struct {
		Page  int32 `query:"page" validate:"required,min=1"`
		Limit int32 `query:"limit" validate:"required,min=1,max=100"`
	}
)

func UserGroup(g *echo.Group, d UserGroupDeps) {
	g.GET("", func(c echo.Context) error {
		ctx := c.Request().Context()
		var query UserListQuery
		if err := c.Bind(&query); err != nil {
			return err
		}

		if err := c.Validate(query); err != nil {
			return err
		}
		users, err := d.Users.ListUsers(ctx, 1, 10)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		return c.JSON(http.StatusOK, users)
	})
}
