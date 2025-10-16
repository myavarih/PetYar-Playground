package entities

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	UserID     uint
	IsApproved bool
	IsPaid     bool
}
