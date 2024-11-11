package externalapi

import (
	"books-api/config"
	"books-api/pkg/server"
	"books-api/proto/books/v1"
	"context"
)

type BookExternalImpl struct {
	config *config.Config
	books  books.BookServiceClient
}

func NewBookExternalImpl(
	config *config.Config,
) BookSvcExternal {
	author := books.NewBookServiceClient(server.NewGRPCListenerClient(config.AppEnvConfig.BookGRPCPort))
	return &BookExternalImpl{
		config: config,
		books:  author,
	}
}
func (b *BookExternalImpl) GetById(ctx context.Context, ID string) (*books.GetBookByIDResponse, error) {

	response, err := b.books.Detail(ctx, &books.GetBookByIDRequest{
		Id: ID,
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *BookExternalImpl) Find(ctx context.Context, req *books.GetAllBookRequest) (*books.GetAllBookResponse, error) {

	response, err := b.books.Find(ctx, &books.GetAllBookRequest{
		Filter: req.GetFilter(),
		Sort:   req.GetSort(),
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (b *BookExternalImpl) Update(ctx context.Context, req *books.UpdateBookRequest) (
	*books.UpdateBookResponse, error,
) {
	response, err := b.books.Update(ctx, &books.UpdateBookRequest{
		Id:         req.GetId(),
		Title:      req.GetTitle(),
		Isbn:       req.GetIsbn(),
		AuthorId:   req.GetAuthorId(),
		CategoryId: req.GetCategoryId(),
		Stock:      req.GetStock(),
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}
