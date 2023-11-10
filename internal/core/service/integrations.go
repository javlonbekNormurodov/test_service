package service

import (
	"context"
	"encoding/json"

	integration_service "test_service/generated/integrations"
	"test_service/internal/core/repository"
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
	err := s.store.CreateIntegration(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &integration_service.FullResponse{}, nil
}

func (s *integrationService) GetIntegrationsList(ctx context.Context, req *integration_service.GetListRequest) (*integration_service.FullResponse, error) {
	var (
		resp            = integration_service.FullResponse{}
		integrationList []*integration_service.FullResponse
	)

	integrations, err := s.store.ListIntegrations(ctx)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(integrations)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &integrationList); err != nil {
		return nil, err
	}

	return &integration_service.FullResponse{
		Id:        resp.Id,
		Name:      resp.Name,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
		DeletedAt: resp.DeletedAt,
	}, nil
}

func (s *integrationService) GetIntegrationById(ctx context.Context, req *integration_service.GetByIDRequest) (*integration_service.FullResponse, error) {

	integrations, err := s.store.GetIntegrationById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &integration_service.FullResponse{
		Id:        integrations.ID,
		Name:      integrations.Name,
		CreatedAt: integrations.CreatedAt.String(),
		UpdatedAt: integrations.UpdatedAt.String(),
		DeletedAt: integrations.DeletedAt.Time.String(),
	}, nil
}
