package service

import (
	"hona/backend/internal/application/dto/login"
	"hona/backend/internal/application/dto/rbac"
	"hona/backend/internal/infrastructure/jwt"
	"hona/backend/internal/infrastructure/repository/postgres"
)

type GeneralService struct {
	unitOfWork *postgres.UnitOfWork
}

func NewGeneralService(unitOfWork *postgres.UnitOfWork) *GeneralService {
	return &GeneralService{
		unitOfWork: unitOfWork,
	}
}

func (gs *GeneralService) Login(loginInfo login.LoginRequest) login.LoginResponse {
	// TODO: actually implement
	user := gs.unitOfWork.Factory().UserRepository().FindUserByEmail(loginInfo.Email)

	// TODO: not a good way
	js := jwt.NewJWTService(jwt.NewJWTKeyManager())
	accessToken, refreshToken := js.GenerateTokens(user.ID)

	return login.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Name:         user.Name,
		Permissions:  []rbac.PermissionResponse{},
	}
}
