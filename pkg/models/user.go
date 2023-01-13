package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	LastName    string    `gorm:"type:varchar(100)`
	FirstName   string    `gorm:"type:varchar(100)`
	Email       string    `gorm:"type:varchar(50);unique_index"`
	PhoneNumber string    `gorm:"type:varchar(15)`
	Password    string    `json:"password"`
}
