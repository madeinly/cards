package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/database"
	appDB "github.com/madeinly/cards/internal/sqlc/app"
	mtgDB "github.com/madeinly/cards/internal/sqlc/cards"
	"github.com/madeinly/core"
)

func GetCardFromID(ctx context.Context, cardScryFallID string, finish string, language string) (card.Card, error) {

	cardsDB, err := database.GetCardsDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	qCards := mtgDB.New(cardsDB)

	fmt.Println("value on service:", cardScryFallID)
	repoCard, err := qCards.GetCard(ctx, cardScryFallID)

	if err != nil {
		return card.Card{}, fmt.Errorf("could not find the card: %w", err)
	}

	nameES, err := qCards.GetCardNameES(ctx, sql.NullString{Valid: true, String: repoCard.Uuid})

	if err != nil && err != sql.ErrNoRows {
		nameES = sql.NullString{String: ""}
	}

	db := core.DB()

	qApp := appDB.New(db)

	stock, err := qApp.GetCardStockById(ctx, appDB.GetCardStockByIdParams{
		ID:       repoCard.Uuid,
		Language: language,
		Finish:   finish,
	})

	if err != nil && err != sql.ErrNoRows {
		return card.Card{}, fmt.Errorf("there was an error getting the stock")
	}

	if err == sql.ErrNoRows {
		stock = 0
	}

	price, err := qApp.GetPrice(ctx, appDB.GetPriceParams{
		CardID: repoCard.Uuid,
		Finish: finish,
	})

	if err != nil {
		return card.Card{}, fmt.Errorf("there was an error getting the price")
	}

	return card.Card{
		ID:        repoCard.Uuid,
		NameEN:    repoCard.Name,
		NameES:    nameES.String,
		ImageURL:  card.GetImageURL(cardScryFallID),
		SetCode:   repoCard.Setcode,
		SetName:   repoCard.Setname,
		ManaValue: int64(repoCard.Manavalue),
		Number:    repoCard.Number,
		Rarity:    repoCard.Rarity,
		Colors:    repoCard.Colors,
		Types:     repoCard.Types,
		Price:     price,
		Stock:     stock,
	}, nil
}

func RegisterCard(ctx context.Context, params RegisterCardParams) error {

	card, err := GetCardFromID(ctx, params.ID, params.Finish, params.Language)

	if err != nil && err != sql.ErrNoRows {

		return err
	}

	db := core.DB()

	qApp := appDB.New(db)

	exists, err := qApp.CardExists(ctx, appDB.CardExistsParams{
		ID:       params.ID,
		Finish:   params.Finish,
		Language: params.Language,
	})

	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("the element already exist")
	}

	hasVendor, err := qApp.GetCardHasVendorById(ctx, params.ID)

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("there was a problem getting the hasVendor")
	}

	if err == sql.ErrNoRows {
		hasVendor = false
	}

	hasVendor = hasVendor || params.Vendor != ""

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("there was a problem getting the stock")
	}

	sku := params.Language + params.Finish + card.SetCode + card.Number

	err = qApp.CreateCard(ctx, appDB.CreateCardParams{
		ID:         params.ID,
		NameEs:     card.NameES,
		NameEn:     card.NameEN,
		Sku:        sku,
		UrlImage:   card.ImageURL,
		SetName:    card.SetName,
		SetCode:    card.SetCode,
		ManaValue:  card.ManaValue,
		Colors:     card.Colors,
		Types:      card.Types,
		Finish:     params.Finish,
		HasVendor:  hasVendor,
		Language:   params.Language,
		Visibility: params.Visibility,
		ImagePath:  sql.NullString{Valid: false},
		ImageUrl:   sql.NullString{Valid: false},
		Stock:      params.Stock,
	})

	if err != nil {
		return err
	}

	return nil

}

func RegisterBulk(ctx context.Context, filename string) error {

	return nil

}
