-- name: StoreFeed :one
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

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeedIDByURL :one
SELECT id FROM feeds
WHERE url = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $2, last_fetched = $3
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
     ORDER BY last_fetched NULLS FIRST
     LIMIT 1;