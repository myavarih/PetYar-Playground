package routing

import (
	httpv1 "github.com/Hona-Tahlil/Backend/internal/presentation/routing/http/v1"
	"github.com/gin-gonic/gin"
)

func SetUpRouts(ginEngine *gin.Engine) {
	v1 := ginEngine.Group("/v1")
	httpv1.SetUpGeneralRouts(v1)
}
