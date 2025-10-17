package service

import (
	"hona/backend/internal/application/dto/rbac"
	"hona/backend/internal/application/dto/user"
	"hona/backend/internal/domain/exceptions"
	"hona/backend/internal/infrastructure/jwt"
	"hona/backend/internal/infrastructure/repository/postgres"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	jwtService *jwt.JWTService
	unitOfWork *postgres.UnitOfWork
}

func NewUserService(unitOfWork *postgres.UnitOfWork, jwtService *jwt.JWTService) *UserService {
	return &UserService{
		unitOfWork: unitOfWork,
		jwtService: jwtService,
	}
}

func (us *UserService) Login(loginInfo user.LoginRequest) (user.LoginResponse, string) {
	foundUser := us.unitOfWork.Factory().UserRepository().FindUserByEmail(loginInfo.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginInfo.Password)); err != nil {
		invalidCredentialsErr := &exceptions.AuthError{
			Type: "INVALID_CREDENTIALS",
		}
		panic(invalidCredentialsErr)
	}

	accessToken, refreshToken := us.jwtService.GenerateTokens(foundUser.ID, loginInfo.RememberMe)

	p := make([]rbac.PermissionResponse, 0)
	for _, role := range foundUser.Roles {
		for _, per := range role.Permissions {
			p = append(p, rbac.PermissionResponse{
				ID:   per.ID,
				Name: per.Type.String(),
			})
		}
	}

	return user.LoginResponse{
		AccessToken: accessToken,
		Permissions: p,
	}, refreshToken
}

func (us *UserService) RefreshToken(refreshTokenInfo user.RefreshTokenRequest) (user.RefreshTokenResponse, string) {
	accessToken, refreshToken, userID := us.jwtService.RefreshTokens(refreshTokenInfo.RefreshToken)

	foundUser := us.unitOfWork.Factory().UserRepository().FindUserByID(userID)

	p := make([]rbac.PermissionResponse, 0)
	for _, role := range foundUser.Roles {
		for _, per := range role.Permissions {
			p = append(p, rbac.PermissionResponse{
				ID:   per.ID,
				Name: per.Type.String(),
			})
		}
	}

	return user.RefreshTokenResponse{
		AccessToken: accessToken,
		Permissions: p,
	}, refreshToken
}
