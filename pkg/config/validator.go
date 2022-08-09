package config

import (
	"github.com/go-playground/validator/v10"
)

//var validate *validator.Validate

func ValidateStruct(entity interface{}) error {
	validate := validator.New()

	err := validate.Struct(entity)
	if err != nil {
		return err
	}
	return nil
}
