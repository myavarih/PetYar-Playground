package middleware

// import (
// 	"strings"

// 	"github.com/Hona-Tahlil/Backend/bootstrap"
// 	"github.com/Hona-Tahlil/Backend/internal/domain/exceptions"
// 	"github.com/gin-gonic/gin"
// )

// func AuthRequired(ctx *gin.Context) {
// 	authHeader := ctx.GetHeader("Authorization")
// 	if authHeader == "" {
// 		unauthorizedError := exceptions.NewUnauthorizedError("empty auth header", nil)
// 		panic(unauthorizedError)
// 	}

// 	parts := strings.Split(authHeader, " ")
// 	if len(parts) != 2 || parts[0] != "Bearer" {
// 		unauthorizedError := exceptions.NewUnauthorizedError("invalid token format", nil)
// 		panic(unauthorizedError)
// 	}

// 	tokenString := parts[1]
// 	if tokenString == "" {
// 		unauthorizedError := exceptions.NewUnauthorizedError("empty token", nil)
// 		panic(unauthorizedError)
// 	}

// 	claims, err := ValidateToken(tokenString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ctx.Set(bootstrap.ProjectConfig.Constants.Context.ID, uint(claims["sub"].(float64)))

// 	ctx.Next()
// }
