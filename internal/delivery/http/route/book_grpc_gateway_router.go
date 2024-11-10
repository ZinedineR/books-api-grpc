package route

import (
	"books-api/pkg/server"
	"books-api/proto/books/v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log/slog"
)

func (h *Router) GRPCBookHTTPSetup(port string) *runtime.ServeMux {
	grpcServer, netListener := server.NewGRPCServer(port)
	grpcGatewayMux := runtime.NewServeMux()

	books.RegisterBookServiceServer(grpcServer, h.BookHandler)

	grpcClient := server.NewGRPCGatewayClient(grpcServer, netListener, port)

	if err := books.RegisterBookServiceHandler(context.Background(), grpcGatewayMux, grpcClient); err != nil {
		slog.Error("failed to register books gateway")
	}

	return grpcGatewayMux
}
