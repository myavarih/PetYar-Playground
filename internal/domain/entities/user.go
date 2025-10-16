package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Name        string
	Address     string
	BirthDate   time.Time
	Phone       string
	City        string
	PictureLink string
	Wallet      Wallet    `gorm:"foreignKey:UserID"`
	Requests    []Request `gorm:"foreignKey:UserID"`
	Role        Role      `gorm:"foreignKey:RoleID"`
	RoleID      uint
}
