

-- name: RegisterUser :one
WITH new_user AS (
    INSERT INTO users (id, username, password, email ) VALUES ($1, $2, $3, $4)
    RETURNING id
), new_token AS (
    INSERT INTO user_tokens (user_id, token) VALUES ((SELECT id FROM new_user), $5)
    RETURNING token
) SELECT id, token FROM new_user, new_token;
