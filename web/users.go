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
	g.POST("", d.createUser)
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

func (u UserGroup) createUser(c echo.Context) error {
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required,min=6"`
	}
	if err := ValidateRequest(c, &body); err != nil {
		return err
	}
	user, err := u.Users.CreateUser(c.Request().Context(), internal.CreateUser{
		Email:    body.Email,
		Name:     body.Name,
		Password: body.Password,
	})

	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, user)
	}
}
