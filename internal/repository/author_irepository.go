package repository

import (
	"books-api/internal/entity"
)

type AuthorRepository interface {
	// Example operations
	CommonQuery[entity.Author]
}
