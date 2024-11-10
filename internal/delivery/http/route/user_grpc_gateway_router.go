package route

import (
	"books-api/pkg/server"
	"books-api/proto/users/v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log/slog"
)

func (h *Router) GRPCUserHTTPSetup(port string) *runtime.ServeMux {
	grpcServer, netListener := server.NewGRPCServer(port)
	grpcGatewayMux := runtime.NewServeMux()

	users.RegisterUserServiceServer(grpcServer, h.UserHandler)

	grpcClient := server.NewGRPCGatewayClient(grpcServer, netListener, port)

	if err := users.RegisterUserServiceHandler(context.Background(), grpcGatewayMux, grpcClient); err != nil {
		slog.Error("failed to register users gateway")
	}

	return grpcGatewayMux
}
