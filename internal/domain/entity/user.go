package entity

// TODO: decide what should be the primary key

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Name     string
}
