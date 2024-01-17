package internal

import (
	"context"

	schema "github.com/kkpagaev/gorofls/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	db schema.Querier
}

func NewUsers(db schema.Querier) *Users {
	return &Users{db: db}
}

func (u Users) ListUsers(ctx context.Context, page, limit int32) ([]schema.User, error) {
	return u.db.ListUsers(ctx, schema.ListUsersParams{
		Offset: (page - 1) * limit,
		Limit:  limit,
	})
}

type CreateUser struct {
	Name     string
	Password string
	Email    string
}

func (u Users) CreateUser(ctx context.Context, arg CreateUser) (schema.User, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
	if err != nil {
		return schema.User{}, err
	}

	user, err := u.db.CreateUser(ctx, schema.CreateUserParams{
		Name:     arg.Name,
		Password: string(hashed_password),
		Email:    arg.Email,
	})

	return user, err
}
