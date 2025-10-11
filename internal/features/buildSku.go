package features

import "strings"

type SkuParams struct {
	Language string
	Finish   string
	SetCode  string
	Number   string
}

func BuildCardSku(params SkuParams) string {

	sku := strings.ToLower(params.Language) + "-" + params.Finish + "-" + params.SetCode + "-" + params.Number

	return sku
}
