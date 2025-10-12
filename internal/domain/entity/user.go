package entity

// TODO: decide what should be the primary key

type User struct {
	Email    string `gorm:"primaryKey"`
	Password string `gorm:"not null"`
}
