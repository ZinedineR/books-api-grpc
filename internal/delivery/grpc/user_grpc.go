package grpc

import (
	"books-api/internal/model"
	"books-api/internal/services"
	"books-api/proto/users/v1"
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

// Find handles the gRPC request for retrieving a list of users with pagination and filtering.
func (h *UserGRPCHandler) Find(ctx context.Context, req *users.GetAllUserRequest) (*users.GetAllUserResponse, error) {
	// Parse pagination, sorting, and filtering parameters
	var request model.ListReq
	var errParse error
	request.Page, request.Order, request.Filter, errParse = h.ParseFindParams(req.GetPage(), req.GetPageSize(), req.GetSort(), req.GetFilter())
	if errParse != nil {
		return nil, status.Error(codes.PermissionDenied, errParse.Error())
	}

	findReq := &model.GetAllUserReq{
		Page:   request.Page,
		Filter: request.Filter,
		Sort:   request.Order,
	}

	// Call the service layer
	response, err := h.UserService.Find(ctx, findReq)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}

	// Map service response to gRPC response
	usersData := make([]*users.User, len(response.Data))
	for i, user := range response.Data {
		usersData[i] = &users.User{
			Id:       user.Id,
			Username: user.Username,
		}
	}

	return &users.GetAllUserResponse{
		Pagination: &users.PaginationResponse{
			PageSize:   int32(response.PageSize),
			Page:       int32(response.Page),
			TotalRows:  response.TotalData,
			TotalPages: response.TotalPage,
		},
		Users: usersData,
		Response: &users.MutationResponse{
			Message: "Users retrieved successfully",
		},
	}, nil
}

// Detail handles the gRPC request for retrieving a single user by ID.
func (h *UserGRPCHandler) Detail(ctx context.Context, req *users.GetUserByIDRequest) (
	*users.GetUserByIDResponse, error,
) {
	// Call the service layer with the provided user ID
	response, err := h.UserService.Detail(ctx, &model.GetUserByIDReq{ID: req.GetId()})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}

	// Map service response to gRPC response
	return &users.GetUserByIDResponse{
		User: &users.User{
			Id:       response.Id,
			Username: response.Username,
		},
	}, nil
}
