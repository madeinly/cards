package flows

import (
	"context"
	"errors"

	"github.com/madeinly/cards/internal/features"
)

func ListUniqueCards(ctx context.Context, cardName string) ([]string, error) {

	list, err := features.GetUniqueCards(ctx, cardName)

	if err != nil && errors.Is(err, features.ErrCardNotFound) {
		return nil, ErrResourceNotFound
	}

	if err != nil {
		return nil, ErrServerFailure
	}

	return list, nil
}
