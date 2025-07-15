package cards

import (
	_ "embed"

	coreModels "github.com/madeinly/core/models"
	"githube.com/madeinly/cards/internal/cmd"
	"githube.com/madeinly/cards/internal/service"
)

//go:embed internal/repository/queries/schemas/initial_schema.sql
var initialSchema string

var Feature = coreModels.FeaturePackage{
	Name:      "cards",
	Migration: coreModels.Migration{Name: "cards", Schema: initialSchema},
	Setup:     setupCards,
	// Routes:    http.Routes,
	Cmd: cmd.Execute,
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
