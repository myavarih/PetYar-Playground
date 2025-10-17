package user

type LoginRequest struct {
	Email      string
	Password   string
	RememberMe bool
}

type RefreshTokenRequest struct {
	RefreshToken string
}
