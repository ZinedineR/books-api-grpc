package externalapi

import authors "books-api/proto/author/v1"

type AuthorSvcExternal interface {
	GetById(ID string) (*authors.GetAuthorByIDResponse, error)
}
