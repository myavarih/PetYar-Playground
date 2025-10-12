package login

import "hona/backend/internal/application/dto/rbac"

type LoginResponse struct {
	AccessToken  string                    `json:"accessToken"`
	RefreshToken string                    `json:"refreshToken"`
	Name         string                    `json:"name"`
	Permissions  []rbac.PermissionResponse `json:"permissions"`
}
