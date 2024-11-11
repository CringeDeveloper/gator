-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6
)
RETURNING *;

-- name: GetFeedsWithAuthor :many
SELECT * FROM feeds
LEFT JOIN users on feeds.user_id = users.id;
