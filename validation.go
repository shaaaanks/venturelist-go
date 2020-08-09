package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var v = validator.New()

func validate(item interface{}) error {
	err := v.Struct(item)

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
