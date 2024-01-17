-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT 
  books.*,
  authors.name AS author_name
  FROM books
JOIN authors ON books.author_id = authors.id
ORDER BY books.id
LIMIT $1
OFFSET $2;

-- name: CreateBook :one
INSERT INTO books (title, author_id, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: UpdateBook :one
UPDATE books
  SET title = COALESCE(sqlc.narg(title), title),
      author_id = COALESCE(sqlc.narg(author_id), author_id),
      description = COALESCE(sqlc.narg(description), description)
WHERE id = $1
RETURNING *;

