package model

import (
	"books-api/internal/entity"
	"github.com/google/uuid"
	"time"
)

type BaseAuthorReq struct {
	Name      string `json:"name" validate:"required,gte=2" example:"Ernest Hemingway"`
	Birthdate string `json:"birthdate" validate:"required,dateLocal" example:"1899-07-21"`
}

type CreateAuthorReq struct {
	BaseAuthorReq
}

func (req BaseAuthorReq) ToEntity() *entity.Author {
	birthDate, _ := time.Parse("2006-01-02", req.Birthdate)
	localBirthdate := birthDate.Local()
	return &entity.Author{
		Id:        uuid.NewString(),
		Name:      req.Name,
		Birthdate: &localBirthdate,
	}
}

type CreateAuthorRes struct {
	entity.Author
}

type UpdateAuthorReq struct {
	BaseAuthorReq
	ID string `swaggerignore:"true"`
}
type UpdateAuthorRes struct {
	entity.Author
}

type DeleteAuthorReq struct {
	ID string `swaggerignore:"true"`
}
type DeleteAuthorRes struct {
	ID string `swaggerignore:"true"`
}

type GetAllAuthorReq struct {
	Page   PaginationParam
	Filter FilterParams
	Sort   OrderParam
}
type GetAllAuthorRes struct {
	PaginationData[entity.Author]
}

type GetAuthorByIDReq struct {
	ID string `swaggerignore:"true"`
}

type GetAuthorByIDRes struct {
	entity.Author
}
