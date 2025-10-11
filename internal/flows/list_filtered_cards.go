package flows

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/madeinly/cards/internal/features"
)

type ListFilteredCardsParams struct {
	CardName     string
	CardType     string
	CardFinish   string
	CardMv       string
	CardPriceMax string
	CardPriceMin string
	CardEn       bool
	CardES       bool
	MatchType    string
	Colors       string
	Limit        string
	Page         string
}

type CardIndex struct {
	CardName   string  `json:"card_name"`
	CardId     string  `json:"card_id"`
	CardImage  string  `json:"card_image"`
	CardPrice  float64 `json:"card_price"`
	Language   string  `json:"card_language"`
	Finish     string  `json:"card_finish"`
	Rarity     string  `json:"card_rarity"`
	ManaValue  int64   `json:"card_manaValue"`
	CardType   string  `json:"card_type"`
	CardColors string  `json:"card_colors"`
}

type CardIndexPage struct {
	Page  int64 `json:"card_pages"`
	Total int64 `json:"card_total"`
	Cards []CardIndex
}

type colorMatch struct {
	B int64
	G int64
	R int64
	U int64
	W int64
	C int64
}

func ListFilteredCards(ctx context.Context, params ListFilteredCardsParams) (CardIndexPage, error) {

	if params.Limit == "" || params.Limit == "-1" {
		params.Limit = "100"
	}

	limit, _ := strconv.ParseInt(params.Limit, 10, 64)
	page, _ := strconv.ParseInt(params.Page, 10, 64)
	var offset int64

	if page == 1 {
		offset = 0
	} else {
		offset = limit * (page - 1)
	}

	if params.CardMv == "" {
		params.CardMv = "-1"
	}

	cardMv, _ := strconv.ParseInt(params.CardMv, 10, 64)

	cardPriceMin, _ := strconv.ParseInt(params.CardPriceMin, 10, 64)
	cardPriceMax, _ := strconv.ParseInt(params.CardPriceMax, 10, 64)

	var cardEn int64
	if params.CardEn {
		cardEn = 1
	}

	var cardEs int64
	if params.CardES {
		cardEs = 1
	}

	colorMatch := parseColorString(params.Colors)

	var colorless int64
	if params.Colors == "C" || params.Colors == "c" || params.Colors == "" {
		colorless = 1
	} else {
		colorless = 0
	}

	var matchType string
	if params.MatchType == "" || params.MatchType == "loose" {
		matchType = "loose"
	} else {
		matchType = "tight"
	}

	list, cardCount, err := features.GetFilteredCards(ctx, features.GetFilteredCardsParams{
		CardName:     params.CardName,
		CardType:     params.CardType,
		CardFinish:   params.CardFinish,
		CardMv:       cardMv,
		CardPriceMin: cardPriceMin,
		CardPriceMax: cardPriceMax,
		MatchType:    matchType,
		LangEn:       cardEn,
		LangEs:       cardEs,
		Colorless:    colorless,
		Colors:       params.Colors,
		ColorB:       colorMatch.B,
		ColorG:       colorMatch.G,
		ColorR:       colorMatch.R,
		ColorU:       colorMatch.U,
		ColorW:       colorMatch.W,
		Limit:        limit,
		Offset:       offset,
	})

	if err != nil && errors.Is(err, features.ErrCardNotFound) {
		return CardIndexPage{}, ErrResourceNotFound
	}

	if err != nil {
		fmt.Println(err.Error())

		return CardIndexPage{}, ErrServerFailure

	}

	var cardList []CardIndex

	for _, item := range list {
		cardList = append(cardList, CardIndex{
			CardName:   item.NameEn,
			CardId:     item.ID,
			CardImage:  item.ImageUrl,
			CardPrice:  item.Price,
			Language:   item.Language,
			Finish:     item.Finish,
			Rarity:     item.Rarity,
			ManaValue:  item.ManaValue,
			CardType:   item.Types,
			CardColors: item.Colors,
		})
	}

	var totalPages int64
	if limit == -1 {
		totalPages = 1
	} else {
		totalPages = (cardCount + limit - 1) / limit //ceiling trick
	}

	catalog := CardIndexPage{
		Page:  totalPages,
		Total: cardCount,
		Cards: cardList,
	}

	return catalog, nil

}

func parseColorString(s string) colorMatch {
	var m colorMatch
	for _, c := range s {
		switch c {
		case 'B', 'b':
			m.B = 1
		case 'G', 'g':
			m.G = 1
		case 'R', 'r':
			m.R = 1
		case 'U', 'u':
			m.U = 1
		case 'W', 'w':
			m.W = 1
		}
	}
	return m
}
