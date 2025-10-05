package entity

type User struct {
	Email    string `gorm:"primaryKey"`
	Name     string
	Password string
}
