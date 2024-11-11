package repository

import (
	"books-api/internal/entity"
)

type CategorySQLRepo struct {
	Repository[entity.Category]
}

func NewCategorySQLRepository() CategoryRepository {
	return &CategorySQLRepo{}
}
