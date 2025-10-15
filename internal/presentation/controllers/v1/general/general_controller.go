package general

import (
	"hona/backend/internal/application/dto/user"
	"hona/backend/internal/application/service"
	"hona/backend/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

type GeneralController struct {
	generalService *service.GeneralService
}

func NewGeneralController(generalService *service.GeneralService) *GeneralController {
	return &GeneralController{
		generalService: generalService,
	}
}

func (gc *GeneralController) Login(ctx *gin.Context) {
	type loginParams struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=64"`
	}

	params := controllers.Receive[loginParams](ctx)
	loginInfo := user.LoginRequest{
		Email:    params.Email,
		Password: params.Password,
	}

	res := gc.generalService.Login(loginInfo)

	msg := controllers.Message{
		Text:   "successMessage.login",
		Params: []string{},
	}
	controllers.Respond(ctx, 200, msg, res)
}
