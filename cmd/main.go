package main

import (
	"hona/backend/bootstrap"
	"hona/backend/internal/presentation/routes"

	"github.com/gin-gonic/gin"
)

// TODO: Add Interfaces!

// TODO: Implement RBAC Response Format

func main() {
	gin.DisableConsoleColor()

	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	bootstrap.Run()

	routes.SetUpRoutes(ginEngine)

	ginEngine.Run()
}
