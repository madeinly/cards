package features

import (
	"context"
	"database/sql"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
)

type CreateCardParams struct {
	ID         string
	NameEn     string
	NameEs     string
	Sku        string
	SetName    string
	SetCode    string
	ManaValue  int64
	Colors     string
	Types      string
	Finish     string
	Rarity     string
	Number     string
	HasVendor  bool
	Language   string
	Visibility int64
	ImagePath  string
	ImageUrl   string
	Stock      int64
}

func CreateCard(ctx context.Context, tx *sql.Tx, params CreateCardParams) error {

	qApp := appDB.New(tx)

	err := qApp.CreateCard(ctx, appDB.CreateCardParams{
		ID:         params.ID,
		NameEn:     params.NameEn,
		NameEs:     params.NameEs,
		Sku:        params.Sku,
		SetName:    params.SetName,
		SetCode:    params.SetCode,
		ManaValue:  params.ManaValue,
		Colors:     params.Colors,
		Types:      params.Colors,
		Finish:     params.Finish,
		Rarity:     params.Rarity,
		Number:     params.Number,
		HasVendor:  params.HasVendor,
		Language:   params.Language,
		Visibility: params.Visibility,
		ImagePath:  sql.NullString{String: params.ImagePath, Valid: true},
		ImageUrl:   params.ImageUrl,
		Stock:      params.Stock,
	})

	if err != nil {
		return err
	}

	return nil

}
