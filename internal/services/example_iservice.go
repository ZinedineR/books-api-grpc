package service

import (
	"boiler-plate-clean/internal/entity"
	"boiler-plate-clean/internal/model"
	"boiler-plate-clean/pkg/exception"
	"context"
)

type ExampleService interface {
	// CRUD operations for Example
	CreateExample(
		ctx context.Context, model *entity.Example,
	) *exception.Exception
}

type ListExampleResp struct {
	Pagination *model.Pagination `json:"pagination"`
	Data       []*entity.Example `json:"data"`
}
