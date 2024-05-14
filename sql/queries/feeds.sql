-- name: CreateFeed :one
INSERT INTO feeds(created_at, updated_at, title, link)
VALUES ($1, $2, $3, $4)
    RETURNING *;
