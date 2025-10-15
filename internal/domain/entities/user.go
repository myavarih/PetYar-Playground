package entities

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Name     string
	Role     Role `gorm:"foreignKey:RoleID"`
	RoleID   uint
}
