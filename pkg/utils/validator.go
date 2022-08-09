package utils

import (
	// "fmt"

	"github.com/go-playground/validator/v10"
)

func CustomValidateStruct(err error) []map[string]string {
	var errMess []map[string]string
	errMess = []map[string]string{
		{
			"field": err.Error(),
			"msg":   err.Error(),
		},
	}
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		errMess = []map[string]string{}
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errMess = append(errMess, map[string]string{
					"field": err.Field(),
					"msg":   "field is required",
				})
			case "email":
				errMess = append(errMess, map[string]string{
					"field": err.Field(),
					"msg":   "not a valid email",
				})
				// default:
			}
		}
	}
	return errMess

}
