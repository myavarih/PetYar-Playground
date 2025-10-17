package middleware

import (
	"net/http"

	"hona/backend/bootstrap"
	"hona/backend/internal/infrastructure/translation"

	"github.com/gin-gonic/gin"
)

// ? what should it have?
type LocalizationMiddleware struct {
	context *bootstrap.Context
}

func NewLocalizationMiddleware(context *bootstrap.Context) *LocalizationMiddleware {
	return &LocalizationMiddleware{
		context: context,
	}
}

func GetLocale(request *http.Request) string {
	return request.Header.Get("Accept-Language")
}

func (lm *LocalizationMiddleware) AddTranslator(ctx *gin.Context) {
	locale := GetLocale(ctx.Request)

	trans := translation.GetTranslator(locale)

	ctx.Set(lm.context.Translator, trans)

	ctx.Next()
}
