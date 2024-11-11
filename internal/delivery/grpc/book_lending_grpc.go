package grpc

import (
	"books-api/internal/model"
	service "books-api/internal/services"
	booklending "books-api/proto/books_lending/v1"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookLendingGRPCHandler struct {
	booklending.UnimplementedBookLendingServiceServer
	BookLendingService service.BookLendingService
	GRPCParamHandler
}

func NewBookLendingGRPCHandler(service service.BookLendingService) *BookLendingGRPCHandler {
	return &BookLendingGRPCHandler{BookLendingService: service}
}

// Create handles the gRPC request for creating a new book lending record.
func (h *BookLendingGRPCHandler) Create(
	ctx context.Context, req *booklending.CreateBookLendingRequest,
) (*booklending.CreateBookLendingResponse, error) {
	lending := &model.CreateBookLendingReq{
		BaseBookLendingReq: model.BaseBookLendingReq{
			UserId: req.GetUserId(),
			BookId: req.GetBookId(),
		},
	}
	response, err := h.BookLendingService.Create(ctx, lending)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &booklending.CreateBookLendingResponse{
		BookLending: &booklending.BookLending{
			Id:        response.Id,
			UserId:    response.UserId,
			BookId:    response.BookId,
			Returned:  response.Returned,
			StartDate: timestamppb.New(response.StartDate),
			EndDate:   timestamppb.New(response.EndDate),
		},
	}, nil
}

// Return handles the gRPC request for marking a book lending as returned.
func (h *BookLendingGRPCHandler) Return(
	ctx context.Context, req *booklending.ReturnBookLendingRequest,
) (*booklending.UpdateBookLendingResponse, error) {
	returnReq := &model.UpdateBookLendingReq{ID: req.GetId()}
	response, err := h.BookLendingService.Return(ctx, returnReq)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &booklending.UpdateBookLendingResponse{
		BookLending: &booklending.BookLending{
			Id:        response.Id,
			UserId:    response.UserId,
			BookId:    response.BookId,
			Returned:  response.Returned,
			StartDate: timestamppb.New(response.StartDate),
			EndDate:   timestamppb.New(response.EndDate),
		},
	}, nil
}

// Extend handles the gRPC request for extending a book lending period.
func (h *BookLendingGRPCHandler) Extend(
	ctx context.Context, req *booklending.ExtendBookLendingRequest,
) (*booklending.UpdateBookLendingResponse, error) {
	extendReq := &model.UpdateBookLendingReq{
		ID: req.GetId(),
	}
	response, err := h.BookLendingService.Extend(ctx, extendReq)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &booklending.UpdateBookLendingResponse{
		BookLending: &booklending.BookLending{
			Id:        response.Id,
			UserId:    response.UserId,
			BookId:    response.BookId,
			Returned:  response.Returned,
			StartDate: timestamppb.New(response.StartDate),
			EndDate:   timestamppb.New(response.EndDate),
		},
	}, nil
}

// Find handles the gRPC request for retrieving a list of book lendings.
func (h *BookLendingGRPCHandler) Find(
	ctx context.Context, req *booklending.GetAllBookLendingRequest,
) (*booklending.GetAllBookLendingResponse, error) {
	var request model.ListReq
	var errParse error
	request.Page, request.Order, request.Filter, errParse = h.ParseFindParams(req.GetPage(), req.GetPageSize(), req.GetSort(), req.GetFilter())
	if errParse != nil {
		return nil, status.Error(codes.PermissionDenied, errParse.Error())
	}

	findReq := &model.GetAllBookLendingReq{
		Page:   request.Page,
		Filter: request.Filter,
		Sort:   request.Order,
	}

	response, err := h.BookLendingService.Find(ctx, findReq)
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}

	// Map the response data to the proto response
	lendingsData := make([]*booklending.BookLending, len(response.Data))
	for i, lending := range response.Data {
		lendingsData[i] = &booklending.BookLending{
			Id:        lending.Id,
			UserId:    lending.UserId,
			BookId:    lending.BookId,
			Returned:  lending.Returned,
			StartDate: timestamppb.New(lending.StartDate),
			EndDate:   timestamppb.New(lending.EndDate),
		}
	}

	return &booklending.GetAllBookLendingResponse{
		Pagination: &booklending.PaginationResponse{
			PageSize:   int32(response.PageSize),
			Page:       int32(response.Page),
			TotalRows:  response.TotalData,
			TotalPages: response.TotalPage,
		},
		BookLendings: lendingsData,
	}, nil
}

// Detail handles the gRPC request for retrieving details of a specific book lending by ID.
func (h *BookLendingGRPCHandler) Detail(
	ctx context.Context, req *booklending.GetBookLendingByIDRequest,
) (*booklending.GetBookLendingByIDResponse, error) {
	response, err := h.BookLendingService.Detail(ctx, &model.GetBookLendingByIDReq{ID: req.GetId()})
	if err != nil {
		return nil, status.Error(codes.Code(err.GetGrpcCode()), fmt.Sprint(err.Message))
	}
	return &booklending.GetBookLendingByIDResponse{
		BookLending: &booklending.BookLending{
			Id:        response.Id,
			UserId:    response.UserId,
			BookId:    response.BookId,
			Returned:  response.Returned,
			StartDate: timestamppb.New(response.StartDate),
			EndDate:   timestamppb.New(response.EndDate),
		},
	}, nil
}
