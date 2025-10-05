package main

import (
	"github.com/Hona-Tahlil/Backend/bootstrap"
	"github.com/Hona-Tahlil/Backend/internal/presentation/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	ginEngine := gin.Default()

	bootstrap.Run()

	bootstrap.DBMigration()

	routing.SetUpRouts(ginEngine)

	ginEngine.Run()
}
