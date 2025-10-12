package postgres

import (
	"hona/backend/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (up *UserRepository) FindUserByEmail(email string) *entity.User {
	var user entity.User
	if err := up.db.First(&user, email); err != nil {
		// TODO: Proper error handling
		panic(err)
	}
	return &user
}
