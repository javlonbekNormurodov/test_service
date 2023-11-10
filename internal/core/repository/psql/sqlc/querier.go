// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateIntegration(ctx context.Context, name string) error
	DeleteIntegration(ctx context.Context, id string) error
	GetIntegrationById(ctx context.Context, id string) (Integration, error)
	ListIntegrations(ctx context.Context) ([]Integration, error)
	UpdateIntegration(ctx context.Context, arg UpdateIntegrationParams) error
}

var _ Querier = (*Queries)(nil)
