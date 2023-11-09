package middleware

import (
	"context"

	"test_service/internal/pkg/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func GrpcLoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logger.Log.Info("----> req: ", zap.String("method", info.FullMethod), zap.Any("req", req))
	resp, err = handler(ctx, req)

	if err != nil {
		logger.Log.Error("failed to make gRPC call: ", zap.Error(err))
		return resp, err
	}

	logger.Log.Info("<---- resp: ", zap.String("method", info.FullMethod))
	return resp, nil
}
