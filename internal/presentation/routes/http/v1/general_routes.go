package httpv1

import (
	"hona/backend/internal/application/service"
	"hona/backend/internal/infrastructure/jwt"
	"hona/backend/internal/infrastructure/repository/postgres"
	"hona/backend/internal/presentation/controllers/v1/general"

	"github.com/gin-gonic/gin"
)

func SetUpGeneralRoutes(v1 *gin.RouterGroup) {
	db := postgres.NewPostgresDatabase()
	u := postgres.NewUnitOfWork(db.DB)
	js := jwt.NewJWTService(jwt.NewJWTKeyManager())
	s := service.NewGeneralService(u, js)
	gc := general.NewGeneralController(s)

	auth := v1.Group("/auth")
	{
		auth.POST("/login", gc.Login)
	}
}
