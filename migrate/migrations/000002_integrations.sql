-- +goose Up
-- +goose StatementBegin
CREATE TABLE integrations (
    id varchar PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    name varchar NOT NULL,
    created_at timestamp WITH time zone NOT NULL DEFAULT (NOW()),
    updated_at timestamp WITH time zone NOT NULL DEFAULT (NOW()),
    deleted_at timestamp WITH time zone DEFAULT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE integrations;

-- +goose StatementEnd