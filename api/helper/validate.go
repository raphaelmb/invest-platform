package helper

import "github.com/go-playground/validator/v10"

func Validate(v any) error {
	validate := validator.New()
	err := validate.Struct(v)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}
