package internal

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	schema "github.com/kkpagaev/gorofls/db/sqlc"
)

type Authors struct {
	db schema.Querier
}

func NewAuthors(db schema.Querier) *Authors {
	return &Authors{db: db}
}

type ListAuthors struct {
	Page  int32 `validate:"required,min=1"`
	Limit int32 `validate:"required,min=1,max=100"`
}

func (a Authors) ListAuthors(ctx context.Context, args ListAuthors) ([]schema.Author, error) {
	return a.db.ListAuthors(ctx, schema.ListAuthorsParams{
		Limit:  args.Limit,
		Offset: (args.Page - 1) * args.Limit,
	})
}

type CreateAuthor struct {
	Name string `validate:"required,max=20"`
	Bio  string `validate:"max=1000"`
}

func (a Authors) CreateAuthor(ctx context.Context, arg CreateAuthor) (schema.Author, error) {
	return a.db.CreateAuthor(ctx, schema.CreateAuthorParams{
		Name: arg.Name,
		Bio: pgtype.Text{
			String: arg.Bio,
			Valid:  arg.Bio != "",
		},
	})
}

type UpdateAuthor struct {
	Name string `validate:"max=20"`
	Bio  string `validate:"max=1000"`
}

func (a Authors) UpdateAuthor(ctx context.Context, id int64, arg UpdateAuthor) (schema.Author, error) {
	err := a.db.UpdateAuthor(ctx, schema.UpdateAuthorParams{
		ID: id,
		Name: pgtype.Text{
			String: arg.Name,
			Valid:  arg.Name != "",
		},
		Bio: pgtype.Text{
			String: arg.Bio,
			Valid:  arg.Bio != "",
		},
	})
	if err != nil {
		return schema.Author{}, err
	}

	author, err := a.db.GetAuthor(ctx, id)
	if err != nil {
		return schema.Author{}, err
	}

	return author, nil
}

func (a Authors) DeleteAuthor(ctx context.Context, id int64) error {
	return a.db.DeleteAuthor(ctx, id)
}

func (a Authors) GetAuthor(ctx context.Context, id int64) (schema.Author, error) {
	return a.db.GetAuthor(ctx, id)
}
