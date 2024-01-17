package web

import (
	schema "github.com/kkpagaev/gorofls/db/sqlc"
	"github.com/kkpagaev/gorofls/internal"
	"github.com/labstack/echo/v4"
)

type BookGroup struct {
	Books *internal.Books
}

func RegisterBookGroup(e *echo.Group, d BookGroup) {
	g := e.Group("/books")

	g.GET("", d.listBooks)
	// g.POST("", d.createBook)
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
