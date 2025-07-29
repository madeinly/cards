package card

import (
	"slices"
	"strings"

	"github.com/madeinly/core/validation"
)

func IdRules(v string) []*validation.Error {

	var errs []*validation.Error
	if v == "" {
		errs = append(errs, &validation.Error{
			Field: "card_name", Code: "EMPTY_VALUE", Message: "card_name is required",
		})
	}
	if len(v) < 3 {
		errs = append(errs, &validation.Error{
			Field: "card_name", Code: "UNEXPECTED_LENGTH", Message: "card_name must be at least 36 chars",
		})
	}
	return errs
}

func VendorRules(v string) []*validation.Error {
	// var errs []*validation.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &validation.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func LanguageRules(v string) []*validation.Error {
	// var errs []*validation.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &validation.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func StockRules(v string) []*validation.Error {
	// var errs []*validation.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &validation.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func FinishRules(v string) []*validation.Error {
	var errs []*validation.Error

	allowed := []string{"foil", "normal", "etched"}

	ok := slices.Contains(allowed, v)

	if !ok {
		errs = append(errs, &validation.Error{
			Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
		})
	}

	return errs
}

func ConditionRules(v string) []*validation.Error {
	// var errs []*validation.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &validation.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func VisibilityRules(v string) []*validation.Error {
	// var errs []*validation.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &validation.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func SetCodeRules(v string) []*validation.Error {
	// var errs []*validation.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &validation.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func CardNameRules(v string) []*validation.Error {
	// var errs []*validation.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &validation.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}
