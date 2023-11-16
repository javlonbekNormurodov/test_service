-- +goose Up
-- +goose StatementBegin
CREATE TABLE integrations (
    id varchar PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    name varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT (NOW()),
    updated_at timestamp NOT NULL DEFAULT (NOW()),
    deleted_at timestamp DEFAULT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE integrations;

-- +goose StatementEnd