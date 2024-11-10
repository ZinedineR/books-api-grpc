package route

import (
	"books-api/pkg/server"
	authors "books-api/proto/author/v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log/slog"
)

func (h *Router) GRPCAuthorHTTPSetup(port string) *runtime.ServeMux {
	grpcServer, netListener := server.NewGRPCServer(port)
	grpcGatewayMux := runtime.NewServeMux()

	authors.RegisterAuthorServiceServer(grpcServer, h.AuthorHandler)

	grpcClient := server.NewGRPCGatewayClient(grpcServer, netListener, port)

	if err := authors.RegisterAuthorServiceHandler(context.Background(), grpcGatewayMux, grpcClient); err != nil {
		slog.Error("failed to register authors gateway")
	}

	return grpcGatewayMux
}
