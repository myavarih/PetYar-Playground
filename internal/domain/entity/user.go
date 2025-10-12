package entity

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Name     string
}
