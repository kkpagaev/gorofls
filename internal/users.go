package internal

import (
	"context"

	schema "github.com/kkpagaev/gorofls/db/sqlc"
)

type Users struct {
	db *schema.Queries
}

func NewUsers(db *schema.Queries) *Users {
	return &Users{db: db}
}

func (u Users) ListUsers(ctx context.Context, page, limit int32) ([]schema.User, error) {
	return u.db.ListUsers(ctx, schema.ListUsersParams{
		Offset: (page - 1) * limit,
		Limit:  limit,
	})
}
