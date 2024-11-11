package externalapi

import (
	authors "books-api/proto/author/v1"
	"context"
)

type AuthorSvcExternal interface {
	GetById(ctx context.Context, ID string) (*authors.GetAuthorByIDResponse, error)
}
