-- +goose Up
CREATE TABLE feeds (
                       id BIGSERIAL PRIMARY KEY,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       title VARCHAR NOT NULL,
                       link VARCHAR NOT NULL
);

-- +goose Down
DROP TABLE feeds;
