package card

import "github.com/madeinly/core/validation"

func IdRules(v string) []*validation.Error {
	var errs []*validation.Error
	if v == "" {
		errs = append(errs, &validation.Error{
			Field: "name", Code: "EMPTY_VALUE", Message: "name is required",
		})
	}
	if len(v) < 3 {
		errs = append(errs, &validation.Error{
			Field: "name", Code: "UNEXPECTED_LENGTH", Message: "name must be at least 36 chars",
		})
	}
	return errs
}
