package service

import (
	"books-api/internal/entity"
	"books-api/internal/model"
	"books-api/pkg/exception"
	"context"
)

type CategoryService interface {
	// CRUD operations for Category
	Create(
		ctx context.Context, req *model.CreateCategoryReq,
	) (*model.CreateCategoryRes, *exception.Exception)
	Update(
		ctx context.Context, req *model.UpdateCategoryReq,
	) (*model.UpdateCategoryRes, *exception.Exception)
	Find(ctx context.Context, req *model.GetAllCategoryReq) (*model.GetAllCategoryRes, *exception.Exception)
	Detail(ctx context.Context, req *model.GetCategoryByIDReq) (*model.GetCategoryByIDRes, *exception.Exception)
	Delete(ctx context.Context, req *model.DeleteCategoryReq) (*model.DeleteCategoryRes, *exception.Exception)
}

type ListCategoryResp struct {
	Pagination *model.Pagination  `json:"pagination"`
	Data       []*entity.Category `json:"data"`
}
