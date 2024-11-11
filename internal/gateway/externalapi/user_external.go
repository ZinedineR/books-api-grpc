package externalapi

import (
	"books-api/proto/users/v1"
	"context"
)

type UserSvcExternal interface {
	GetById(ctx context.Context, ID string) (*users.GetUserByIDResponse, error)
}
