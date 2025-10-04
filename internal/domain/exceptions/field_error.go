package exceptions

import "fmt"

type FieldError struct {
	Field string
	Tag   string
}

func (fe *FieldError) Error() string {
	return fmt.Sprintf("Error:Field validation for '%s' failed on the '%s' tag", fe.Field, fe.Tag)
}

func NewFieldError(field, tag string) *FieldError {
	return &FieldError{
		Field: field,
		Tag:   tag,
	}
}
