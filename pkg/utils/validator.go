package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func CustomValidateStruct(err error) string {
	errMessage := err.Error()
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errMessage = fmt.Sprintf("%s is required", err.Field())
			case "email":
				errMessage = fmt.Sprintf("%s is not a valid email", err.Field())
			default:
				errMessage = fmt.Sprintf("%s is invalid", err.Field())
			}
		}
	}
	return errMessage

}
