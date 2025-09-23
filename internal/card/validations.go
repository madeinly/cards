package card

import (
	"slices"
	"strings"

	"github.com/madeinly/core"
)

func IdRules(v string) []*core.Error {

	var errs []*core.Error
	if v == "" {
		errs = append(errs, &core.Error{
			Field: "card_name", Code: "EMPTY_VALUE", Message: "card_name is required",
		})
	}
	if len(v) < 3 {
		errs = append(errs, &core.Error{
			Field: "card_name", Code: "UNEXPECTED_LENGTH", Message: "card_name must be at least 36 chars",
		})
	}
	return errs
}

func VendorRules(v string) []*core.Error {
	// var errs []*models.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &models.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func LanguageRules(v string) []*core.Error {
	// var errs []*models.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &models.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func StockRules(v string) []*core.Error {
	// var errs []*models.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &models.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func FinishRules(v string) []*core.Error {
	var errs []*core.Error

	allowed := []string{"foil", "normal", "etched"}

	ok := slices.Contains(allowed, v)

	if !ok {
		errs = append(errs, &core.Error{
			Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
		})
	}

	return errs
}

func ConditionRules(v string) []*core.Error {
	// var errs []*models.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &models.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func VisibilityRules(v string) []*core.Error {
	// var errs []*models.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &models.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func SetCodeRules(v string) []*core.Error {
	// var errs []*models.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &models.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}

func CardNameRules(v string) []*core.Error {
	// var errs []*models.Error

	// allowed := []string{"foil", "normal", "etched"}

	// ok := slices.Contains(allowed, v)

	// if !ok {
	// 	errs = append(errs, &models.Error{
	// 		Field: "card_finish", Code: "UNEXPECTED_VALUE", Message: "must be one of: " + strings.Join(allowed, ", "),
	// 	})
	// }

	return nil
}
