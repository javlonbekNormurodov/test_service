package service

import (
	"context"
	"encoding/json"
	"fmt"

	integration_service "test_service/generated/integrations"
	"test_service/internal/core/repository"
	"test_service/internal/core/repository/psql/sqlc"
	"test_service/internal/pkg/logger"

	"go.uber.org/zap"
)

type integrationService struct {
	store repository.Store
	integration_service.UnimplementedIntegrationServer
}

func NewIntegrationService(store repository.Store) integration_service.IntegrationServer {
	return &integrationService{
		store: store,
	}
}

func (s *integrationService) CreateIntegration(ctx context.Context, req *integration_service.CreateRequest) (*integration_service.FullResponse, error) {
	res, err := s.store.CreateIntegration(ctx, req.Name)
	if err != nil {
		logger.Log.Error("failed while creating integration: ", zap.Error(err))
		return nil, err
	}

	return mapIntegrationToFullResponse(&res), nil
}

func (s *integrationService) GetIntegrationsList(ctx context.Context, req *integration_service.GetListRequest) (*integration_service.GetListResponse, error) {
	var (
		resp integration_service.GetListResponse
	)

	integrations, err := s.store.ListIntegrations(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get integration list: %v", err)
	}

	bytes, err := json.Marshal(integrations)
	if err != nil {
		return nil, fmt.Errorf("failed to marshalling integration %v", err)
	}

	if err := json.Unmarshal(bytes, &resp.Response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal integration %v", err)
	}

	return &resp, nil
}

func (s *integrationService) GetIntegrationById(ctx context.Context, req *integration_service.GetByIDRequest) (*integration_service.FullResponse, error) {
	integration, err := s.store.GetIntegrationById(ctx, req.Id)
	if err != nil {
		logger.Log.Error("failed while getting integration by id: ", zap.Error(err))
		return nil, err
	}

	return mapIntegrationToFullResponse(&integration), nil
}

func (s *integrationService) UpdateIntegration(ctx context.Context, req *integration_service.UpdateRequest) (*integration_service.FullResponse, error) {
	params := sqlc.UpdateIntegrationParams{
		ID:   req.Id,
		Name: req.Name,
	}

	integration, err := s.store.UpdateIntegration(ctx, params)
	if err != nil {
		logger.Log.Error("failed while updating integration: ", zap.Error(err))
		return nil, err
	}

	return mapIntegrationToFullResponse(&integration), nil
}

func (s *integrationService) DeleteIntegration(ctx context.Context, req *integration_service.DeleteRequest) error {
	return s.store.DeleteIntegration(ctx, req.Id)
}

func mapIntegrationToFullResponse(i *sqlc.Integration) *integration_service.FullResponse {
	return &integration_service.FullResponse{
		Id:        i.ID,
		Name:      i.Name,
		CreatedAt: i.CreatedAt.String(),
		UpdatedAt: i.UpdatedAt.String(),
		DeletedAt: i.DeletedAt.Time.String(),
	}
}
