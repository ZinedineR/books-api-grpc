package service

import (
	"books-api/internal/model"
	"books-api/pkg/exception"
	"context"
)

type BookLendingService interface {
	// CRUD operations for BookLending
	Create(
		ctx context.Context, req *model.CreateBookLendingReq,
	) (*model.CreateBookLendingRes, *exception.Exception)
	Return(
		ctx context.Context, req *model.UpdateBookLendingReq,
	) (*model.UpdateBookLendingRes, *exception.Exception)
	Extend(
		ctx context.Context, req *model.UpdateBookLendingReq,
	) (*model.UpdateBookLendingRes, *exception.Exception)
	Find(ctx context.Context, req *model.GetAllBookLendingReq) (*model.GetAllBookLendingRes, *exception.Exception)
	Detail(ctx context.Context, req *model.GetBookLendingByIDReq) (*model.GetBookLendingByIDRes, *exception.Exception)
}
