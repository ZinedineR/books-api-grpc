package grpc

type BaseGRPCHandler struct {
	User *UserGRPCHandler
}

func NewBaseGRPCHandler(
	user *UserGRPCHandler,
) *BaseGRPCHandler {
	return &BaseGRPCHandler{
		User: user,
	}
}
