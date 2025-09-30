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

type RegisterCardParams struct {
	ID         string `json:"scryfall_id"`
	Vendor     string `json:"vendor"`
	Language   string `json:"language"`
	Finish     string `json:"finish"`
	Stock      string `json:"stock"`
	Visibility string `json:"visibility"`
}

func RegisterCard(ctx context.Context, params RegisterCardParams) error {

	card, _ := GetCardfromId(ctx, params.ID, params.Finish, params.Language)

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

	sku := strings.ToLower(params.Language) + "-" + params.Finish + "-" + card.SetCode + "-" + card.Number

	stock, _ := strconv.ParseInt(params.Stock, 10, 64)

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
		Rarity:     card.Rarity,
		Number:     card.Number,
		Finish:     params.Finish,
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
