package grpc

import (
	"books-api/internal/model"
	"books-api/internal/services"
	"books-api/proto/users/v1"

	//wallet_finance "books-api/proto/wallet-finance/v1"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserGRPCHandler struct {
	users.UnimplementedUserServiceServer
	UserService service.UserService
	GRPCParamHandler
}

func NewUserGRPCHandler(service service.UserService) *UserGRPCHandler {
	return &UserGRPCHandler{UserService: service}
}

func (h *UserGRPCHandler) Register(ctx context.Context, req *users.RegisterRequest) (
	*users.RegisterResponse, error,
) {
	user := &model.CreateUserReq{
		BaseUserReq: model.BaseUserReq{
			Username: req.GetUsername(),
			Password: req.GetPassword(),
		},
	}
	response, err := h.UserService.Register(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &users.RegisterResponse{
		User: &users.User{
			Id:       response.Id,
			Username: response.Username,
			Password: response.Password,
		},
		Response: &users.MutationResponse{
			Message: "Register User Success",
		},
	}, nil
}

func (h *UserGRPCHandler) Login(ctx context.Context, req *users.LoginRequest) (*users.LoginResponse, error) {
	user := &model.CreateUserReq{
		BaseUserReq: model.BaseUserReq{
			Username: req.GetUsername(),
			Password: req.GetPassword(),
		},
	}
	response, err := h.UserService.Login(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &users.LoginResponse{
		Username: response.Username,
		Token:    response.Token,
	}, nil
}
