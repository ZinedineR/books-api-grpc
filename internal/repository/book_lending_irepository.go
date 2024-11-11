package repository

import (
	"books-api/internal/entity"
)

type BookLendingRepository interface {
	CommonQuery[entity.BookLending]
}
