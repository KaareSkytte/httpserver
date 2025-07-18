// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: login.sql

package database

import (
	"context"
	"database/sql"
)

const checkEmail = `-- name: CheckEmail :one
SELECT id, created_at, updated_at, email, hashed_password, is_chirpy_red FROM users
WHERE email = $1
`

func (q *Queries) CheckEmail(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, checkEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}
