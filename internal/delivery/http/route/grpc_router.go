package route

import (
	"books-api/pkg/server"
	"books-api/proto/users/v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log/slog"
)

func (h *Router) GRPCUserSetup() *runtime.ServeMux {
	grpcServer, netListener := server.NewGRPCServer("50051")
	grpcGatewayMux := runtime.NewServeMux()

	users.RegisterUserServiceServer(grpcServer, h.GRPCHandler.User)

	grpcClient := server.NewGRPCClient(grpcServer, netListener, "50051")

	if err := users.RegisterUserServiceHandler(context.Background(), grpcGatewayMux, grpcClient); err != nil {
		slog.Error("failed to register users gateway")
	}

	return grpcGatewayMux
}
