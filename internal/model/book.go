package model

import (
	"books-api/internal/entity"
	"books-api/proto/books/v1"
	"github.com/google/uuid"
)

type BaseBookReq struct {
	Title      string `json:"title" validate:"required,gte=2" example:"The Great Gatsby"`
	ISBN       string `json:"isbn" validate:"required,gte=2" example:"978-3-16-148410-0"`
	AuthorId   string `json:"author_id" validate:"required,uuid" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	CategoryId string `json:"category_id" validate:"required,uuid" example:"7ea2b005-a3a5-4f4b-b762-75234fb6b4bd"`
	Stock      uint32 `json:"stock" example:"1"`
}

type CreateBookReq struct {
	BaseBookReq
}

func (req BaseBookReq) ToEntity() *entity.Book {
	return &entity.Book{
		Id:         uuid.NewString(),
		Title:      req.Title,
		ISBN:       req.ISBN,
		AuthorId:   req.AuthorId,
		CategoryId: req.CategoryId,
		Stock:      uint(req.Stock),
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

func BookGRPCToEntity(input *books.Book) *entity.Book {
	return &entity.Book{
		Id:         input.Id,
		Title:      input.Title,
		ISBN:       input.Isbn,
		AuthorId:   input.AuthorId,
		CategoryId: input.CategoryId,
		Stock:      uint(input.Stock),
	}
}

func BookListGRPCToEntity(list []*books.Book) []*entity.Book {
	var results []*entity.Book
	if len(list) > 0 {
		for _, book := range list {
			results = append(results, BookGRPCToEntity(book))
		}
	}
	return results
}
