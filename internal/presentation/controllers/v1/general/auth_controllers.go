package general

import (
	"hona/backend/internal/application/dto/user"
	"hona/backend/internal/application/service"
	"hona/backend/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

type GeneralAuthController struct {
	generalService *service.UserService
}

func NewGeneralAuthController(generalService *service.UserService) *GeneralAuthController {
	return &GeneralAuthController{
		generalService: generalService,
	}
}

func (gc *GeneralAuthController) Login(ctx *gin.Context) {
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
