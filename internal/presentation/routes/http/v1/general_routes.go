package httpv1

import (
	"hona/backend/wire"

	"github.com/gin-gonic/gin"
)

func SetUpGeneralRoutes(v1 *gin.RouterGroup, app *wire.Application) {

	auth := v1.Group("/auth")
	{
		auth.POST("/login", app.Controllers.GeneralControllers.GeneralAuthController.Login)
	}
}
