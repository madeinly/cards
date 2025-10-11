package http

import (
	"net/http"

	"github.com/madeinly/core"
)

var Routes = []core.Route{
	{
		Type:    "GET",
		Pattern: "/cards/unique",
		Handler: http.HandlerFunc(GetCard),
	},
	{
		Type:    "GET",
		Pattern: "/cards/names",
		Handler: http.HandlerFunc(ListCardNames),
	},
	{
		Type:    "POST",
		Pattern: "/cards",
		Handler: http.HandlerFunc(CreateCard),
	},
	{
		Type:    "POST",
		Pattern: "/cards/bulk",
		Handler: http.HandlerFunc(BulkCreate),
	},
	{
		Type:    "GET",
		Pattern: "/cards",
		Handler: http.HandlerFunc(ListCards),
	},
	{
		Type:    "GET",
		Pattern: "/cards/sets",
		Handler: http.HandlerFunc(GetSets),
	},
	{
		Type:    "PATCH",
		Pattern: "/cards",
		Handler: http.HandlerFunc(UpdateCardStock),
	},
	{
		Type:    "GET",
		Pattern: "/cards/available",
		Handler: http.HandlerFunc(ListCardsAvailable),
	},
	{
		Type:    "GET",
		Pattern: "/cards/filtered",
		Handler: http.HandlerFunc(GetFilteredCards),
	},
}
