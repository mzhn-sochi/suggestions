package usecases

import "suggestions/pkg/suggestions"

type (
	ShopsRepository interface {
	}

	ShopsUsecases interface {
		SuggestShop(lon float32, lat float32, results uint) ([]suggestions.Suggestion, error)
	}
)
