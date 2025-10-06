package exceptions

import "fmt"

type AuthError struct {
	Type        string
	Message     string
	OriginalErr error
}

func (e AuthError) Error() string {
	if e.OriginalErr != nil {
		return fmt.Sprintf("%s: %s (%s)", e.Type, e.Message, e.OriginalErr.Error())
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func NewUnauthorizedError(message string, originalErr error) *AuthError {
	if message == "" {
		message = "Unauthorized access"
	}

	return &AuthError{
		Type:        "UNAUTHORIZED",
		Message:     message,
		OriginalErr: originalErr,
	}
}
