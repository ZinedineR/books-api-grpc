package service

import (
	"books-api/internal/model"
	"books-api/pkg/exception"
	"context"
)

type BookService interface {
	// CRUD operations for Book
	Create(
		ctx context.Context, req *model.CreateBookReq,
	) (*model.CreateBookRes, *exception.Exception)
	Update(
		ctx context.Context, req *model.UpdateBookReq,
	) (*model.UpdateBookRes, *exception.Exception)
	Find(ctx context.Context, req *model.GetAllBookReq) (*model.GetAllBookRes, *exception.Exception)
	Detail(ctx context.Context, req *model.GetBookByIDReq) (*model.GetBookByIDRes, *exception.Exception)
	Delete(ctx context.Context, req *model.DeleteBookReq) (*model.DeleteBookRes, *exception.Exception)
}
