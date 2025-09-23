package http

import (
	"net/http"

	"github.com/madeinly/core"
)

var Routes = []core.Route{
	{
		Type:    "GET",
		Pattern: "/card",
		Handler: http.HandlerFunc(GetCard),
	},
	{
		Type:    "POST",
		Pattern: "/card",
		Handler: http.HandlerFunc(PostCreateCard),
	},
	{
		Type:    "POST",
		Pattern: "/card/bulk",
		Handler: http.HandlerFunc(BulkCreate),
	},
	{
		Type:    "GET",
		Pattern: "/cards",
		Handler: http.HandlerFunc(GetDashboardCards),
	},
	{
		Type:    "GET",
		Pattern: "/card/sets",
		Handler: http.HandlerFunc(GetSets),
	},
}
