package route

import (
	"books-api/internal/delivery/grpc"
	api "books-api/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	App                *gin.Engine
	UserHandler        *grpc.UserGRPCHandler
	AuthorHandler      *grpc.AuthorGRPCHandler
	CategoryHandler    *grpc.CategoryGRPCHandler
	BookHandler        *grpc.BookGRPCHandler
	BookLendingHandler *grpc.BookLendingGRPCHandler
	AuthMiddleware     *api.AuthMiddleware
}

func (h *Router) UserSetup(port string) {
	h.App.Use(h.AuthMiddleware.ErrorHandler)
	guestApi := h.App.Group("/api/v1")
	{
		userGRPC := h.GRPCUserHTTPSetup(port)
		guestApi.Any("/*grpc_gateway", gin.WrapH(userGRPC))
	}
}

func (h *Router) BookSetup(port string) {
	bookApi := h.App.Group("/api/v1")
	bookApi.Use(h.AuthMiddleware.ErrorHandler, h.AuthMiddleware.JWTAuthentication)
	{
		bookGRPC := h.GRPCBookHTTPSetup(port)
		bookApi.Any("/*grpc_gateway", gin.WrapH(bookGRPC))
	}
}

func (h *Router) AuthorSetup(port string) {
	authorApi := h.App.Group("/api/v1")
	authorApi.Use(h.AuthMiddleware.ErrorHandler, h.AuthMiddleware.JWTAuthentication)
	{
		authorGRPC := h.GRPCAuthorHTTPSetup(port)
		authorApi.Any("/*grpc_gateway", gin.WrapH(authorGRPC))
	}
}

func (h *Router) CategorySetup(port string) {
	categoryApi := h.App.Group("/api/v1")
	categoryApi.Use(h.AuthMiddleware.ErrorHandler, h.AuthMiddleware.JWTAuthentication)
	{
		categoryGRPC := h.GRPCCategoryHTTPSetup(port)
		categoryApi.Any("/*grpc_gateway", gin.WrapH(categoryGRPC))
	}
}
