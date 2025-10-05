package features

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

type GetFilteredCardsParams struct {
	CardName     string
	CardType     string
	CardFinish   string
	CardMv       int64
	CardPriceMin int64
	CardPriceMax int64
	MatchType    string
	Colors       string
	AnyColor     int64
	ColorB       int64
	ColorG       int64
	ColorR       int64
	ColorU       int64
	ColorW       int64
	LangEn       int64
	LangEs       int64
	Limit        int64
	Offset       int64
}

func GetFilteredCards(ctx context.Context, params GetFilteredCardsParams) ([]appDB.GetFilteredCardsRow, int64, error) {

	db := core.DB()

	queryApp := appDB.New(db)

	filteredParams := appDB.GetFilteredCardsParams{
		CardType:     params.CardType,
		CardName:     params.CardName,
		LangEn:       params.LangEn,
		LangES:       params.LangEs,
		CardMv:       params.CardMv,
		CardFinish:   params.CardFinish,
		CardPriceMin: params.CardPriceMin,
		CardPriceMax: params.CardPriceMax,
		MatchType:    params.MatchType,
		CardColor:    params.Colors,
		AnyColor:     params.AnyColor,
		ColorB:       params.ColorB,
		ColorG:       params.ColorG,
		ColorR:       params.ColorR,
		ColorU:       params.ColorU,
		ColorW:       params.ColorW,
		Offset:       params.Offset,
		Limit:        params.Limit,
	}

	debugVal, _ := json.MarshalIndent(filteredParams, "", " ")
	fmt.Println(string(debugVal))

	list, _ := queryApp.GetFilteredCards(ctx, filteredParams)

	countParams := appDB.CountFilteredCardsParams{
		CardType:     params.CardType,
		CardName:     params.CardName,
		LangEn:       params.LangEn,
		LangES:       params.LangEs,
		CardMv:       params.CardMv,
		CardFinish:   params.CardFinish,
		CardPriceMin: params.CardPriceMin,
		CardPriceMax: params.CardPriceMax,
		MatchType:    params.MatchType,
		CardColor:    params.Colors,
		AnyColor:     params.AnyColor,
		ColorB:       params.ColorB,
		ColorG:       params.ColorG,
		ColorR:       params.ColorR,
		ColorU:       params.ColorU,
		ColorW:       params.ColorW,
	}

	cardCount, err := queryApp.CountFilteredCards(ctx, countParams)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, 0, ErrCardNotFound
	}

	if err != nil {
		return nil, 0, err
	}

	return list, cardCount, nil

}
