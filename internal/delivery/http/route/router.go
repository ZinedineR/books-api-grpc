package route

import (
	"books-api/internal/delivery/grpc"
	api "books-api/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	App            *gin.Engine
	UserHandler    *grpc.UserGRPCHandler
	BookHandler    *grpc.BookGRPCHandler
	AuthMiddleware *api.AuthMiddleware
}

func (h *Router) UserSetup() {
	h.App.Use(h.AuthMiddleware.ErrorHandler)
	guestApi := h.App.Group("/api/v1/auth")
	{
		userGRPC := h.GRPCUserSetup()
		guestApi.Any("/*grpc_gateway", gin.WrapH(userGRPC))
	}
}

func (h *Router) BookSetup() {
	h.App.Use(h.AuthMiddleware.ErrorHandler)
	bookApi := h.App.Group("/api/v1")
	bookApi.Use(h.AuthMiddleware.JWTAuthentication)
	{
		bookGRPC := h.GRPCBookSetup()
		bookApi.Any("/*grpc_gateway", gin.WrapH(bookGRPC))
	}

}
