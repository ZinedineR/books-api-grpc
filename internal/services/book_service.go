package service

import (
	"books-api/internal/model"
	"books-api/internal/repository"
	"context"
	//"books-api/pkg/exception"
	"books-api/pkg/exception"
	"books-api/pkg/xvalidator"
	"gorm.io/gorm"
)

type BookServiceImpl struct {
	db       *gorm.DB
	bookRepo repository.BookRepository
	validate *xvalidator.Validator
}

func NewBookService(
	db *gorm.DB, repo repository.BookRepository,
	validate *xvalidator.Validator,
) BookService {
	return &BookServiceImpl{
		db:       db,
		bookRepo: repo,
		validate: validate,
	}
}

func (s *BookServiceImpl) Create(
	ctx context.Context, req *model.CreateBookReq,
) (*model.CreateBookRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	if err := s.bookRepo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("err", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.CreateBookRes{
		Book: *body,
	}, nil
}

func (s *BookServiceImpl) Update(
	ctx context.Context, req *model.UpdateBookReq,
) (*model.UpdateBookRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	body.Id = req.ID
	if err := s.bookRepo.UpdateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("err", err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.UpdateBookRes{
		Book: *body,
	}, nil
}

func (s *BookServiceImpl) Delete(ctx context.Context, req *model.DeleteBookReq) (
	*model.DeleteBookRes, *exception.Exception,
) {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.bookRepo.DeleteByIDTx(ctx, tx, req.ID); err != nil {
		return nil, exception.Internal("err", err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.DeleteBookRes{
		ID: req.ID,
	}, nil
}

func (s *BookServiceImpl) Find(ctx context.Context, req *model.GetAllBookReq) (
	*model.GetAllBookRes, *exception.Exception,
) {
	result, err := s.bookRepo.FindByPagination(ctx, s.db, req.Page, req.Sort, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get book", err)
	}
	return &model.GetAllBookRes{
		PaginationData: *result,
	}, nil
}

func (s *BookServiceImpl) Detail(ctx context.Context, req *model.GetBookByIDReq) (
	*model.GetBookByIDRes, *exception.Exception,
) {
	result, err := s.bookRepo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal("err", err)
	}
	return &model.GetBookByIDRes{
		Book: *result,
	}, nil
}
