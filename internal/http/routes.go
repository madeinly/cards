package http

import (
	coreModels "github.com/madeinly/core/models"
)

var Routes = []coreModels.Route{
	{
		Type:    "GET",
		Pattern: "/card",
		Handler: GetCard,
	},
	{
		Type:    "POST",
		Pattern: "/card",
		Handler: PostCreateCard,
	},
}
