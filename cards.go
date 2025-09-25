package cards

import (
	_ "embed"
	"fmt"

	"github.com/madeinly/cards/internal/flows"
	"github.com/madeinly/cards/internal/gateways/cmd"
	"github.com/madeinly/cards/internal/gateways/http"
	"github.com/madeinly/core"
)

//go:embed internal/drivers/sqlite/queries/app/migration.sql
var initialSchema string

var migration = core.Migration{
	Name:   "cards",
	Schema: initialSchema,
}

var Feature = core.FeaturePackage{
	Name:      "cards",
	Migration: migration,
	Setup:     setupCards,
	Routes:    http.Routes,
	Cmd:       cmd.Execute,
}

func setupCards(params map[string]string) error {

	var err error

	fmt.Println("updating the cards database")
	err = flows.UpdateCardsDB()

	if err != nil {
		return err
	}

	fmt.Println("updatinf the card prices database")
	err = flows.InitCardPrices()

	if err != nil {
		return err
	}

	return nil
}
