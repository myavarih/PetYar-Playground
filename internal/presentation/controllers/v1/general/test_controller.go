package general

import (
	testdto "github.com/Hona-Tahlil/Backend/internal/application/dto/test"
	"github.com/Hona-Tahlil/Backend/internal/application/service"
	"github.com/Hona-Tahlil/Backend/internal/infrastructure/translation"
	"github.com/Hona-Tahlil/Backend/internal/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	type testParams struct {
		Num1 int `json:"num1" validate:"required"`
		Num2 int `json:"num2" validate:"required"`
	}

	params := controllers.Receive[testParams](ctx)

	data := testdto.TestRequest{
		Num1: params.Num1,
		Num2: params.Num2,
	}

	res := service.Test(data)

	controllers.Respond(ctx, res, translation.Message{Text: "success.test", Params: []string{}})
}
