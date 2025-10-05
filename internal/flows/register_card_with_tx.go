package flows

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
	"github.com/madeinly/cards/internal/features"
)

func RegisterCardTx(ctx context.Context, tx *sql.Tx, params RegisterCardParams) error {

	card, err := GetCardFromIDTx(ctx, tx, params.ID, params.Finish, params.Language)

	if err != nil && err != sql.ErrNoRows {

		return err
	}

	qApp := appDB.New(tx)

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

	sku := strings.ToLower(params.Language) + "-" + params.Finish + "-" + card.SetCode + "-" + card.Number

	stock, err := strconv.ParseInt(params.Stock, 10, 64)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = qApp.CreateCard(ctx, appDB.CreateCardParams{
		ID:         card.ID,
		NameEn:     card.NameEN,
		NameEs:     card.NameES,
		Sku:        sku,
		UrlImage:   card.ImageURL,
		SetName:    card.SetName,
		SetCode:    card.SetCode,
		ManaValue:  card.ManaValue,
		Colors:     card.Colors,
		Types:      card.Types,
		Finish:     params.Finish,
		Rarity:     card.Rarity,
		Number:     card.Number,
		HasVendor:  hasVendor,
		Language:   params.Language,
		Visibility: params.Visibility,
		ImagePath:  sql.NullString{Valid: false},
		ImageUrl:   sql.NullString{Valid: false},
		Stock:      stock,
	})

	if err != nil {
		return err
	}

	return nil

}

func GetCardFromIDTx(ctx context.Context, tx *sql.Tx, cardScryFallID string, finish string, language string) (Card, error) {

	cardsDB := features.GetCardsDB()

	qCards := mtgDB.New(cardsDB)

	repoCard, err := qCards.GetCard(ctx, cardScryFallID)

	if err != nil {
		return Card{}, fmt.Errorf("could not find the card: %w", err)
	}

	nameES, err := qCards.GetCardNameES(ctx, sql.NullString{Valid: true, String: repoCard.Uuid})

	if err != nil && err != sql.ErrNoRows {
		nameES = sql.NullString{String: ""}
	}

	qApp := appDB.New(tx)

	stock, err := qApp.GetCardStockById(ctx, appDB.GetCardStockByIdParams{
		ID:       repoCard.Uuid,
		Language: language,
		Finish:   finish,
	})

	if err != nil && err != sql.ErrNoRows {
		return Card{}, fmt.Errorf("there was an error getting the stock")
	}

	if err == sql.ErrNoRows {
		stock = 0
	}

	price, err := qApp.GetPrice(ctx, appDB.GetPriceParams{
		CardID: repoCard.Uuid,
		Finish: finish,
	})

	if err != nil {
		return Card{}, nil
	}

	return Card{
		CardBase: CardBase{
			ID:        repoCard.Uuid,
			NameEN:    repoCard.Name,
			NameES:    nameES.String,
			ImageURL:  features.BuildImageURL(cardScryFallID),
			SetCode:   repoCard.Setcode,
			SetName:   repoCard.Setname,
			ManaValue: int64(repoCard.Manavalue),
			Number:    repoCard.Number,
			Rarity:    repoCard.Rarity,
			Colors:    repoCard.Colors,
			Types:     repoCard.Types,
			Price:     price,
			Stock:     stock,
		},
	}, nil
}
