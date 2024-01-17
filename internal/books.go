package internal

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	schema "github.com/kkpagaev/gorofls/db/sqlc"
)

type Books struct {
	db schema.Querier
}

func NewBooks(db schema.Querier) *Books {
	return &Books{db: db}
}

func (b Books) ListBooks(ctx context.Context, params schema.ListBooksParams) ([]schema.ListBooksRow, error) {
	return b.db.ListBooks(ctx, params)
}

func (b Books) GetBook(ctx context.Context, id int32) (schema.Book, error) {
	return b.db.GetBook(ctx, id)
}

type CreateBook struct {
	Title       string `validate:"required,min=1,max=100"`
	AuthorID    int64  `validate:"required"`
	Description string `validate:"max=1000"`
}

func (b Books) CreateBook(ctx context.Context, arg CreateBook) (schema.Book, error) {
	return b.db.CreateBook(ctx, schema.CreateBookParams{
		Title:    arg.Title,
		AuthorID: arg.AuthorID,
		Description: pgtype.Text{
			String: arg.Description,
			Valid:  arg.Description != "",
		},
	})
}

type UpdateBook struct {
	Title       string `validate:"min=1,max=100"`
	AuthorID    int64
	Description string `validate:"max=1000"`
}

func (b Books) UpdateBook(ctx context.Context, id int32, arg UpdateBook) (schema.Book, error) {
	return b.db.UpdateBook(ctx, schema.UpdateBookParams{
		ID: id,
		Title: pgtype.Text{
			String: arg.Title,
			Valid:  arg.Title != "",
		},
		AuthorID: pgtype.Int8{
			Int64: arg.AuthorID,
			Valid: arg.AuthorID != 0,
		},
		Description: pgtype.Text{
			String: arg.Description,
			Valid:  arg.Description != "",
		},
	})
}

func (b Books) DeleteBook(ctx context.Context, id int32) error {
	return b.db.DeleteBook(ctx, id)
}
