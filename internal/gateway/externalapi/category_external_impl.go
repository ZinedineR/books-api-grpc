package externalapi

import (
	"books-api/config"
	"books-api/pkg/server"
	categories "books-api/proto/category/v1"
	"context"
)

type CategoryExternalImpl struct {
	config     *config.Config
	categories categories.CategoryServiceClient
}

func NewCategoryExternalImpl(
	config *config.Config,
) CategorySvcExternal {
	author := categories.NewCategoryServiceClient(server.NewGRPCListenerClient(config.AppEnvConfig.CategoryGRPCPort))
	return &CategoryExternalImpl{
		config:     config,
		categories: author,
	}
}
func (b *CategoryExternalImpl) GetById(ctx context.Context, ID string) (*categories.GetCategoryByIDResponse, error) {

	response, err := b.categories.Detail(ctx, &categories.GetCategoryByIDRequest{
		Id: ID,
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}
