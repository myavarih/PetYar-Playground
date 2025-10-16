package middleware

import (
	"hona/backend/bootstrap"
	"hona/backend/internal/domain/enums"
	"hona/backend/internal/domain/exceptions"
	"hona/backend/internal/infrastructure/repository/postgres"
	"log"

	"github.com/gin-gonic/gin"
)

type RBACMiddleware struct {
	userRepository *postgres.UserRepository
}

func NewRBACMiddleware(userRepository *postgres.UserRepository) *RBACMiddleware {
	return &RBACMiddleware{
		userRepository: userRepository,
	}
}

func (rm *RBACMiddleware) HasAccess(allowedPermissions []enums.Permission) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _ := ctx.Get(bootstrap.ProjectConfig.Constants.Context.ID)
		user := rm.userRepository.FindUserByID(userID.(uint))
		var allowed bool = false
		for _, permission := range allowedPermissions {
			for _, p := range user.Role.Permissions {
				if p.Type == permission {
					allowed = true
					break
				}
			}
			if allowed {
				break
			}
		}
		if !allowed {
			log.Println("works!")
			unauthorizedErr := exceptions.NewUnauthorizedError("you don't have the required access")
			panic(unauthorizedErr)
		}
		ctx.Next()
	}
}
