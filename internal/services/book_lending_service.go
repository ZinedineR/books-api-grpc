package service

import (
	"books-api/internal/gateway/externalapi"
	"books-api/internal/model"
	"books-api/internal/repository"
	"context"
	//"books-api/pkg/exception"
	"books-api/pkg/exception"
	"books-api/pkg/xvalidator"
	"gorm.io/gorm"
)

type BookLendingServiceImpl struct {
	db              *gorm.DB
	bookLendingRepo repository.BookLendingRepository
	bookRepo        repository.BookRepository
	userService     externalapi.UserSvcExternal
	validate        *xvalidator.Validator
}

func NewBookLendingService(
	db *gorm.DB, repo repository.BookLendingRepository,
	bookRepo repository.BookRepository,
	userService externalapi.UserSvcExternal,
	validate *xvalidator.Validator,
) BookLendingService {
	return &BookLendingServiceImpl{
		db:              db,
		bookLendingRepo: repo,
		bookRepo:        bookRepo,
		userService:     userService,
		validate:        validate,
	}
}

func (s *BookLendingServiceImpl) Create(
	ctx context.Context, req *model.CreateBookLendingReq,
) (*model.CreateBookLendingRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	userCheck, err := s.userService.GetById(ctx, req.UserId)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if userCheck == nil {
		return nil, exception.NotFound("user not found, id: " + req.UserId)
	}
	bookCheck, err := s.bookRepo.FindByID(ctx, s.db, req.BookId)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if bookCheck == nil {
		return nil, exception.NotFound("book not found, id: " + req.BookId)
	}
	if bookCheck.Stock < 1 {
		return nil, exception.NotFound("book stock is 0: id" + req.BookId)
	}
	if err := s.bookLendingRepo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	//update stock of book
	bookCheck.Stock -= 1
	if err := s.bookRepo.UpdateTx(ctx, tx, bookCheck); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.CreateBookLendingRes{
		BookLending: *body,
	}, nil
}

func (s *BookLendingServiceImpl) Return(
	ctx context.Context, req *model.UpdateBookLendingReq,
) (*model.UpdateBookLendingRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	bookLending, err := s.bookLendingRepo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if bookLending == nil {
		return nil, exception.NotFound("book lending not found, id: " + req.ID)
	}
	bookCheck, err := s.bookRepo.FindByID(ctx, s.db, bookLending.BookId)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if bookCheck == nil {
		return nil, exception.NotFound("book not found, id: " + req.ID)
	}
	bookLending.Returned = true
	if err := s.bookLendingRepo.UpdateTx(ctx, tx, bookLending); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	//update stock of book
	bookCheck.Stock += 1
	if err := s.bookRepo.UpdateTx(ctx, tx, bookCheck); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.UpdateBookLendingRes{
		BookLending: *bookLending,
	}, nil
}

func (s *BookLendingServiceImpl) Extend(
	ctx context.Context, req *model.UpdateBookLendingReq,
) (*model.UpdateBookLendingRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	bookLending, err := s.bookLendingRepo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if bookLending == nil {
		return nil, exception.NotFound("book lending not found, id: " + req.ID)
	}
	bookLending.EndDate = bookLending.EndDate.AddDate(0, 0, 7)
	if err := s.bookLendingRepo.UpdateTx(ctx, tx, bookLending); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.UpdateBookLendingRes{
		BookLending: *bookLending,
	}, nil
}

func (s *BookLendingServiceImpl) Find(ctx context.Context, req *model.GetAllBookLendingReq) (
	*model.GetAllBookLendingRes, *exception.Exception,
) {
	result, err := s.bookLendingRepo.FindByPagination(ctx, s.db, req.Page, req.Sort, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get book", err)
	}
	return &model.GetAllBookLendingRes{
		PaginationData: *result,
	}, nil
}

func (s *BookLendingServiceImpl) Detail(ctx context.Context, req *model.GetBookLendingByIDReq) (
	*model.GetBookLendingByIDRes, *exception.Exception,
) {
	result, err := s.bookLendingRepo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if result == nil {
		return nil, exception.NotFound("book lending not found, id: " + req.ID)
	}
	return &model.GetBookLendingByIDRes{
		BookLending: *result,
	}, nil
}
