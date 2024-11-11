package service

import (
	"books-api/internal/model"
	"books-api/pkg/exception"
	"context"
)

type UserService interface {
	// CRUD operations for Book
	Register(
		ctx context.Context, req *model.CreateUserReq,
	) (*model.CreateUserRes, *exception.Exception)
	Login(ctx context.Context, req *model.CreateUserReq) (*model.LoginUserRes, *exception.Exception)
	Find(ctx context.Context, req *model.GetAllUserReq) (*model.GetAllUserRes, *exception.Exception)
	Detail(ctx context.Context, req *model.GetUserByIDReq) (*model.GetUserByIDRes, *exception.Exception)
}
