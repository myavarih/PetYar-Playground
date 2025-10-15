package postgres

import (
	"hona/backend/internal/domain/entities"
	"hona/backend/internal/domain/exceptions"

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

func (up *UserRepository) FindUserByEmail(email string) *entities.User {
	var foundUser *entities.User

	if result := up.db.First(foundUser, email); result.Error != nil {
		notFoundErr := &exceptions.NotFoundError{
			Item: "email",
		}
		panic(notFoundErr)
	}
	return foundUser
}
