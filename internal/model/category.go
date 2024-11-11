package model

import (
	"books-api/internal/entity"
	"github.com/google/uuid"
)

type BaseCategoryReq struct {
	Title string `json:"title" validate:"required,gte=2" example:"Romance"`
}

type CreateCategoryReq struct {
	BaseCategoryReq
}

func (req BaseCategoryReq) ToEntity() *entity.Category {
	return &entity.Category{
		Id:    uuid.NewString(),
		Title: req.Title,
	}
}

type CreateCategoryRes struct {
	entity.Category
}

type UpdateCategoryReq struct {
	BaseCategoryReq
	ID string `swaggerignore:"true"`
}
type UpdateCategoryRes struct {
	entity.Category
}

type DeleteCategoryReq struct {
	ID string `swaggerignore:"true"`
}
type DeleteCategoryRes struct {
	ID string `swaggerignore:"true"`
}

type GetAllCategoryReq struct {
	Page   PaginationParam
	Filter FilterParams
	Sort   OrderParam
}
type GetAllCategoryRes struct {
	PaginationData[entity.Category]
}

type GetCategoryByIDReq struct {
	ID string `swaggerignore:"true"`
}

type GetCategoryByIDRes struct {
	entity.Category
	Books []*entity.Book
}
