package grpc

import (
	"books-api/internal/model"
	service "books-api/internal/services"
	"books-api/proto/books/v1"
	categories "books-api/proto/category/v1"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryGRPCHandler struct {
	categories.UnimplementedCategoryServiceServer
	CategoryService service.CategoryService
	GRPCParamHandler
}

func NewCategoryGRPCHandler(service service.CategoryService) *CategoryGRPCHandler {
	return &CategoryGRPCHandler{CategoryService: service}
}

// Create handles the gRPC request for creating a new category.
func (h *CategoryGRPCHandler) Create(
	ctx context.Context, req *categories.CreateCategoryRequest,
) (*categories.CreateCategoryResponse, error) {
	category := &model.CreateCategoryReq{
		BaseCategoryReq: model.BaseCategoryReq{
			Title: req.GetTitle(),
		},
	}
	response, err := h.CategoryService.Create(ctx, category)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &categories.CreateCategoryResponse{
		Category: &categories.Category{
			Id:    response.Id,
			Title: response.Title,
		},
		Response: &categories.MutationResponse{
			Message: "Category created successfully",
		},
	}, nil
}

// Update handles the gRPC request for updating an category.
func (h *CategoryGRPCHandler) Update(
	ctx context.Context, req *categories.UpdateCategoryRequest,
) (*categories.UpdateCategoryResponse, error) {
	category := &model.UpdateCategoryReq{
		BaseCategoryReq: model.BaseCategoryReq{
			Title: req.GetTitle(),
		},
		ID: req.GetId(),
	}
	response, err := h.CategoryService.Update(ctx, category)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &categories.UpdateCategoryResponse{
		Category: &categories.Category{
			Id:    response.Id,
			Title: response.Id,
		},
		Response: &categories.MutationResponse{
			Message: "Category updated successfully",
		},
	}, nil
}

// Find handles the gRPC request for retrieving a list of categories.
func (h *CategoryGRPCHandler) Find(ctx context.Context, req *categories.GetAllCategoryRequest) (
	*categories.GetAllCategoryResponse, error,
) {
	var request model.ListReq
	var errParse error
	request.Page, request.Order, request.Filter, errParse = h.ParseFindParams(req.GetPage(), req.GetPageSize(), req.GetSort(), req.GetFilter())
	if errParse != nil {
		return nil, status.Error(codes.PermissionDenied, errParse.Error())
	}
	findReq := &model.GetAllCategoryReq{
		Page:   request.Page,
		Filter: request.Filter,
		Sort:   request.Order,
	}
	response, err := h.CategoryService.Find(ctx, findReq)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}

	// Map the response data to the proto response
	categorysData := make([]*categories.Category, len(response.Data))
	for i, category := range response.Data {
		categorysData[i] = &categories.Category{
			Id:    category.Id,
			Title: category.Title,
		}
	}
	return &categories.GetAllCategoryResponse{
		Pagination: &categories.PaginationResponse{
			PageSize:   int32(response.PageSize),
			Page:       int32(response.Page),
			TotalRows:  response.TotalData,
			TotalPages: response.TotalPage,
		},
		Categories: categorysData,
		Response: &categories.MutationResponse{
			Message: "Categorys retrieved successfully",
		},
	}, nil
}

// Detail handles the gRPC request for retrieving a specific category by ID.
func (h *CategoryGRPCHandler) Detail(
	ctx context.Context, req *categories.GetCategoryByIDRequest,
) (*categories.GetCategoryByIDResponse, error) {
	category, err := h.CategoryService.Detail(ctx, &model.GetCategoryByIDReq{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	// Map the response data to the proto response
	booksData := make([]*books.Book, len(category.Books))
	for i, book := range category.Books {
		booksData[i] = &books.Book{
			Id:         book.Id,
			Title:      book.Title,
			Isbn:       book.ISBN,
			AuthorId:   book.AuthorId,
			CategoryId: book.CategoryId,
		}
	}
	return &categories.GetCategoryByIDResponse{
		Category: &categories.Category{
			Id:    category.Id,
			Title: category.Title,
		},
		Books: booksData,
	}, nil
}

// Delete handles the gRPC request for deleting an category by ID.
func (h *CategoryGRPCHandler) Delete(
	ctx context.Context, req *categories.DeleteCategoryRequest,
) (*categories.DeleteCategoryResponse, error) {
	_, err := h.CategoryService.Delete(ctx, &model.DeleteCategoryReq{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &categories.DeleteCategoryResponse{
		Response: &categories.MutationResponse{
			Message: "Category deleted successfully",
		},
	}, nil
}
