package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func ParseErrors(errors validate.Errors) string {
	var errorDescriptions string
	for key, val := range errors.Errors {
		errorDescriptions += key + ": " + strings.Join(val, ", ") + "\n"
	}
	return errorDescriptions
}
