package entity

import (
	"os"
	"time"
)

type Author struct {
	Id        string     `json:"id" validate:"required,uuid" gorm:"primaryKey;type:uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name      string     `json:"name" validate:"required,title" example:"F. Scott Fitzgerald"`
	Birthdate *time.Time `json:"birthdate" example:"1896-09-24"`
}

func (model *Author) TableName() string {
	return os.Getenv("DB_PREFIX") + "author"
}
