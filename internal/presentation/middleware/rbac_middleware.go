package middleware

import (
	"hona/backend/bootstrap"
	"hona/backend/internal/domain/enums"
	"hona/backend/internal/domain/exceptions"
	"hona/backend/internal/infrastructure/repository/postgres"

	"github.com/gin-gonic/gin"
)

type RBACMiddleware struct {
	uintOfWork *postgres.UnitOfWork
}

func NewRBACMiddleware(uintOfWork *postgres.UnitOfWork) *RBACMiddleware {
	return &RBACMiddleware{
		uintOfWork: uintOfWork,
	}
}

func (rm *RBACMiddleware) NeedsPermission(allowedPermissions []enums.Permission) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _ := ctx.Get(bootstrap.ProjectConfig.Constants.Context.ID)
		user := rm.uintOfWork.Factory().UserRepository().FindUserByID(userID.(uint))
		var allowed bool = false
		for _, permission := range allowedPermissions {
			for _, role := range user.Roles {
				for _, p := range role.Permissions {
					if p.Type == permission {
						allowed = true
						break
					}
				}
				if allowed {
					break
				}
			}
		}
		if !allowed {
			accessDeniedErr := exceptions.NewAccessDeniedError("you don't have the required access")
			panic(accessDeniedErr)
		}
		ctx.Next()
	}
}
