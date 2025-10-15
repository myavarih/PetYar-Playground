package service

import (
	"hona/backend/internal/application/dto/rbac"
	"hona/backend/internal/application/dto/user"
	"hona/backend/internal/domain/exceptions"
	"hona/backend/internal/infrastructure/jwt"
	"hona/backend/internal/infrastructure/repository/postgres"

	"golang.org/x/crypto/bcrypt"
)

type GeneralService struct {
	jwtService *jwt.JWTService
	unitOfWork *postgres.UnitOfWork
}

func NewGeneralService(unitOfWork *postgres.UnitOfWork, jwtService *jwt.JWTService) *GeneralService {
	return &GeneralService{
		unitOfWork: unitOfWork,
		jwtService: jwtService,
	}
}

func (gs *GeneralService) Login(loginInfo user.LoginRequest) user.LoginResponse {
	foundUser := gs.unitOfWork.Factory().UserRepository().FindUserByEmail(loginInfo.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginInfo.Password)); err != nil {
		invalidCredentialsErr := &exceptions.AuthError{
			Type: "INVALID_CREDENTIALS",
		}
		panic(invalidCredentialsErr)
	}

	accessToken, refreshToken := gs.jwtService.GenerateTokens(foundUser.ID)

	// foundUser := entities.User{}

	permissions := foundUser.Role.Permissions

	p := make([]rbac.PermissionResponse, 0)

	for _, per := range permissions {
		p = append(p, rbac.PermissionResponse{
			ID:   per.ID,
			Name: per.Name,
		})
	}

	return user.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Role: rbac.RoleResponse{
			ID:          foundUser.Role.ID,
			Name:        foundUser.Role.Name,
			Permissions: p,
		},
	}
}
