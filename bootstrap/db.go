package bootstrap

import (
	"fmt"

	"github.com/Hona-Tahlil/Backend/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: and this
var DB *gorm.DB

// TODO: move this to infra/database
func DBMigration() {
	dsn := GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database"))
	}

	DB = db

	DB.AutoMigrate(&entity.User{})
}

func GetDSN() string {
	dbConfig := ProjectConfig.Env.PrimaryDatabase
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)
	return dsn
}
