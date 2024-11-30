-- name: CreateCollection :one
INSERT INTO collections (id, user_id, title, description, type)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUsersCollectionsList :many
SELECT * FROM collections WHERE user_id = $1 OFFSET $2 LIMIT $3;

-- name: GetUsersCollectionListTotal :one
SELECT count(*) FROM collections WHERE user_id = $1;

-- name: GetUserCollectionByID :one
SELECT * FROM collections WHERE id = $1 AND user_id = $2;

-- name: DeleteUserCollectionByID :exec
DELETE FROM collections where id = $1 AND user_id = $2;