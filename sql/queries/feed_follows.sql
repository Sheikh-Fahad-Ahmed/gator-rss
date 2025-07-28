-- name: CreateFeedFollow :one
WITH feed_follow_record AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    feed_follow_record.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follow_record
INNER JOIN feeds ON feeds.id = feed_follow_record.feed_id
INNER JOIN users ON users.id = feed_follow_record.user_id;

-- name: GetFeedFollowForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
INNER JOIN users ON users.id = feed_follows.user_id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_id = $1 AND user_id = $2;