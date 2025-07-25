package cards

import (
	_ "embed"

	"github.com/madeinly/cards/internal/cmd"
	"github.com/madeinly/cards/internal/http"
	"github.com/madeinly/cards/internal/service"
	coreModels "github.com/madeinly/core/models"
)

//go:embed internal/schemas/initial_schema.sql
var initialSchema string

var Feature = coreModels.FeaturePackage{
	Name:      "cards",
	Migration: coreModels.Migration{Name: "cards", Schema: initialSchema},
	Setup:     setupCards,
	Routes:    http.Routes,
	Cmd:       cmd.Execute,
}

func setupCards() error {

	var err error

	err = service.UpdateCardsDB()

	if err != nil {
		return err
	}

	err = service.InitCardPrices()

	if err != nil {
		return err
	}

	return nil
}
