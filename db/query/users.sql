-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users
(name, password, email)
VALUES
($1, $2, $3)
RETURNING *;

-- name: UserEmailExists :one

SELECT EXISTS (SELECT 1 FROM users WHERE email = $1);

-- name: UserNameExists :one
SELECT EXISTS (SELECT 1 FROM users WHERE name = $1);

