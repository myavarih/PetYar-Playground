package user

import "hona/backend/internal/application/dto/rbac"

type LoginResponse struct {
	AccessToken  string            `json:"accessToken"`
	RefreshToken string            `json:"refreshToken"`
	Role         rbac.RoleResponse `json:"role"`
}
