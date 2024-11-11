package grpc

import (
	"books-api/internal/model"
	service "books-api/internal/services"
	authors "books-api/proto/author/v1"
	"books-api/proto/books/v1"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthorGRPCHandler struct {
	authors.UnimplementedAuthorServiceServer
	AuthorService service.AuthorService
	GRPCParamHandler
}

func NewAuthorGRPCHandler(service service.AuthorService) *AuthorGRPCHandler {
	return &AuthorGRPCHandler{AuthorService: service}
}

// Create handles the gRPC request for creating a new author.
func (h *AuthorGRPCHandler) Create(
	ctx context.Context, req *authors.CreateAuthorRequest,
) (*authors.CreateAuthorResponse, error) {
	author := &model.CreateAuthorReq{
		BaseAuthorReq: model.BaseAuthorReq{
			Name:      req.GetName(),
			Birthdate: req.GetBirthdate(),
		},
	}
	response, err := h.AuthorService.Create(ctx, author)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &authors.CreateAuthorResponse{
		Author: &authors.Author{
			Id:        response.Id,
			Name:      response.Name,
			Birthdate: timestamppb.New(*response.Birthdate),
		},
		Response: &authors.MutationResponse{
			Message: "Author created successfully",
		},
	}, nil
}

// Update handles the gRPC request for updating an author.
func (h *AuthorGRPCHandler) Update(
	ctx context.Context, req *authors.UpdateAuthorRequest,
) (*authors.UpdateAuthorResponse, error) {
	author := &model.UpdateAuthorReq{
		BaseAuthorReq: model.BaseAuthorReq{
			Name:      req.GetName(),
			Birthdate: req.GetBirthdate(),
		},
		ID: req.GetId(),
	}
	response, err := h.AuthorService.Update(ctx, author)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &authors.UpdateAuthorResponse{
		Author: &authors.Author{
			Id:        response.Id,
			Name:      response.Name,
			Birthdate: timestamppb.New(*response.Birthdate),
		},
		Response: &authors.MutationResponse{
			Message: "Author updated successfully",
		},
	}, nil
}

// Find handles the gRPC request for retrieving a list of authors.
func (h *AuthorGRPCHandler) Find(ctx context.Context, req *authors.GetAllAuthorRequest) (
	*authors.GetAllAuthorResponse, error,
) {
	var request model.ListReq
	var errParse error
	request.Page, request.Order, request.Filter, errParse = h.ParseFindParams(req.GetPage(), req.GetPageSize(), req.GetSort(), req.GetFilter())
	if errParse != nil {
		return nil, status.Error(codes.PermissionDenied, errParse.Error())
	}
	findReq := &model.GetAllAuthorReq{
		Page:   request.Page,
		Filter: request.Filter,
		Sort:   request.Order,
	}
	response, err := h.AuthorService.Find(ctx, findReq)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}

	// Map the response data to the proto response
	authorsData := make([]*authors.Author, len(response.Data))
	for i, author := range response.Data {
		authorsData[i] = &authors.Author{
			Id:        author.Id,
			Name:      author.Name,
			Birthdate: timestamppb.New(*author.Birthdate),
		}
	}
	return &authors.GetAllAuthorResponse{
		Pagination: &authors.PaginationResponse{
			PageSize:   int32(response.PageSize),
			Page:       int32(response.Page),
			TotalRows:  response.TotalData,
			TotalPages: response.TotalPage,
		},
		Authors: authorsData,
		Response: &authors.MutationResponse{
			Message: "Authors retrieved successfully",
		},
	}, nil
}

// Detail handles the gRPC request for retrieving a specific author by ID.
func (h *AuthorGRPCHandler) Detail(
	ctx context.Context, req *authors.GetAuthorByIDRequest,
) (*authors.GetAuthorByIDResponse, error) {
	author, err := h.AuthorService.Detail(ctx, &model.GetAuthorByIDReq{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	// Map the response data to the proto response
	booksData := make([]*books.Book, len(author.Books))
	for i, book := range author.Books {
		booksData[i] = &books.Book{
			Id:       book.Id,
			Title:    book.Title,
			Isbn:     book.ISBN,
			AuthorId: book.AuthorId,
		}
	}
	return &authors.GetAuthorByIDResponse{
		Author: &authors.Author{
			Id:        author.Id,
			Name:      author.Name,
			Birthdate: timestamppb.New(*author.Birthdate),
		},
		Books: booksData,
	}, nil
}

// Delete handles the gRPC request for deleting an author by ID.
func (h *AuthorGRPCHandler) Delete(
	ctx context.Context, req *authors.DeleteAuthorRequest,
) (*authors.DeleteAuthorResponse, error) {
	_, err := h.AuthorService.Delete(ctx, &model.DeleteAuthorReq{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &authors.DeleteAuthorResponse{
		Response: &authors.MutationResponse{
			Message: "Author deleted successfully",
		},
	}, nil
}
