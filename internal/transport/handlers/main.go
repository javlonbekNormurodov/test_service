package handlers

import (
	"context"
	"fmt"

	integration_service "test_service/generated/integrations"
	"test_service/internal/config"
	"test_service/pkg/wrapper"

	mainGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func New(ctx context.Context, cfg *config.Config) *runtime.ServeMux {
	gwMux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(wrapper.CustomMatcher),
	)

	connIntegrationService, err := mainGrpc.Dial(
		fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port),
		mainGrpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil
	}

	if err := integration_service.RegisterIntegrationHandler(ctx, gwMux, connIntegrationService); err != nil {
		return nil
	}

	return gwMux
}

func makeHost(host string, port int32) string {
	return host + ":" + fmt.Sprintf("%d", port)
}
