-- +goose Up
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name       VARCHAR   NOT NULL
);

-- +goose Down
DROP TABLE users;
