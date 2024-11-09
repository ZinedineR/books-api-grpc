package model

import (
	"books-api/internal/entity"
	"github.com/google/uuid"
)

type BaseBookReq struct {
	Title    string `json:"title" validate:"required,gte=2" example:"The Great Gatsby"`
	ISBN     string `json:"isbn" validate:"required,gte=2" example:"978-3-16-148410-0"`
	AuthorId string `json:"author_id" validate:"required,uuid" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
}

type CreateBookReq struct {
	BaseBookReq
}

func (req BaseBookReq) ToEntity() *entity.Book {
	return &entity.Book{
		Id:       uuid.NewString(),
		Title:    req.Title,
		ISBN:     req.ISBN,
		AuthorId: req.AuthorId,
	}
}

type CreateBookRes struct {
	entity.Book
}

type UpdateBookReq struct {
	BaseBookReq
	ID string `swaggerignore:"true"`
}
type UpdateBookRes struct {
	entity.Book
}

type DeleteBookReq struct {
	ID string `swaggerignore:"true"`
}
type DeleteBookRes struct {
	ID string `swaggerignore:"true"`
}

type GetAllBookReq struct {
	Page   PaginationParam
	Filter FilterParams
	Sort   OrderParam
}
type GetAllBookRes struct {
	PaginationData[entity.Book]
}

type GetBookByIDReq struct {
	ID string `swaggerignore:"true"`
}

type GetBookByIDRes struct {
	entity.Book
}
