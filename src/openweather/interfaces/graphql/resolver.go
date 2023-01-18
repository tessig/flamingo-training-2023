package graphql

import (
	"context"

	"flamingo.me/training/src/openweather/domain"
)

type (
	WeatherResolver struct {
		weatherService domain.Service
	}
)

// Inject dependencies
func (r *WeatherResolver) Inject(
	weatherService domain.Service,
) *WeatherResolver {
	r.weatherService = weatherService

	return r
}

func (r *WeatherResolver) Openweather_Weather(ctx context.Context, city string) (*domain.Weather, error) {
	weather, err := r.weatherService.GetByCity(ctx, city)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
