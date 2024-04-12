package router

import (
	"context"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log/slog"
	"suggestions/api/shop_suggestions"
	"suggestions/internal/usecases"
	"suggestions/pkg/suggestions"
)

type ShopSuggestionsService struct {
	shop_suggestions.UnimplementedShopSuggestionsServiceServer
	uc  usecases.ShopsUsecases
	log *slog.Logger
}

func NewGRPCServer(grpcServer *grpc.Server, uc usecases.ShopsUsecases, logger *slog.Logger) shop_suggestions.ShopSuggestionsServiceServer {
	srv := ShopSuggestionsService{
		uc:  uc,
		log: logger,
	}

	shop_suggestions.RegisterShopSuggestionsServiceServer(grpcServer, &srv)
	reflection.Register(grpcServer)

	return &srv
}

func (s *ShopSuggestionsService) GetSuggestion(ctx context.Context, options *shop_suggestions.SuggestionOptions) (*shop_suggestions.ShopSuggestionsResponse, error) {
	s.log.InfoContext(ctx, "new suggestion request has been received",
		slog.Float64("lat", float64(options.Lat)),
		slog.Float64("lon", float64(options.Lon)),
		slog.Uint64("results", uint64(options.Count)),
	)
	sugg, err := s.uc.SuggestShop(options.Lon, options.Lat, uint(options.Count))
	if err != nil {
		s.log.ErrorContext(ctx, "unable to get suggestions", err)
		return nil, status.Errorf(codes.InvalidArgument, "the suggestion cannot be executed")
	}
	s.log.DebugContext(ctx, "suggestion results", slog.Any("sugg", sugg), slog.Int("count", len(sugg)))

	convSuggestions := lo.Map(sugg, func(item suggestions.Suggestion, index int) *shop_suggestions.Suggestion {
		return &shop_suggestions.Suggestion{
			Title:    item.Title.Text,
			Subtitle: item.Subtitle.Text,
			Distance: item.Distance.Text,
		}
	})

	return &shop_suggestions.ShopSuggestionsResponse{
		Suggestions: convSuggestions,
	}, nil
}
