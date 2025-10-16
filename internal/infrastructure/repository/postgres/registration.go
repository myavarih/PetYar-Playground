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

		db.Migrator().DropTable(
			&entities.User{},
			&entities.Role{},
			&entities.Permission{},
			&entities.Wallet{},
			&entities.Request{},
			&entities.CalenderSlot{},
			entities.Pet{},
			entities.PetSitter{},
			entities.Reserve{},
			entities.Service{},
			entities.Chat{},
			entities.Comment{},
		)
		db.AutoMigrate(
			&entities.User{},
			&entities.Role{},
			&entities.Permission{},
			&entities.Wallet{},
			&entities.Request{},
			&entities.CalenderSlot{},
			entities.Pet{},
			entities.PetSitter{},
			entities.Reserve{},
			entities.Service{},
			entities.Chat{},
			entities.Comment{},
		)

	})
	return dbInstance
}
