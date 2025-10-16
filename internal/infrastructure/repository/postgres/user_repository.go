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

	if result := up.db.First(foundUser, "email = ?", email); result.Error != nil {
		invalidCredentialsErr := &exceptions.AuthError{
			Type: "INVALID_CREDENTIALS",
		}
		panic(invalidCredentialsErr)
	}
	return foundUser
}

func (up *UserRepository) FindUserByID(userID uint) *entities.User {
	var foundUser *entities.User

	if result := up.db.First(foundUser, userID); result.Error != nil {
		// invalidCredentialsErr := &exceptions.AuthError{
		// 	Type: "INVALID_CREDENTIALS",
		// }
		// panic(invalidCredentialsErr)
	}
	return foundUser
}
