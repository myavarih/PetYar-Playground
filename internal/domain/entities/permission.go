package entities

import "hona/backend/internal/domain/enums"

type Permission struct {
	ID     uint `gorm:"primaryKey"`
	Type   enums.Permission
	RoleID uint
}
