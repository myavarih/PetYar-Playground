package main

import (
	"hona/backend/bootstrap"
	"hona/backend/internal/presentation/routes"
	"hona/backend/wire"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	bootstrap.Run()

	app, err := wire.InitializeApplication(bootstrap.ProjectConfig)
	if err != nil {
		panic(err)
	}

	routes.SetUpRoutes(ginEngine, app)

	ginEngine.Run()
}
