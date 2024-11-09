package service

import (
	"books-api/internal/entity"
	"books-api/internal/model"
	"books-api/pkg/exception"
	"context"
)

type AuthorService interface {
	// CRUD operations for Author
	Create(
		ctx context.Context, req *model.CreateAuthorReq,
	) (*model.CreateAuthorRes, *exception.Exception)
	Update(
		ctx context.Context, req *model.UpdateAuthorReq,
	) (*model.UpdateAuthorRes, *exception.Exception)
	Find(ctx context.Context, req *model.GetAllAuthorReq) (*model.GetAllAuthorRes, *exception.Exception)
	Detail(ctx context.Context, req *model.GetAuthorByIDReq) (*model.GetAuthorByIDRes, *exception.Exception)
	Delete(ctx context.Context, req *model.DeleteAuthorReq) (*model.DeleteAuthorRes, *exception.Exception)
}

type ListAuthorResp struct {
	Pagination *model.Pagination `json:"pagination"`
	Data       []*entity.Author  `json:"data"`
}
