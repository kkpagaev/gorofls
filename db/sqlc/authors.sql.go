// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: authors.sql

package schema

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING id, name, bio
`

type CreateAuthorParams struct {
	Name string      `json:"name"`
	Bio  pgtype.Text `json:"bio"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRow(ctx, createAuthor, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRow(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio FROM authors
ORDER BY name
LIMIT $1
OFFSET $2
`

type ListAuthorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error) {
	rows, err := q.db.Query(ctx, listAuthors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthor = `-- name: UpdateAuthor :exec
UPDATE authors
  set name = COALESCE($2, name),
  bio = COALESCE($3, bio)
WHERE id = $1
`

type UpdateAuthorParams struct {
	ID   int64       `json:"id"`
	Name pgtype.Text `json:"name"`
	Bio  pgtype.Text `json:"bio"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.db.Exec(ctx, updateAuthor, arg.ID, arg.Name, arg.Bio)
	return err
}
