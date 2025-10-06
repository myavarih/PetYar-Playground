package main

import (
	"github.com/Hona-Tahlil/Backend/bootstrap"
	"github.com/Hona-Tahlil/Backend/internal/presentation/routing"
	"github.com/gin-gonic/gin"
)

// TODO: read all files: verify one by one (+compare with barghe-no)

// TODO: make sure the receivers and their constructors are set up correctly

// TODO: try to understand wire

// TODO: try to understand interfaces
func main() {
	gin.DisableConsoleColor()

	ginEngine := gin.Default()

	bootstrap.Run()

	bootstrap.DBMigration()

	routing.SetUpRouts(ginEngine)

	ginEngine.Run()
}
