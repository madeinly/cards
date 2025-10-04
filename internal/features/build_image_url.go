package features

import "net/url"

// this is uses the SCRYFALL ID
func BuildImageURL(scryfallId string) string {
	base, _ := url.Parse("https://cards.scryfall.io")

	url := base.JoinPath("normal", "front", string(scryfallId[0]), string(scryfallId[1]), scryfallId+".jpg")

	return url.String()

}
