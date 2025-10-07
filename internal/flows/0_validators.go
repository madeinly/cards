package flows

import "github.com/madeinly/core"

// This are all the parameters that enter to the application

var (
	ScryfallId   string // 36 characters (uuid google package)
	CardFinish   string // foil, normal, etched
	CardLanguage string // English, Spanish
	CardSetCode  string // 3 characters alphanumeric
	CardName     string // alpha numeric include space
	CardId       string // 36 character alpha numeric and - uuid google package
	Page         string // must be able to convert into integer
	Limit        string // must be able to convert into integer
	CardType     string // must be one of: Creature, Sorcery, Enchantment, Instant, Artifact, Land, Planeswalker, Vanguard, Summon, Conspiracy, Plane, Scheme, Battle, Phenomenon, Hero
	CardMv       string // must be able to convert to integer from 0 to 100
	CardPriceMax string // must be able to convert to integer from 0 to 1000
	CardPriceMin string // must be able to convert to integer from 0 to 1000
	MatchType    string // one of "tight", "loose"
	Colors       string // can be empty, a string that include only (in capital letters): B, G, R, U, W
	Vendor       string //
	Visibility   string // 1 or 0
)

type CardsValidator struct{}

func (c *CardsValidator) ScryfallId(ScryfallId string) core.ErrorBag {

	errBag := core.Validate()

	if len(ScryfallId) != 36 {

		errBag.Add("ScryfallId", "invalid_length", "scryfallId is exactly 36 characters")
	}

	return *errBag

}
