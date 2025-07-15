package card

import (
	"net/url"
	"strings"
)

func TransformColors(strColors string) []string {

	var colors []string

	if strColors != "" {
		colors = strings.Split(strColors, ", ")
	} else {
		colors = []string{"colorless"}
	}

	return colors
}

func TransformTypes(strTypes string) []string {

	types := strings.Split(strTypes, ", ")

	return types
}

func GetImageURL(cardID string) string {
	base, _ := url.Parse("https://cards.scryfall.io")

	url := base.JoinPath("normal", "front", string(cardID[0]), string(cardID[1]), cardID+".jpg")

	return url.String()

}
