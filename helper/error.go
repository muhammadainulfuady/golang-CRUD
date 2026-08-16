package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errorsMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := strings.ToLower(fieldError.Field())
			switch fieldError.Tag() {
			case "required":
				errorsMap[fieldName] = fmt.Sprintf("field %s wajib di isi", fieldName)
			case "min":
				errorsMap[fieldName] = fmt.Sprintf("field %s minimal %s kareakter", fieldName, fieldError.Param())
			case "max":
				errorsMap[fieldName] = fmt.Sprintf("field %s maximal %s kareakter", fieldName, fieldError.Param())
			case "datetime":
				errorsMap[fieldName] = fmt.Sprintf("field %s tidak valid", fieldName)
			default:
				errorsMap[fieldName] = fmt.Sprintf("field %s tidak valid", fieldName)
			}
		}
	} else {
		errorsMap["error"] = err.Error()
	}
	return errorsMap
}
