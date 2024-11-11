package entity

import (
	"os"
)

type Book struct {
	Id         string `json:"id" gorm:"primaryKey;type:uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title      string `json:"title" example:"The Great Gatsby"`
	ISBN       string `json:"isbn" example:"978-3-16-148410-0" gorm:"uniqueIndex"`
	AuthorId   string `json:"author_id" gorm:"uuid" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	CategoryId string `json:"category_id" gorm:"uuid" example:"7ea2b005-a3a5-4f4b-b762-75234fb6b4bd"`
	Stock      uint   `json:"stock" example:"1" gorm:"default:1"`
}

func (model *Book) TableName() string {
	return os.Getenv("DB_PREFIX") + "book"
}
