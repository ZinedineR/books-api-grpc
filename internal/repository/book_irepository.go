package repository

import (
	"books-api/internal/entity"
)

type BookRepository interface {
	CommonQuery[entity.Book]
}
