package httpv1

import (
	"github.com/Hona-Tahlil/Backend/internal/presentation/controllers/v1/general"
	"github.com/gin-gonic/gin"
)

func SetUpGeneralRouts(v1 *gin.RouterGroup) {
	v1.POST("/test", general.Test)
}
