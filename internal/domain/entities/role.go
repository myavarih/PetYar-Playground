package entities

import "hona/backend/internal/domain/enums"

type Role struct {
	ID          uint `gorm:"primaryKey"`
	Type        enums.Role
	Permissions []Permission `gorm:"foreignKey:RoleID"`
}
