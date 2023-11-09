package handlers

import (
	"context"
	"fmt"

	"test_service/internal/config"
	"test_service/pkg/wrapper"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func New(ctx context.Context, cfg *config.Config) *runtime.ServeMux {
	gwMux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(wrapper.CustomMatcher),
	)

	return gwMux
}

func makeHost(host string, port int32) string {
	return host + ":" + fmt.Sprintf("%d", port)
}
