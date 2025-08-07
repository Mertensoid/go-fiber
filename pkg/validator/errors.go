package validator

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/validate"
)

func ParseErrors(errors validate.Errors) string {
	var errorDescriptions string
	for key, val := range errors.Errors {
		errorDescriptions += fmt.Sprintf("%s: %s\n", key, strings.Join(val, ", "))
	}
	return errorDescriptions
}
