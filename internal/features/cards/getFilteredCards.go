package cards

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
	LangEn       int64
	LangEs       int64
	Limit        int64
	Offset       int64
}

func GetFilteredCards(ctx context.Context, params GetFilteredCardsParams) ([]appDB.GetFilteredCardsRow, error) {

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
		Offset:       params.Offset,
		Limit:        params.Limit,
	}

	debugVal, _ := json.MarshalIndent(filteredParams, "", " ")
	fmt.Println(string(debugVal))

	list, err := queryApp.GetFilteredCards(ctx, filteredParams)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, ErrCardNotFound
	}

	if err != nil {
		return nil, err
	}

	return list, nil

}
