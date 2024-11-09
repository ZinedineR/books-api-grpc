package route

import (
	"books-api/internal/delivery/grpc"
	api "books-api/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	App            *gin.Engine
	GRPCHandler    *grpc.BaseGRPCHandler
	AuthMiddleware *api.AuthMiddleware
}

func (h *Router) Setup() {
	h.App.Use(h.AuthMiddleware.ErrorHandler)
	//h.App.Group("/api/v1/*{grpc_gateway}").Any("", gin.WrapH(h.GRPCUserSetup()))
	guestApi := h.App.Group("/api/v1/auth")
	{
		userGRPC := h.GRPCUserSetup()
		guestApi.Any("/*grpc_gateway", gin.WrapH(userGRPC))
	}
	//signatureAPI := h.App.Group("/auth/signature")
	//signatureAPI.Use(h.AuthMiddleware.JWTAuthentication)
	//{
	//	signatureAPI.POST("", h.SignatureHandler.Signature)
	//}
	//userApi := h.App.Group("")
	//userApi.Use(h.AuthMiddleware.JWTAuthentication, h.AuthMiddleware.SignatureAuthentication)
	//{
	//	booksApi := userApi.Group("/books")
	//	{
	//		booksApi.POST("", h.BookHandler.Create)
	//		booksApi.GET("", h.BookHandler.List)
	//		booksApi.GET("/:id", h.BookHandler.FindOne)
	//		booksApi.PUT("/:id", h.BookHandler.Update)
	//		booksApi.DELETE("/:id", h.BookHandler.Delete)
	//	}
	//	authorApi := userApi.Group("/authors")
	//	{
	//		authorApi.POST("", h.AuthorHandler.Create)
	//		authorApi.GET("", h.AuthorHandler.List)
	//		authorApi.GET("/:id", h.AuthorHandler.FindOne)
	//		authorApi.PUT("/:id", h.AuthorHandler.Update)
	//		authorApi.DELETE("/:id", h.AuthorHandler.Delete)
	//	}
	//}
}
