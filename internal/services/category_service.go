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

type CategoryServiceImpl struct {
	db           *gorm.DB
	categoryRepo repository.CategoryRepository
	bookService  externalapi.BookSvcExternal
	validate     *xvalidator.Validator
}

func NewCategoryService(
	db *gorm.DB, repo repository.CategoryRepository,
	bookService externalapi.BookSvcExternal,
	validate *xvalidator.Validator,
) CategoryService {
	return &CategoryServiceImpl{
		db:           db,
		categoryRepo: repo,
		bookService:  bookService,
		validate:     validate,
	}
}

func (s *CategoryServiceImpl) Create(
	ctx context.Context, req *model.CreateCategoryReq,
) (*model.CreateCategoryRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	checker, err := s.categoryRepo.FindByFilter(ctx, s.db,
		model.FilterParams{
			&model.FilterParam{
				Field:    "title",
				Value:    req.Title,
				Operator: "=",
			},
		},
		model.OrderParam{
			Order:   "desc",
			OrderBy: "title",
		})
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if checker != nil {
		return nil, exception.PermissionDenied("category with this title already exists")
	}
	body := req.ToEntity()
	if err := s.categoryRepo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.CreateCategoryRes{
		Category: *body,
	}, nil
}

func (s *CategoryServiceImpl) Update(
	ctx context.Context, req *model.UpdateCategoryReq,
) (*model.UpdateCategoryRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	checker, err := s.categoryRepo.FindByFilter(ctx, s.db,
		model.FilterParams{
			&model.FilterParam{
				Field:    "title",
				Value:    req.Title,
				Operator: "=",
			},
		},
		model.OrderParam{
			Order:   "desc",
			OrderBy: "title",
		})
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if checker != nil && checker.Id != body.Id {
		return nil, exception.PermissionDenied("category with this title already exists")
	}
	body.Id = req.ID
	if err := s.categoryRepo.UpdateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.UpdateCategoryRes{
		Category: *body,
	}, nil
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, req *model.DeleteCategoryReq) (
	*model.DeleteCategoryRes, *exception.Exception,
) {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.categoryRepo.DeleteByIDTx(ctx, tx, req.ID); err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	booksResponse, err := s.bookService.Find(ctx, &books.GetAllBookRequest{
		Filter: "category_id:" + req.ID + ":eq",
	})
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	//making category id as null in book
	if len(booksResponse.Books) > 0 {
		for _, book := range booksResponse.Books {
			if _, err := s.bookService.Update(ctx, &books.UpdateBookRequest{
				Id:         book.Id,
				Title:      book.Title,
				Isbn:       book.Isbn,
				AuthorId:   book.Id,
				CategoryId: "",
			}); err != nil {
				return nil, exception.Internal(err.Error(), err)
			}
		}
	}
	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("commit transaction", err)
	}
	return &model.DeleteCategoryRes{
		ID: req.ID,
	}, nil
}

func (s *CategoryServiceImpl) Find(ctx context.Context, req *model.GetAllCategoryReq) (
	*model.GetAllCategoryRes, *exception.Exception,
) {
	result, err := s.categoryRepo.FindByPagination(ctx, s.db, req.Page, req.Sort, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get category", err)
	}
	return &model.GetAllCategoryRes{
		PaginationData: *result,
	}, nil
}

func (s *CategoryServiceImpl) Detail(ctx context.Context, req *model.GetCategoryByIDReq) (
	*model.GetCategoryByIDRes, *exception.Exception,
) {
	result, err := s.categoryRepo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	if result == nil {
		return nil, exception.NotFound("category not found, id: " + req.ID)
	}
	booksResponse, err := s.bookService.Find(ctx, &books.GetAllBookRequest{
		Filter: "category_id:" + req.ID + ":eq",
	})
	if err != nil {
		return nil, exception.Internal(err.Error(), err)
	}
	return &model.GetCategoryByIDRes{
		Category: *result,
		Books:    model.BookListGRPCToEntity(booksResponse.Books),
	}, nil
}
