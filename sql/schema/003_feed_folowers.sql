-- +goose Up
CREATE TABLE feeds_followers
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    feed_id    BIGSERIAL NOT NULL REFERENCES feeds (id) ON DELETE CASCADE,
    user_id    BIGSERIAL NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (feed_id, user_id)
);

-- +goose Down
DROP TABLE feeds_followers;
