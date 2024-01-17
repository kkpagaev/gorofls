package internal

import (
	"context"
	"errors"

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
	Email    string `validate:"required,email"`
	Name     string `validate:"required"`
	Password string `validate:"required,min=6"`
}

var UserEmailExists = errors.New("user with email already exists")
var UserNameExists = errors.New("user with name already exists")

func (arg CreateUser) validate(ctx context.Context, db schema.Querier) error {
	email_exists, err := db.UserEmailExists(ctx, arg.Email)

	if err != nil {
		return err
	}
	if email_exists {
		return UserEmailExists
	}

	name_exisits, err := db.UserNameExists(ctx, arg.Name)

	if name_exisits {
		return UserNameExists
	}

	return err
}

func (u Users) CreateUser(ctx context.Context, arg CreateUser) (schema.User, error) {
	if err := arg.validate(ctx, u.db); err != nil {
		return schema.User{}, err
	}

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
