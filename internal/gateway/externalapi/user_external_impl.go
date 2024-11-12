package externalapi

import (
	"books-api/config"
	"books-api/pkg/server"
	"books-api/proto/users/v1"
	"context"
)

type UserExternalImpl struct {
	config *config.Config
	users  users.UserServiceClient
}

func NewUserExternalImpl(
	config *config.Config,
) UserSvcExternal {
	author := users.NewUserServiceClient(server.NewGRPCListenerClient(config.AppEnvConfig.UserListenerHost, config.AppEnvConfig.UserGRPCPort))
	return &UserExternalImpl{
		config: config,
		users:  author,
	}
}
func (b *UserExternalImpl) GetById(ctx context.Context, ID string) (*users.GetUserByIDResponse, error) {

	response, err := b.users.Detail(ctx, &users.GetUserByIDRequest{
		Id: ID,
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}
