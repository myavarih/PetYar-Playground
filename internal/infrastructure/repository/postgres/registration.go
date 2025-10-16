package postgres

import (
	"fmt"
	"hona/backend/bootstrap"
	"hona/backend/internal/domain/entities"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var dbOnce sync.Once

func NewPostgresDatabase() *gorm.DB {
	dsn := bootstrap.ProjectConfig.Env.DSN

	dbOnce.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("failed to connect database"))
		}

		dbInstance = db

		dbInstance.AutoMigrate(&entities.User{}, entities.Role{}, entities.Permission{})

	})
	return dbInstance
}
