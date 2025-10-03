package middleware

import (
	"net/http"

	"github.com/Hona-Tahlil/Backend/internal/infrastructure/translation"
	"github.com/gin-gonic/gin"
)

func GetLocale(request *http.Request) string {
	return request.Header.Get("Accept-Language")
}

func AddTranslator(ctx *gin.Context) {
	locale := GetLocale(ctx.Request)

	trans := translation.GetTranslator(locale)

	ctx.Set("translator", trans)

	ctx.Next()
}
