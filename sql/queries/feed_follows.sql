-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
        VALUES ($1, $2, $3 ,$4, $5)
        RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
         INNER JOIN users on inserted_feed_follow.user_id = users.id
         INNER JOIN feeds on inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedsFollowedByUser :many
WITH inserted_feeds AS (
    SELECT * FROM feed_follows
    WHERE user_id = (
        SELECT id FROM users
        WHERE users.name = $1
    )
)
SELECT
    inserted_feeds.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feeds
         INNER JOIN users on inserted_feeds.user_id = users.id
         INNER JOIN feeds on inserted_feeds.feed_id = feeds.id;

-- name: Unfollow :exec
DELETE FROM feed_follows
WHERE feed_id = (
    SELECT id FROM feeds
    WHERE feeds.user_id = $1
);