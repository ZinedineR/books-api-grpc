package model

import "books-api/internal/entity"

type BaseUserReq struct {
	Username string `json:"username" validate:"required" example:"john_doe"`
	Password string `json:"password" validate:"required,password,gte=8" example:"SecurePass123!"`
}

type CreateUserReq struct {
	BaseUserReq
}

func (req BaseUserReq) ToEntity() *entity.User {
	return &entity.User{
		Username: req.Username,
		Password: req.Password,
	}
}

type CreateUserRes struct {
	entity.User
}

type UpdateUserReq struct {
	BaseUserReq
	ID string `json:"id" name:"id"`
}
type UpdateUserRes struct {
	entity.User
}

type DeleteUserReq struct {
	BaseUserReq
	ID string `json:"id" name:"id"`
}
type DeleteUserRes struct {
	ID string `json:"id" name:"id"`
}

type GetAllUserReq struct {
	Page   PaginationParam
	Filter FilterParams
	Sort   OrderParam
}
type GetAllUserRes struct {
	PaginationData[entity.User]
}

type GetUserByIDReq struct {
	ID string `json:"id" name:"id"`
}

type GetUserByIDRes struct {
	entity.User
}
