package main

import (
	"hona/backend/bootstrap"
	"hona/backend/internal/infrastructure/repository/postgres"
	"hona/backend/internal/presentation/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	bootstrap.Run()

	postgres.NewPostgresDatabase()

	routes.SetUpRoutes(ginEngine)

	ginEngine.Run()
}
