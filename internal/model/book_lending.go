package model

import (
	"books-api/internal/entity"
	"github.com/google/uuid"
	"time"
)

type BaseBookLendingReq struct {
	UserId string `json:"user_id" gorm:"uuid" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	BookId string `json:"book_id" gorm:"uuid" example:"7ea2b005-a3a5-4f4b-b762-75234fb6b4bd"`
}

type CreateBookLendingReq struct {
	BaseBookLendingReq
}

func (req BaseBookLendingReq) ToEntity() *entity.BookLending {
	now := time.Now()
	return &entity.BookLending{
		Id:        uuid.NewString(),
		UserId:    req.UserId,
		BookId:    req.BookId,
		StartDate: now,
		EndDate:   now.AddDate(0, 0, 7),
	}
}

type CreateBookLendingRes struct {
	entity.BookLending
}

type UpdateBookLendingReq struct {
	ID string `swaggerignore:"true"`
}
type UpdateBookLendingRes struct {
	entity.BookLending
}

type GetAllBookLendingReq struct {
	Page   PaginationParam
	Filter FilterParams
	Sort   OrderParam
}
type GetAllBookLendingRes struct {
	PaginationData[entity.BookLending]
}

type GetBookLendingByIDReq struct {
	ID string `swaggerignore:"true"`
}

type GetBookLendingByIDRes struct {
	entity.BookLending
}

//func BookLendingGRPCToEntity(input *books.BookLending) *entity.BookLending {
//	return &entity.BookLending{
//		Id:         input.Id,
//		Title:      input.Title,
//		ISBN:       input.Isbn,
//		AuthorId:   input.AuthorId,
//		CategoryId: input.CategoryId,
//	}
//}
//
//func BookLendingListGRPCToEntity(list []*books.BookLending) []*entity.BookLending {
//	var results []*entity.BookLending
//	if len(list) > 0 {
//		for _, book := range list {
//			results = append(results, BookLendingGRPCToEntity(book))
//		}
//	}
//	return results
//}
