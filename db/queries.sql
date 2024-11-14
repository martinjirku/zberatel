

-- name: RegisterUser :one
WITH new_user AS (
    INSERT INTO users (id, username, password, email ) VALUES ($1, $2, $3, $4)
    RETURNING id
), new_token AS (
    INSERT INTO user_tokens (user_id, token) VALUES ((SELECT id FROM new_user), $5)
    RETURNING token
) SELECT id, token FROM new_user, new_token;

-- name: GetUserLogin :one
SELECT id, username, password, email FROM users WHERE username = $1;

-- name: CreateCollection :one
INSERT INTO collections (id, user_id, title, description, type)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetCollectionsList :many
SELECT * FROM collections WHERE user_id = $1 OFFSET $2 LIMIT $3;

-- name: GetCollectionListTotal :one
SELECT count(*) FROM collections WHERE user_id = $1;

-- name: GetCollectionByID :one
SELECT * FROM collections WHERE id = $1;