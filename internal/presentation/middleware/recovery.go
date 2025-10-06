package middleware

import (
	"log"
	"strconv"

	"github.com/Hona-Tahlil/Backend/bootstrap"
	"github.com/Hona-Tahlil/Backend/internal/domain/exceptions"
	"github.com/Hona-Tahlil/Backend/internal/infrastructure/translation"
	"github.com/Hona-Tahlil/Backend/internal/presentation/controllers"
	"github.com/gin-gonic/gin"
)

type RecoveryMiddleware struct {
	constants *bootstrap.Constants
}

func NewRecoveryMiddleware(constants *bootstrap.Constants) *RecoveryMiddleware {
	return &RecoveryMiddleware{
		constants: constants,
	}
}

func (rm *RecoveryMiddleware) Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				msgs := rm.handleError(err)
				if len(msgs) == 1 {
					controllers.Respond(ctx, nil, msgs[0])
				} else {
					controllers.Respond(ctx, nil, msgs)
				}
			}
			ctx.Abort()
		}
	}()
}

func (rm *RecoveryMiddleware) handleError(err error) []translation.Message {
	if bindingErr, ok := err.(*exceptions.BindingError); ok {
		return rm.handleBindingError(bindingErr)
	} else if validationErrors, ok := err.(*exceptions.ValidationErrors); ok {
		return rm.handleValidationErrors(validationErrors)
	}
	return rm.unhandledErrors(err)
}

func (rm *RecoveryMiddleware) handleBindingError(bindingErr *exceptions.BindingError) []translation.Message {
	if numError, ok := bindingErr.Err.(*strconv.NumError); ok {
		msg := translation.Message{
			Text:   "errors.numeric",
			Params: []string{numError.Num},
		}
		return []translation.Message{msg}
	}
	msg := translation.Message{
		Text:   "errors.binding",
		Params: []string{},
	}
	return []translation.Message{msg}
}

func (rm *RecoveryMiddleware) handleValidationErrors(validationErrs *exceptions.ValidationErrors) []translation.Message {
	var msgs []translation.Message
	for i, fieldErr := range validationErrs.FieldErrors {
		msgs[i] = translation.Message{
			Text:       "errors.validation",
			FieldError: &fieldErr,
		}
	}
	return msgs
}

func (rm *RecoveryMiddleware) unhandledErrors(err error) []translation.Message {
	log.Println("an unhandled error occurred", err.Error())
	msg := translation.Message{
		Text:   "errors.generic",
		Params: []string{},
	}
	return []translation.Message{msg}
}
