package postgres

import (
	"hona/backend/internal/domain/entities"

	"gorm.io/gorm"
)

// TODO: manteghie?
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (up *UserRepository) FindUserByEmail(email string) *entities.User {
	var user *entities.User
	if user := up.db.First(user, email); user == nil {
		// TODO: Proper error handling
		panic("user not found")
	}
	return user
}
