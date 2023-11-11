-- name: ListIntegrations :many
SELECT * FROM integrations;

-- name: CreateIntegration :one
INSERT INTO integrations (name, created_at, updated_at, deleted_at)
VALUES ($1, NOW(), NOW(), NULL)
RETURNING id, name, created_at, updated_at, deleted_at;


-- name: GetIntegrationById :one
SELECT * FROM integrations where id = $1;

-- name: UpdateIntegration :one
UPDATE integrations SET name = $1,updated_at = current_timestamp, deleted_at = NULL WHERE id = $2 RETURNING id, name, created_at, updated_at, deleted_at;

-- name: DeleteIntegration :exec
UPDATE integrations SET deleted_at = current_timestamp WHERE id = $1;