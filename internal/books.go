package internal

import (
	"context"

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

func (b Books) CreateBook(ctx context.Context, arg schema.CreateBookParams) (schema.Book, error) {
	return b.db.CreateBook(ctx, arg)
}

func (b Books) DeleteBook(ctx context.Context, id int32) error {
	return b.db.DeleteBook(ctx, id)
}
