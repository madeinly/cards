package flows

import (
	"context"
	"database/sql"
	"strconv"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

type GetDashboardCardsParams struct {
	SetCode  string
	CardName string
	Page     string
	Limit    string
}

func GetDashboardCards(ctx context.Context, params GetDashboardCardsParams) ([]Card, error) {

	db := core.DB()

	qApp := appDB.New(db)

	page, err := strconv.ParseInt(params.Page, 10, 64)
	if err != nil {
		return nil, err
	}

	limit, err := strconv.ParseInt(params.Limit, 10, 64)
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit

	getCardsParams := appDB.GetCardsWithPriceParams{
		SetCode: params.SetCode,
		Name:    params.CardName,
		Offset:  offset,
		Limit:   limit,
	}

	repoCards, err := qApp.GetCardsWithPrice(ctx, getCardsParams)

	if err != nil && err == sql.ErrNoRows {

		return []Card{}, nil
	}

	if err != nil {
		return []Card{}, err
	}

	var cards []Card

	for _, repoCard := range repoCards {

		cards = append(cards, Card{
			ID:        repoCard.ID,
			NameEN:    repoCard.NameEn,
			NameES:    repoCard.NameEs,
			ImageURL:  repoCard.UrlImage,
			SetCode:   repoCard.SetCode,
			SetName:   repoCard.SetName,
			ManaValue: repoCard.ManaValue,
			Colors:    repoCard.Colors,
			Types:     repoCard.Types,
			Price:     repoCard.Price,
			Stock:     repoCard.Stock,
		})

	}

	return cards, nil
}
