// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package schema

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (CreateAuthorRow, error)
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAuthor(ctx context.Context, id int64) error
	DeleteBook(ctx context.Context, id int32) error
	GetAuthor(ctx context.Context, id int64) (Author, error)
	GetBook(ctx context.Context, id int32) (Book, error)
	ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]ListBooksRow, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error)
}

var _ Querier = (*Queries)(nil)
