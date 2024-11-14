-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1,
        $2,
        $3,
        $4)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE name = $1
LIMIT 1;

-- name: DeleteAllUsers :exec
DELETE
FROM users;

-- name: GetAllUsers :many
SELECT *
FROM users;

-- name: GetFeedFollowsForUser :many
SELECT users.name as user_name, feeds.name as feeds_name
from users
         JOIN feed_follows
              ON users.id = feed_follows.user_id
         JOIN feeds
              ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;
