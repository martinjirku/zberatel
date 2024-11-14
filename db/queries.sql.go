// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"

	"jirku.sk/zberatel/model"
)

const createCollection = `-- name: CreateCollection :one
INSERT INTO collections (id, user_id, title, description, type)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, title, description, type, created_at, updated_at
`

type CreateCollectionParams struct {
	ID          model.KSUID
	UserID      model.KSUID
	Title       string
	Description sql.NullString
	Type        string
}

func (q *Queries) CreateCollection(ctx context.Context, arg CreateCollectionParams) (Collection, error) {
	row := q.db.QueryRowContext(ctx, createCollection,
		arg.ID,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.Type,
	)
	var i Collection
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCollectionByID = `-- name: GetCollectionByID :one
SELECT id, user_id, title, description, type, created_at, updated_at FROM collections WHERE id = $1
`

func (q *Queries) GetCollectionByID(ctx context.Context, id model.KSUID) (Collection, error) {
	row := q.db.QueryRowContext(ctx, getCollectionByID, id)
	var i Collection
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCollectionListTotal = `-- name: GetCollectionListTotal :one
SELECT count(*) FROM collections WHERE user_id = $1
`

func (q *Queries) GetCollectionListTotal(ctx context.Context, userID model.KSUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCollectionListTotal, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getCollectionsList = `-- name: GetCollectionsList :many
SELECT id, user_id, title, description, type, created_at, updated_at FROM collections WHERE user_id = $1 OFFSET $2 LIMIT $3
`

type GetCollectionsListParams struct {
	UserID model.KSUID
	Offset int32
	Limit  int32
}

func (q *Queries) GetCollectionsList(ctx context.Context, arg GetCollectionsListParams) ([]Collection, error) {
	rows, err := q.db.QueryContext(ctx, getCollectionsList, arg.UserID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserLogin = `-- name: GetUserLogin :one
SELECT id, username, password, email FROM users WHERE username = $1
`

type GetUserLoginRow struct {
	ID       model.KSUID
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
	ID       model.KSUID
	Username string
	Password string
	Email    string
	Token    model.KSUID
}

type RegisterUserRow struct {
	ID    model.KSUID
	Token model.KSUID
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
