package graphql

import (
	"context"

	"flamingo.me/flamingo/v3/core/auth"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/training/src/openweather/application"
	"flamingo.me/training/src/openweather/domain"
)

type (
	WeatherResolver struct {
		weatherService  application.Service
		identityService *auth.WebIdentityService
	}
)

// Inject dependencies
func (r *WeatherResolver) Inject(
	weatherService application.Service,
	identityService *auth.WebIdentityService,
) *WeatherResolver {
	r.weatherService = weatherService
	r.identityService = identityService

	return r
}

func (r *WeatherResolver) Openweather_Weather(ctx context.Context, city string) (*domain.Weather, error) {
	identity := r.identityService.Identify(ctx, web.RequestFromContext(ctx))

	weather, err := r.weatherService.GetByCity(ctx, city)
	if err != nil {
		return nil, err
	}

	if identity == nil {
		weather.Humidity = 0
	}

	return &weather, nil
}
