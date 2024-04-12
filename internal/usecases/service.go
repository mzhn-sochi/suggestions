package usecases

import (
	"sort"
	"suggestions/pkg/suggestions"
)

type usecases struct {
	proxy suggestions.Suggestions
}

func NewUsecases(suggestionsProxy suggestions.Suggestions) ShopsUsecases {
	return &usecases{proxy: suggestionsProxy}
}

func (u *usecases) SuggestShop(lon float32, lat float32, results uint) ([]suggestions.Suggestion, error) {
	suggest, err := u.proxy.GetSuggestions("Супермаркет", lon, lat, 10)
	if err != nil {
		return nil, err
	}

	sort.Slice(suggest, func(i, j int) bool {
		return suggest[i].Distance.Value < suggest[j].Distance.Value
	})

	return suggest[:results], nil
}
