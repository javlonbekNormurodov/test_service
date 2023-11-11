package grpc

import (
	integration_service "test_service/generated/integrations"
	"test_service/internal/config"
	"test_service/internal/core/repository"
	"test_service/internal/core/service"
	"test_service/internal/transport/grpc/middleware"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(repo repository.Store, cfg *config.Config) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.GrpcLoggerMiddleware,
			),
		),
	)

	reflection.Register(grpcServer)
	integration_service.RegisterIntegrationServer(grpcServer, service.NewIntegrationService(repo))

	return grpcServer
}
