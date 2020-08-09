package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

func validate(item interface{}) error {
	validate := validator.New()
	err := validate.Struct(item)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("The %v field is %v", err.Field(), err.Tag())
		}
	}

	return nil
}
