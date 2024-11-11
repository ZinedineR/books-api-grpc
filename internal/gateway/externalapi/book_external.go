package externalapi

import (
	"books-api/proto/books/v1"
	"context"
)

type BookSvcExternal interface {
	GetById(ctx context.Context, ID string) (*books.GetBookByIDResponse, error)
	Find(ctx context.Context, req *books.GetAllBookRequest) (*books.GetAllBookResponse, error)
	Update(ctx context.Context, req *books.UpdateBookRequest) (*books.UpdateBookResponse, error)
}
