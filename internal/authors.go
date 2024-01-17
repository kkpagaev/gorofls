package internal

import (
	"context"

	schema "github.com/kkpagaev/gorofls/db/sqlc"
)

type Authors struct {
	db schema.Querier
}

func NewAuthors(db schema.Querier) *Authors {
	return &Authors{db: db}
}

func (a Authors) ListAuthors(ctx context.Context, page, limit int32) ([]schema.Author, error) {
	return a.db.ListAuthors(ctx, schema.ListAuthorsParams{
		Offset: (page - 1) * limit,
		Limit:  limit,
	})
}
