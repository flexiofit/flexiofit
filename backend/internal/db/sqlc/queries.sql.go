// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one


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
RETURNING id, email, username, full_name, password_hash, bio, avatar_url, is_active, created_at, updated_at
`

type CreateUserParams struct {
	Email        string
	Username     string
	FullName     string
	PasswordHash string
	Bio          pgtype.Text
	AvatarUrl    pgtype.Text
}

// internal/db/sqlc/queries.sql
// User Queries
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.Username,
		arg.FullName,
		arg.PasswordHash,
		arg.Bio,
		arg.AvatarUrl,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.FullName,
		&i.PasswordHash,
		&i.Bio,
		&i.AvatarUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users
SET
    is_active = false,
    updated_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, username, full_name, password_hash, bio, avatar_url, is_active, created_at, updated_at FROM users
WHERE email = $1 AND is_active = true
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.FullName,
		&i.PasswordHash,
		&i.Bio,
		&i.AvatarUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, username, full_name, password_hash, bio, avatar_url, is_active, created_at, updated_at FROM users
WHERE id = $1 AND is_active = true
`

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.FullName,
		&i.PasswordHash,
		&i.Bio,
		&i.AvatarUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, email, username, full_name, password_hash, bio, avatar_url, is_active, created_at, updated_at FROM users
WHERE username = $1 AND is_active = true
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.FullName,
		&i.PasswordHash,
		&i.Bio,
		&i.AvatarUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserCount = `-- name: GetUserCount :one
SELECT COUNT(*) FROM users
WHERE is_active = true
`

func (q *Queries) GetUserCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getUserCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, username, full_name, password_hash, bio, avatar_url, is_active, created_at, updated_at FROM users
WHERE is_active = true
ORDER BY created_at DESC
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Username,
			&i.FullName,
			&i.PasswordHash,
			&i.Bio,
			&i.AvatarUrl,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchUsers = `-- name: SearchUsers :many
SELECT id, email, username, full_name, password_hash, bio, avatar_url, is_active, created_at, updated_at FROM users
WHERE
    is_active = true AND
    (
        username ILIKE '%' || $1 || '%' OR
        full_name ILIKE '%' || $1 || '%' OR
        email ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type SearchUsersParams struct {
	Column1 pgtype.Text
	Limit   int32
	Offset  int32
}

func (q *Queries) SearchUsers(ctx context.Context, arg SearchUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, searchUsers, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Username,
			&i.FullName,
			&i.PasswordHash,
			&i.Bio,
			&i.AvatarUrl,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    email = COALESCE($1, email),
    username = COALESCE($2, username),
    full_name = COALESCE($3, full_name),
    bio = COALESCE($4, bio),
    avatar_url = COALESCE($5, avatar_url),
    updated_at = NOW()
WHERE id = $6 AND is_active = true
RETURNING id, email, username, full_name, password_hash, bio, avatar_url, is_active, created_at, updated_at
`

type UpdateUserParams struct {
	Email     pgtype.Text
	Username  pgtype.Text
	FullName  pgtype.Text
	Bio       pgtype.Text
	AvatarUrl pgtype.Text
	ID        pgtype.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Email,
		arg.Username,
		arg.FullName,
		arg.Bio,
		arg.AvatarUrl,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.FullName,
		&i.PasswordHash,
		&i.Bio,
		&i.AvatarUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
UPDATE users
SET
    password_hash = $2,
    updated_at = NOW()
WHERE id = $1 AND is_active = true
RETURNING id, email, username, full_name, password_hash, bio, avatar_url, is_active, created_at, updated_at
`

type UpdateUserPasswordParams struct {
	ID           pgtype.UUID
	PasswordHash string
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserPassword, arg.ID, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.FullName,
		&i.PasswordHash,
		&i.Bio,
		&i.AvatarUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
