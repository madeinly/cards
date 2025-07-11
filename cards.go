package cards

import (
	_ "embed"

	coreModels "github.com/madeinly/core/models"
	"githube.com/madeinly/cards/internal/cmd"
)

//go:embed internal/repository/queries/initial_schema.sql
var initialSchema string

var Feature = coreModels.FeaturePackage{
	Name:      "cards",
	Migration: coreModels.Migration{Name: "cards", Schema: initialSchema},
	Setup:     setupCards,
	// Routes:    http.Routes,
	Cmd: cmd.Execute,
}

func setupCards() error {
	return nil
}
