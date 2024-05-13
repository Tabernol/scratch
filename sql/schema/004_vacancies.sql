-- +goose Up
CREATE TABLE vacancies
(
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    title       VARCHAR   NOT NULL,
    description VARCHAR   NOT NULL,
    link        VARCHAR   NOT NULL,
    category    VARCHAR   NOT NULL
);

-- +goose Down
DROP TABLE vacancies;
