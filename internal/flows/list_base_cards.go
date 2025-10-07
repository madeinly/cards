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

type CardsPage struct {
	Limit int64
	Page  int64
	Total int64
	Cards []Card
}

func GetDashboardCards(ctx context.Context, params GetDashboardCardsParams) (CardsPage, error) {

	db := core.DB()

	qApp := appDB.New(db)

	page, err := strconv.ParseInt(params.Page, 10, 64)
	if err != nil {
		return CardsPage{}, err
	}

	limit, err := strconv.ParseInt(params.Limit, 10, 64)
	if err != nil {
		return CardsPage{}, err
	}

	var offset int64

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	getCardsParams := appDB.GetCardsWithPriceParams{
		SetCode: params.SetCode,
		Name:    params.CardName,
		Offset:  offset,
		Limit:   limit,
	}

	repoCards, err := qApp.GetCardsWithPrice(ctx, getCardsParams)

	if err != nil && err == sql.ErrNoRows {

		return CardsPage{}, ErrResourceNotFound
	}

	if err != nil {
		return CardsPage{}, err
	}

	cardsCount, _ := qApp.CountCardsWithPrice(ctx, appDB.CountCardsWithPriceParams{
		SetCode: params.SetCode,
		Name:    params.CardName,
	})

	totalPages := (cardsCount + limit - 1) / limit //ceiling trick

	var cards []Card

	for _, repoCard := range repoCards {

		cards = append(cards, Card{
			CardBase: CardBase{
				ID:        repoCard.ID,
				NameEN:    repoCard.NameEn,
				NameES:    repoCard.NameEs,
				ImageURL:  repoCard.ImageUrl,
				SetCode:   repoCard.SetCode,
				SetName:   repoCard.SetName,
				ManaValue: repoCard.ManaValue,
				Colors:    repoCard.Colors,
				Types:     repoCard.Types,
				Number:    repoCard.Number,
				Rarity:    repoCard.Rarity,
				Price:     repoCard.Price,
				Stock:     repoCard.Stock,
			},
			Language:  repoCard.Language,
			Finish:    repoCard.Finish,
			HasVendor: repoCard.HasVendor,
		})

	}

	return CardsPage{
		Page:  totalPages,
		Cards: cards,
		Total: cardsCount,
	}, nil
}
