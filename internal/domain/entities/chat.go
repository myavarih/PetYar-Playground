package entities

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	History            string // don't know how to store it
	IsUserBlocked      bool
	IsPetSitterBlocked bool
	IsAccepted         bool
	RequestID          uint
}
