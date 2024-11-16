-- internal/db/sqlc/queries.sql

-- User Queries

-- name: CreateUser :one
INSERT INTO users (
    email,
    username,
    full_name,
    password_hash,
    bio,
    avatar_url
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 AND is_active = true;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND is_active = true;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 AND is_active = true;

-- name: ListUsers :many
SELECT * FROM users
WHERE is_active = true
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
    email = COALESCE(sqlc.narg('email'), email),
    username = COALESCE(sqlc.narg('username'), username),
    full_name = COALESCE(sqlc.narg('full_name'), full_name),
    bio = COALESCE(sqlc.narg('bio'), bio),
    avatar_url = COALESCE(sqlc.narg('avatar_url'), avatar_url),
    updated_at = NOW()
WHERE id = sqlc.arg('id') AND is_active = true
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
SET
    password_hash = $2,
    updated_at = NOW()
WHERE id = $1 AND is_active = true
RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET
    is_active = false,
    updated_at = NOW()
WHERE id = $1;

-- name: SearchUsers :many
SELECT * FROM users
WHERE
    is_active = true AND
    (
        username ILIKE '%' || $1 || '%' OR
        full_name ILIKE '%' || $1 || '%' OR
        email ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetUserCount :one
SELECT COUNT(*) FROM users
WHERE is_active = true;
