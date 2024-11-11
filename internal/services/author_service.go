package service

import (
	"books-api/internal/gateway/externalapi"
	"books-api/internal/model"
	"books-api/internal/repository"
	"books-api/proto/books/v1"
	"context"
	//"books-api/pkg/exception"
	"books-api/pkg/exception"
	"books-api/pkg/xvalidator"
	"gorm.io/gorm"
)

type AuthorServiceImpl struct {
	db          *gorm.DB
	authorRepo  repository.AuthorRepository
	bookService externalapi.BookSvcExternal
	validate    *xvalidator.Validator
}

func NewAuthorService(
	db *gorm.DB, repo repository.AuthorRepository,
	bookService externalapi.BookSvcExternal,
	validate *xvalidator.Validator,
) AuthorService {
	return &AuthorServiceImpl{
		db:          db,
		authorRepo:  repo,
		bookService: bookService,
		validate:    validate,
	}
}

func (s *AuthorServiceImpl) Create(
	ctx context.Context, req *model.CreateAuthorReq,
) (*model.CreateAuthorRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	checker, err := s.authorRepo.FindByFilter(ctx, s.db,
		model.FilterParams{
			&model.FilterParam{
				Field:    "name",
				Value:    req.Name,
				Operator: "=",
			},
		},
		model.OrderParam{
			Order:   "desc",
			OrderBy: "name",
		})
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if checker != nil {
		return nil, exception.PermissionDenied("author with this name already exists")
	}
	body := req.ToEntity()
	if err := s.authorRepo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.CreateAuthorRes{
		Author: *body,
	}, nil
}

func (s *AuthorServiceImpl) Update(
	ctx context.Context, req *model.UpdateAuthorReq,
) (*model.UpdateAuthorRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	checker, err := s.authorRepo.FindByFilter(ctx, s.db,
		model.FilterParams{
			&model.FilterParam{
				Field:    "name",
				Value:    req.Name,
				Operator: "=",
			},
		},
		model.OrderParam{
			Order:   "desc",
			OrderBy: "name",
		})
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if checker != nil && checker.Id != body.Id {
		return nil, exception.PermissionDenied("author with this name already exists")
	}
	body.Id = req.ID
	if err := s.authorRepo.UpdateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.UpdateAuthorRes{
		Author: *body,
	}, nil
}

func (s *AuthorServiceImpl) Delete(ctx context.Context, req *model.DeleteAuthorReq) (
	*model.DeleteAuthorRes, *exception.Exception,
) {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.authorRepo.DeleteByIDTx(ctx, tx, req.ID); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	booksResponse, err := s.bookService.Find(ctx, &books.GetAllBookRequest{
		Filter: "author_id:" + req.ID + ":eq",
	})
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	//making author id as null in book
	if len(booksResponse.Books) > 0 {
		for _, book := range booksResponse.Books {
			if _, err := s.bookService.Update(ctx, &books.UpdateBookRequest{
				Id:       book.Id,
				Title:    book.Title,
				Isbn:     book.Isbn,
				AuthorId: "",
			}); err != nil {
				return nil, exception.Internal(err.Error(), err)
			}
		}
	}
	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.DeleteAuthorRes{
		ID: req.ID,
	}, nil
}

func (s *AuthorServiceImpl) Find(ctx context.Context, req *model.GetAllAuthorReq) (
	*model.GetAllAuthorRes, *exception.Exception,
) {
	result, err := s.authorRepo.FindByPagination(ctx, s.db, req.Page, req.Sort, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get author", err)
	}
	return &model.GetAllAuthorRes{
		PaginationData: *result,
	}, nil
}

func (s *AuthorServiceImpl) Detail(ctx context.Context, req *model.GetAuthorByIDReq) (
	*model.GetAuthorByIDRes, *exception.Exception,
) {
	result, err := s.authorRepo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if result == nil {
		return nil, exception.NotFound("author not found, id: " + req.ID)
	}
	booksResponse, err := s.bookService.Find(ctx, &books.GetAllBookRequest{
		Filter: "author_id:" + req.ID + ":eq",
	})
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	return &model.GetAuthorByIDRes{
		Author: *result,
		Books:  model.BookListGRPCToEntity(booksResponse.Books),
	}, nil
}
