package route

import (
	"books-api/pkg/server"
	"books-api/proto/books/v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log/slog"
)

func (h *Router) GRPCBookSetup() *runtime.ServeMux {
	grpcServer, netListener := server.NewGRPCServer("50052")
	grpcGatewayMux := runtime.NewServeMux()

	books.RegisterBookServiceServer(grpcServer, h.BookHandler)

	grpcClient := server.NewGRPCClient(grpcServer, netListener, "50052")

	if err := books.RegisterBookServiceHandler(context.Background(), grpcGatewayMux, grpcClient); err != nil {
		slog.Error("failed to register users gateway")
	}

	return grpcGatewayMux
}
