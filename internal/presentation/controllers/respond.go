package controllers

import (
	"net/http"

	"github.com/Hona-Tahlil/Backend/bootstrap"
	statuscodes "github.com/Hona-Tahlil/Backend/internal/domain/statusCodes"
	"github.com/Hona-Tahlil/Backend/internal/infrastructure/translation"
	"github.com/gin-gonic/gin"
)

type RespondController struct {
	constants *bootstrap.Constants
}

func NewRespondController(constants *bootstrap.Constants) *RespondController {
	return &RespondController{
		constants: constants,
	}
}

type singleMessageResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type multipleMessageResponse struct {
	StatusCode int               `json:"statusCode"`
	Messages   map[string]string `json:"messages"`
	Data       interface{}       `json:"data"`
}

func Respond[T translation.Message | []translation.Message](ctx *gin.Context, data interface{}, msgs T) {
	translator := GetTranslator(ctx, NewRespondController(bootstrap.NewConstants()).constants.Context.Translator)

	switch msg := any(msgs).(type) {
	case translation.Message:
		sCode := statuscodes.StatusCodes[msg.Text]
		if msg.Text == "" {
			msg.Text = http.StatusText(sCode)
		}
		message, _ := translator.T(msg.Text, msg.Params...)
		ctx.JSON(sCode, singleMessageResponse{
			StatusCode: sCode,
			Message:    message,
			Data:       data,
		})
	case []translation.Message:
		sCode := statuscodes.StatusCodes[msg[0].Text]
		mms := multipleMessageResponse{
			StatusCode: sCode,
			Messages:   map[string]string{},
			Data:       data,
		}
		for _, ms := range msg {
			translatedTagValue, _ := translator.T(ms.FieldError.Field)
			translatedTag, _ := translator.T("errors."+ms.FieldError.Tag, translatedTagValue)
			mms.Messages[ms.FieldError.Tag] = translatedTag
		}
		ctx.JSON(sCode, mms)
	}
}
