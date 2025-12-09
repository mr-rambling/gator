-- name: MarkFeedFetched :exec
UPDATE feeds 
SET last_fetched = NOW(), updated_at = NOW() 
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched NULLS FIRST, updated_at
LIMIT 1;