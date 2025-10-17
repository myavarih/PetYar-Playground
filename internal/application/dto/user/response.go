package user

import "hona/backend/internal/application/dto/rbac"

type LoginResponse struct {
	AccessToken string                    `json:"accessToken"`
	Permissions []rbac.PermissionResponse `json:"permissions"`
}

type RefreshTokenResponse struct {
	AccessToken string                    `json:"accessToken"`
	Permissions []rbac.PermissionResponse `json:"permissions"`
}
