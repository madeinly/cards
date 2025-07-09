package cards

import (
	coreModels "github.com/madeinly/core/models"
	"githube.com/madeinly/cards/internal/cmd"
)

var Feature = coreModels.FeaturePackage{
	Name: "cards",
	// Migration: UserMigration,
	Setup: setupCards,
	// Routes:    http.Routes,
	Cmd: cmd.Execute,
}

func setupCards() error {
	return nil
}
