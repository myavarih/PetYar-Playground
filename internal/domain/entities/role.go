package entities

type Role struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Permissions []Permission `gorm:"foreignKey:RoleID"`
}
