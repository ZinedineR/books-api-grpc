package repository

import (
	"books-api/internal/entity"
)

type BookLendingSQLRepo struct {
	Repository[entity.BookLending]
}

func NewBookLendingSQLRepository() BookLendingRepository {
	return &BookLendingSQLRepo{}
}
