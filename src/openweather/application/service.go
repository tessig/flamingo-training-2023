package application

import (
	"context"

	"flamingo.me/training/src/openweather/domain"
)

type (
	Service interface {
		GetByCity(ctx context.Context, city string) (domain.Weather, error)
	}
)
