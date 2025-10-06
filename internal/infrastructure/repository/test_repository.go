package repository

import (
	"github.com/Hona-Tahlil/Backend/bootstrap"
	"github.com/Hona-Tahlil/Backend/internal/domain/entity"
)

// TODO: should have receiver
func Test() {
	bootstrap.DB.Create(entity.User{
		Email:    "email",
		Name:     "asghar",
		Password: "myPass",
	})
}
