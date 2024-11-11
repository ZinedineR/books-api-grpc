package route

import (
	"books-api/pkg/server"
	categories "books-api/proto/category/v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log/slog"
)

func (h *Router) GRPCCategoryHTTPSetup(port string) *runtime.ServeMux {
	grpcServer, netListener := server.NewGRPCServer(port)
	grpcGatewayMux := runtime.NewServeMux()

	categories.RegisterCategoryServiceServer(grpcServer, h.CategoryHandler)

	grpcClient := server.NewGRPCGatewayClient(grpcServer, netListener, port)

	if err := categories.RegisterCategoryServiceHandler(context.Background(), grpcGatewayMux, grpcClient); err != nil {
		slog.Error("failed to register authors gateway")
	}

	return grpcGatewayMux
}
