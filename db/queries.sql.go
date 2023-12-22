// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: queries.sql

package db

import (
	"context"
)

const getUserLogin = `-- name: GetUserLogin :one
SELECT id, username, password, email FROM users WHERE username = $1
`

type GetUserLoginRow struct {
	ID       []byte
	Username string
	Password string
	Email    string
}

func (q *Queries) GetUserLogin(ctx context.Context, username string) (GetUserLoginRow, error) {
	row := q.db.QueryRowContext(ctx, getUserLogin, username)
	var i GetUserLoginRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
	)
	return i, err
}

const registerUser = `-- name: RegisterUser :one
WITH new_user AS (
    INSERT INTO users (id, username, password, email ) VALUES ($1, $2, $3, $4)
    RETURNING id
), new_token AS (
    INSERT INTO user_tokens (user_id, token) VALUES ((SELECT id FROM new_user), $5)
    RETURNING token
) SELECT id, token FROM new_user, new_token
`

type RegisterUserParams struct {
	ID       []byte
	Username string
	Password string
	Email    string
	Token    []byte
}

type RegisterUserRow struct {
	ID    []byte
	Token []byte
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) (RegisterUserRow, error) {
	row := q.db.QueryRowContext(ctx, registerUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.Token,
	)
	var i RegisterUserRow
	err := row.Scan(&i.ID, &i.Token)
	return i, err
}
