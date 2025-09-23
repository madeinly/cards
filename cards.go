package cards

import (
	_ "embed"
	"fmt"

	"github.com/madeinly/cards/internal/cmd"
	"github.com/madeinly/cards/internal/http"
	"github.com/madeinly/cards/internal/service"
	"github.com/madeinly/core"
)

//go:embed internal/schemas/initial_schema.sql
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
	err = service.UpdateCardsDB()

	if err != nil {
		return err
	}

	fmt.Println("updatinf the card prices database")
	err = service.InitCardPrices()

	if err != nil {
		return err
	}

	return nil
}
