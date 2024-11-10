package grpc

import (
	"books-api/internal/model"
	"books-api/internal/services"
	"books-api/proto/books/v1"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookGRPCHandler struct {
	books.UnimplementedBookServiceServer
	BookService service.BookService
	GRPCParamHandler
}

func NewBookGRPCHandler(service service.BookService) *BookGRPCHandler {
	return &BookGRPCHandler{BookService: service}
}

// Create handles the gRPC request for creating a new book.
func (h *BookGRPCHandler) Create(ctx context.Context, req *books.CreateBookRequest) (*books.CreateBookResponse, error) {
	book := &model.CreateBookReq{
		BaseBookReq: model.BaseBookReq{
			Title:    req.GetTitle(),
			ISBN:     req.GetIsbn(),
			AuthorId: req.GetAuthorId(),
		},
	}
	response, err := h.BookService.Create(ctx, book)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &books.CreateBookResponse{
		Book: &books.Book{
			Id:       response.Id,
			Title:    response.Title,
			Isbn:     response.ISBN,
			AuthorId: response.AuthorId,
		},
		Response: &books.MutationResponse{
			Message: "Book created successfully",
		},
	}, nil
}

// Update handles the gRPC request for updating an existing book.
func (h *BookGRPCHandler) Update(ctx context.Context, req *books.UpdateBookRequest) (*books.UpdateBookResponse, error) {
	book := &model.UpdateBookReq{
		BaseBookReq: model.BaseBookReq{
			Title:    req.GetTitle(),
			ISBN:     req.GetIsbn(),
			AuthorId: req.GetAuthorId(),
		},
		ID: req.GetId(),
	}
	response, err := h.BookService.Update(ctx, book)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &books.UpdateBookResponse{
		Book: &books.Book{
			Id:       response.Id,
			Title:    response.Title,
			Isbn:     response.ISBN,
			AuthorId: response.AuthorId,
		},
		Response: &books.MutationResponse{
			Message: "Book updated successfully",
		},
	}, nil
}

// Find handles the gRPC request for retrieving a list of books.
func (h *BookGRPCHandler) Find(ctx context.Context, req *books.GetAllBookRequest) (*books.GetAllBookResponse, error) {
	var request model.ListReq
	var errParse error
	request.Page, request.Order, request.Filter, errParse = h.ParseFindParams(req.GetPage(), req.GetPageSize(), req.GetSort(), req.GetFilter())
	if errParse != nil {
		return nil, status.Error(codes.PermissionDenied, errParse.Error())
	}
	findReq := &model.GetAllBookReq{
		Page:   request.Page,
		Filter: request.Filter,
		Sort:   request.Order,
	}
	response, err := h.BookService.Find(ctx, findReq)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}

	// Map the response data to the proto response
	booksData := make([]*books.Book, len(response.Data))
	for i, book := range response.Data {
		booksData[i] = &books.Book{
			Id:       book.Id,
			Title:    book.Title,
			Isbn:     book.ISBN,
			AuthorId: book.AuthorId,
		}
	}
	return &books.GetAllBookResponse{
		Pagination: &books.PaginationResponse{
			PageSize:   int32(response.PageSize),
			Page:       int32(response.Page),
			TotalRows:  response.TotalData,
			TotalPages: response.TotalPage,
		},
		Books: booksData,
		Response: &books.MutationResponse{
			Message: "Books retrieved successfully",
		},
	}, nil
}

// Detail handles the gRPC request for retrieving a single book by ID.
func (h *BookGRPCHandler) Detail(ctx context.Context, req *books.GetBookByIDRequest) (
	*books.GetBookByIDResponse, error,
) {
	response, err := h.BookService.Detail(ctx, &model.GetBookByIDReq{ID: req.GetId()})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &books.GetBookByIDResponse{
		Book: &books.Book{
			Id:       response.Id,
			Title:    response.Title,
			Isbn:     response.ISBN,
			AuthorId: response.AuthorId,
		},
	}, nil
}

// Delete handles the gRPC request for deleting a book by ID.
func (h *BookGRPCHandler) Delete(ctx context.Context, req *books.DeleteBookRequest) (*books.DeleteBookResponse, error) {
	_, err := h.BookService.Delete(ctx, &model.DeleteBookReq{ID: req.GetId()})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &books.DeleteBookResponse{
		Response: &books.MutationResponse{
			Message: "Book deleted successfully",
		},
	}, nil
}
