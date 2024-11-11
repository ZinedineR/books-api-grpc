package repository

import (
	"books-api/internal/entity"
)

type CategoryRepository interface {
	// Example operations
	CommonQuery[entity.Category]
}
