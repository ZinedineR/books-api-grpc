package externalapi

import (
	"books-api/config"
	"books-api/pkg/server"
	authors "books-api/proto/author/v1"
	"context"
)

type AuthorExternalImpl struct {
	config  *config.Config
	authors authors.AuthorServiceClient
}

func NewAuthorExternalImpl(
	config *config.Config,
) AuthorSvcExternal {
	author := authors.NewAuthorServiceClient(server.NewGRPCListenerClient(config.AppEnvConfig.AuthorGRPCPort))
	return &AuthorExternalImpl{
		config:  config,
		authors: author,
	}
}
func (b *AuthorExternalImpl) GetById(ID string) (*authors.GetAuthorByIDResponse, error) {

	response, err := b.authors.Detail(context.Background(), &authors.GetAuthorByIDRequest{
		Id: ID,
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}
