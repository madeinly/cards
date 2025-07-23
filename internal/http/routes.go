package http

import (
	coreModels "github.com/madeinly/core/models"
)

var Routes = []coreModels.Route{
	{
		Type:    "GET",
		Pattern: "/user",
		Handler: GetCard,
	},
}
