package flows

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
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
	CardName     string  `json:"card_name"`
	CardId       string  `json:"card_id"`
	CardImage    string  `json:"card_image"`
	CardPriceMin float64 `json:"card_priceMin"`
	CardPriceMax float64 `json:"card_priceMax"`
	IncludeEn    bool    `json:"card_includeEn"`
	IncludeEs    bool    `json:"card_includeEs"`
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

	var anyColor int64
	if params.Colors == "" {
		anyColor = 1
	} else {

		anyColor = 0
	}

	list, cardCount, err := features.GetFilteredCards(ctx, features.GetFilteredCardsParams{
		CardName:     params.CardName,
		CardType:     params.CardType,
		CardFinish:   params.CardFinish,
		CardMv:       cardMv,
		CardPriceMin: cardPriceMin,
		CardPriceMax: cardPriceMax,
		MatchType:    params.MatchType,
		LangEn:       cardEn,
		LangEs:       cardEs,
		AnyColor:     anyColor,
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

	cardList := listUniqueCards(list)

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

func listUniqueCards(rows []appDB.GetFilteredCardsRow) []CardIndex {
	// temp map to aggregate while we scan
	m := make(map[string]*CardIndex) // key: NameEn

	for _, r := range rows {
		// already seen?  update aggregates
		if e, ok := m[r.NameEn]; ok {
			if r.Price < e.CardPriceMin {
				e.CardPriceMin = r.Price
			}
			if r.Price > e.CardPriceMax {
				e.CardPriceMax = r.Price
			}
			if r.Language == "English" {
				e.IncludeEn = true
			}
			if r.Language == "Spanish" {
				e.IncludeEs = true
			}
			continue
		}

		// first time we see this card
		m[r.NameEn] = &CardIndex{
			CardId:       r.ID,
			CardName:     r.NameEn,
			CardImage:    r.ImageUrl,
			CardPriceMin: r.Price,
			CardPriceMax: r.Price,
			IncludeEn:    r.Language == "English",
			IncludeEs:    r.Language == "Spanish",
		}
	}

	// map -> slice (stable order: alphabetical by name)
	out := make([]CardIndex, 0, len(m))
	for _, v := range m {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CardId < out[j].CardId })
	return out
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
