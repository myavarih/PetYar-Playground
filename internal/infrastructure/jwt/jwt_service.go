package jwt

import (
	"errors"
	"hona/backend/internal/domain/exceptions"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	keyManager *JWTKeyManager
}

func NewJWTService(keyManager *JWTKeyManager) *JWTService {
	return &JWTService{
		keyManager: keyManager,
	}
}

func (js *JWTService) GenerateTokens(userID uint) (accessTokenString string, refreshTokenString string) {
	accessTokenClaims, refreshTokenClaims := js.GenerateClaims(userID)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenClaims)

	var err error
	accessTokenString, err = accessToken.SignedString(js.keyManager.GetPrivateKey())
	if err != nil {
		panic(err)
	}
	refreshTokenString, err = refreshToken.SignedString(js.keyManager.GetPrivateKey())
	if err != nil {
		panic(err)
	}

	return
}

func (js *JWTService) GenerateClaims(userID uint) (accessTokenClaims jwt.MapClaims, refreshTokenClaims jwt.MapClaims) {
	accessTokenClaims = jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Minute * 2).Unix(),
		"iat": time.Now().Unix(),
	}
	refreshTokenClaims = jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
	}
	return
}

func (js *JWTService) ValidateToken(tokenString string) uint {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, exceptions.NewInvalidTokenError()
		}
		return js.keyManager.GetPublicKey(), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			expiredTokenErr := exceptions.NewExpiredTokenError()
			panic(expiredTokenErr)
		}
		invalidTokenErr := exceptions.NewInvalidTokenError()
		panic(invalidTokenErr)
	}

	if !token.Valid {
		invalidTokenErr := exceptions.NewInvalidTokenError()
		panic(invalidTokenErr)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		invalidTokenErr := exceptions.NewInvalidTokenError()
		panic(invalidTokenErr)
	}

	userID := uint(claims["sub"].(float64))
	return userID
}

func (js *JWTService) RefreshTokens(refreshTokenString string) (accessTokenString string, newRefreshTokenString string) {
	userID := js.ValidateToken(refreshTokenString)

	accessTokenString, newRefreshTokenString = js.GenerateTokens(userID)
	return
}
