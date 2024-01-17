package web

import (
	schema "github.com/kkpagaev/gorofls/db/sqlc"
	"github.com/kkpagaev/gorofls/internal"
	"github.com/labstack/echo/v4"
)

type BookGroup struct {
	Books *internal.Books
}

func RegisterBookGroup(g *echo.Group, d BookGroup) {
	g.GET("", d.listBooks)
	g.POST("", d.createBook)
	g.PATCH("/:id", d.updateBook)
}

func (d BookGroup) listBooks(c echo.Context) error {
	var query struct {
		Page  int32 `query:"page" validate:"required,min=1"`
		Limit int32 `query:"limit" validate:"required,min=1,max=100"`
	}
	if err := ValidateRequest(c, &query); err != nil {
		return err
	}

	books, err := d.Books.ListBooks(c.Request().Context(), schema.ListBooksParams{
		Limit:  query.Limit,
		Offset: (query.Page - 1) * query.Limit,
	})

	if err != nil {
		return err
	}

	return ok(c, books)
}

func (d BookGroup) createBook(c echo.Context) error {
	var body struct {
		Title       string `json:"title"`
		AuthorID    int64  `json:"author_id"`
		Description string `json:"description"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	createBook := internal.CreateBook{
		Title:       body.Title,
		AuthorID:    body.AuthorID,
		Description: body.Description,
	}
	if err := c.Validate(&createBook); err != nil {
		return err
	}
	ctx := c.Request().Context()
	book, err := d.Books.CreateBook(ctx, createBook)
	if err != nil {
		return err
	}
	return ok(c, book)
}

func (d BookGroup) updateBook(c echo.Context) error {
	var body struct {
		ID          int32  `param:"id"`
		Title       string `json:"title"`
		AuthorID    int64  `json:"author_id"`
		Description string `json:"description"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	updateBook := internal.UpdateBook{
		Title:       body.Title,
		AuthorID:    body.AuthorID,
		Description: body.Description,
	}
	if err := c.Validate(&updateBook); err != nil {
		return err
	}
	ctx := c.Request().Context()
	book, err := d.Books.UpdateBook(ctx, body.ID, updateBook)
	if err != nil {
		return err
	}
	return ok(c, book)
}
