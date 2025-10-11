package flows

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

func RegisterCard(ctx context.Context, params RegisterCardParams) error {

	card, _ := GetCardfromId(ctx, params.ScryfallId, params.Finish, params.Language)

	db := core.DB()

	qApp := appDB.New(db)

	exists, err := qApp.CardExists(ctx, appDB.CardExistsParams{
		ID:       params.ScryfallId,
		Finish:   params.Finish,
		Language: params.Language,
	})

	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("the element already exist")
	}

	hasVendor, err := qApp.GetCardHasVendorById(ctx, params.ScryfallId)

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

	stock, _ := strconv.ParseInt(params.Stock, 10, 64)
	visibility, _ := strconv.ParseInt(params.Visibility, 10, 64)

	err = qApp.CreateCard(ctx, appDB.CreateCardParams{
		ID:         card.ID,
		NameEn:     card.NameEN,
		NameEs:     card.NameES,
		Sku:        sku,
		SetName:    card.SetName,
		SetCode:    card.SetCode,
		ManaValue:  card.ManaValue,
		Colors:     card.Colors,
		Types:      card.Types,
		Rarity:     card.Rarity,
		Number:     card.Number,
		Finish:     params.Finish,
		HasVendor:  hasVendor,
		Language:   params.Language,
		Visibility: visibility,
		ImagePath:  sql.NullString{Valid: false},
		ImageUrl:   card.ImageURL,
		Stock:      stock,
	})

	if err != nil {
		return err
	}

	return nil

}
