package entity

import (
	"os"
)

type Category struct {
	Id    string `json:"id" gorm:"primaryKey;type:uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title string `json:"title" example:"Romance"`
}

func (model *Category) TableName() string {
	return os.Getenv("DB_PREFIX") + "category"
}
