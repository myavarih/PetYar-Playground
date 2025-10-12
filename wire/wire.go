//go:build wireinject
// +build wireinject

package wire

import (
	"hona/backend/bootstrap"
	"hona/backend/internal/infrastructure/database"

	"github.com/google/wire"
)

var DatabaseProviderSet = wire.NewSet(
	database.NewPostgresDatabase,
	wire.Bind(new(database.Database), new(*database.PostgresDatabase)),
	wire.Struct(new(Database), "*"),
)

var ProviderSet = wire.NewSet(
	DatabaseProviderSet,
)

type Database struct {
	DB database.Database
}

type Application struct {
	Database *Database
}

func NewApplication(db *Database) *Application {
	return &Application{
		Database: db,
	}
}

func InitializeApplication(container *bootstrap.Config) (*Application, error) {
	wire.Build(
		ProviderSet,
		NewApplication,
	)
	return &Application{}, nil
}
