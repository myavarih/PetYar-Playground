package jwt

// TODO: continue developing this
// import (
// 	"errors"

// 	"github.com/Hona-Tahlil/Backend/internal/domain/exceptions"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func ValidateToken(tokenString string) (jwt.Claims, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, exceptions.NewInvalidTokenError(nil)
// 		}
// 		return jwtService.keyManager.GetPublicKey(), nil
// 	})

// 	if err != nil {
// 		if errors.Is(err, jwt.ErrTokenExpired) {
// 			return nil, exception.NewExpiredTokenError(err)
// 		}
// 		return nil, exception.NewInvalidTokenError(err)
// 	}

// 	if !token.Valid {
// 		return nil, exception.NewInvalidTokenError(nil)
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, exception.NewInvalidTokenError(nil)
// 	}

// 	return claims, nil

// }
