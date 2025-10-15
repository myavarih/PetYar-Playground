package entities

type Permission struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	RoleID uint
}
