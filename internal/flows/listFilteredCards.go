package flows

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/madeinly/cards/internal/features/cards"
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
	CardId       string  `json:"card_id"`
	CardImage    string  `json:"card_image"`
	CardPriceMin float64 `json:"card_priceMin"`
	CardPriceMax float64 `json:"card_priceMax"`
	IncludeEn    bool    `json:"card_includeEn"`
	IncludeEs    bool    `json:"card_includeEs"`
}

type CardIndexPage struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
	Cards map[string]CardIndex
}

func ListFilteredCards(ctx context.Context, params ListFilteredCardsParams) (CardIndexPage, error) {

	limit, _ := strconv.ParseInt(params.Limit, 10, 64)
	page, _ := strconv.ParseInt(params.Page, 10, 64)
	offset := limit * (page - 1)

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

	//ORDER OF COLORS IN MTGJSON: B, G, R, U, W

	list, err := cards.GetFilteredCards(ctx, cards.GetFilteredCardsParams{
		CardName:     params.CardName,
		CardType:     params.CardType,
		CardFinish:   params.CardFinish,
		CardMv:       cardMv,
		CardPriceMin: cardPriceMin,
		CardPriceMax: cardPriceMax,
		MatchType:    params.MatchType,
		Colors:       params.Colors,
		LangEn:       cardEn,
		LangEs:       cardEs,
		Limit:        limit,
		Offset:       offset,
	})

	if err != nil && errors.Is(err, cards.ErrCardNotFound) {
		return CardIndexPage{}, ErrResourceNotFound
	}

	if err != nil {
		fmt.Println(err.Error())

		return CardIndexPage{}, ErrServerFailure

	}

	cardList := make(map[string]CardIndex)

	catalog := CardIndexPage{
		Limit: limit,
		Page:  page,
		Total: int64(len(list)),
		Cards: cardList,
	}

	for _, item := range list {
		idx, ok := cardList[item.NameEn]
		if !ok {
			idx = CardIndex{
				CardId:       item.ID,
				CardImage:    item.ImageUrl.String,
				CardPriceMin: item.Price,
				CardPriceMax: item.Price,
				IncludeEn:    item.Language == "English",
				IncludeEs:    item.Language == "Spanish",
			}
		} else {
			if item.Price < idx.CardPriceMin {
				idx.CardPriceMin = item.Price
			}
			if item.Price > idx.CardPriceMax {
				idx.CardPriceMax = item.Price
			}
			// accumulate languages
			if item.Language == "English" {
				idx.IncludeEn = true
			}
			if item.Language == "Spanish" {
				idx.IncludeEs = true
			}
		}
		cardList[item.NameEn] = idx
	}

	return catalog, nil

}
