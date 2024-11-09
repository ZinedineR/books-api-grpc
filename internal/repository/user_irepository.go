package repository

import (
	"books-api/internal/entity"
)

type UserRepository interface {
	// Example operations
	CommonQuery[entity.User]
}
