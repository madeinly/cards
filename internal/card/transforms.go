package card

import (
	"net/url"
)

func GetImageURL(cardID string) string {
	base, _ := url.Parse("https://cards.scryfall.io")

	url := base.JoinPath("normal", "front", string(cardID[0]), string(cardID[1]), cardID+".jpg")

	return url.String()

}
