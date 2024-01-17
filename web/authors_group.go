package web

import (
	"github.com/kkpagaev/gorofls/internal"
	"github.com/labstack/echo/v4"
)

type AuthorsGroup struct {
	Authors *internal.Authors
}

func RegisterAuthorsGroup(g *echo.Group, d AuthorsGroup) {
	g.GET("", d.listAuthors)
	g.POST("", d.createAuthor)
	g.PATCH("/:id", d.updateAuthor)
	g.DELETE("/:id", d.deleteAuthor)
}

func (d AuthorsGroup) listAuthors(c echo.Context) error {
	var query struct {
		Page  int32 `query:"page"`
		Limit int32 `query:"limit"`
	}
	if err := c.Bind(&query); err != nil {
		return err
	}
	listAuthors := internal.ListAuthors{
		Limit: query.Limit,
		Page:  query.Page,
	}
	if err := c.Validate(&listAuthors); err != nil {
		return nil
	}
	ctx := c.Request().Context()
	authors, err := d.Authors.ListAuthors(ctx, listAuthors)

	if err != nil {
		return err
	}
	return ok(c, authors)
}

func (d AuthorsGroup) findAuthor(c echo.Context) error {
	var params struct {
		ID int64 `param:"id"`
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	ctx := c.Request().Context()
	author, err := d.Authors.GetAuthor(ctx, params.ID)
	if err != nil {
		return err
	}
	return ok(c, author)
}

func (d AuthorsGroup) createAuthor(c echo.Context) error {
	var body struct {
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	createAuthor := internal.CreateAuthor{
		Name: body.Name,
		Bio:  body.Bio,
	}
	if err := c.Validate(createAuthor); err != nil {
		return err
	}

	ctx := c.Request().Context()
	author, err := d.Authors.CreateAuthor(ctx, createAuthor)

	if err != nil {
		return err
	}
	return ok(c, author)
}

func (d AuthorsGroup) updateAuthor(c echo.Context) error {
	var params struct {
		ID   int64  `param:"id"`
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}
	if err := c.Bind(&params); err != nil {
		return err
	}
	updateAuthor := internal.UpdateAuthor{
		Name: params.Name,
		Bio:  params.Bio,
	}
	if err := c.Validate(updateAuthor); err != nil {
		return err
	}
	ctx := c.Request().Context()
	author, err := d.Authors.UpdateAuthor(ctx, params.ID, updateAuthor)
	if err != nil {
		return err
	}

	return ok(c, author)
}

func (d AuthorsGroup) deleteAuthor(c echo.Context) error {
	var params struct {
		ID int64 `param:"id"`
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	ctx := c.Request().Context()
	err := d.Authors.DeleteAuthor(ctx, params.ID)

	if err != nil {
		return err
	}

	return message(c, "deleted")
}
