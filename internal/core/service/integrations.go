package service

import (
	"context"

	integration_service "test_service/generated/integrations"
	"test_service/internal/core/repository"
	"test_service/internal/core/repository/psql/sqlc"
	"test_service/internal/pkg/logger"

	"go.uber.org/zap"
)

type IntegrationSService interface {
	CreateIntegration(ctx context.Context, req *integration_service.CreateRequest) (*integration_service.FullResponse, error)
	GetIntegrationsList(ctx context.Context, req *integration_service.GetListRequest) ([]*integration_service.GetListResponse, error)
	GetIntegrationById(ctx context.Context, req *integration_service.GetByIDRequest) (*integration_service.FullResponse, error)
	UpdateIntegration(ctx context.Context, req *integration_service.UpdateRequest) (*integration_service.FullResponse, error)
	DeleteIntegration(ctx context.Context, req *integration_service.DeleteRequest) error
}

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

func (s *integrationService) GetIntegrationsList(ctx context.Context, req *integration_service.GetListRequest) ([]*integration_service.GetListResponse, error) {
	integrations, err := s.store.ListIntegrations(ctx)
	if err != nil {
		logger.Log.Error("failed while getting list of integrations: ", zap.Error(err))
		return nil, err
	}

	getList := mapIntegrationListToFullResponse(integrations)
	return getList, nil
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

func mapIntegrationListToFullResponse(integrations []sqlc.Integration) []*integration_service.GetListResponse {

	var (
		result []*integration_service.GetListResponse
		resp   []*integration_service.FullResponse
	)
	for _, i := range integrations {
		resp = append(resp, mapIntegrationToFullResponse(&i))
		result = append(result, &integration_service.GetListResponse{
			Response: resp,
		})
	}
	return result
}
