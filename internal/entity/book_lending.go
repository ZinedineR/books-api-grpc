package entity

import (
	"os"
	"time"
)

type BookLending struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	UserId    string    `json:"user_id" gorm:"uuid" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	BookId    string    `json:"book_id" gorm:"uuid" example:"7ea2b005-a3a5-4f4b-b762-75234fb6b4bd"`
	Returned  bool      `json:"returned" gorm:"default:false"`
	StartDate time.Time `json:"start_date" example:"2020-01-01 00:00:00 +0000 UTC"`
	EndDate   time.Time `json:"end_date" example:"2020-01-01 00:00:00 +0000 UTC"`

	//Relationship
	Book *Book `gorm:"foreignKey:BookId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"book,omitempty"`
}

func (model *BookLending) TableName() string {
	return os.Getenv("DB_PREFIX") + "book_lending"
}
