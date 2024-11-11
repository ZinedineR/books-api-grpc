package externalapi

import (
	categories "books-api/proto/category/v1"
	"context"
)

type CategorySvcExternal interface {
	GetById(ctx context.Context, ID string) (*categories.GetCategoryByIDResponse, error)
}
