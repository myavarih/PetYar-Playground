package routing

import (
	"github.com/Hona-Tahlil/Backend/bootstrap"
	"github.com/Hona-Tahlil/Backend/internal/presentation/middleware"
	httpv1 "github.com/Hona-Tahlil/Backend/internal/presentation/routing/http/v1"
	"github.com/gin-gonic/gin"
)

func SetUpRouts(ginEngine *gin.Engine) {
	ginEngine.Use(middleware.AddTranslator)
	ginEngine.Use(middleware.NewRecoveryMiddleware(bootstrap.NewConstants()).Recover)

	v1 := ginEngine.Group("/v1")
	httpv1.SetUpGeneralRouts(v1)
}
