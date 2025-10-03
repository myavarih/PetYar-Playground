package validation

import "github.com/go-playground/validator/v10"

func ValidateFields(params any) error {
	val := validator.New()
	if err := val.Struct(params); err != nil {
		// TODO: Proper Error Handling
		return err
	}
	return nil
}
