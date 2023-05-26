package infrastructure

import (
	"context"
	"errors"

	"flamingo.me/training/src/openweather/domain"
)

type (
	FakeWeather struct {
	}
)

func (f *FakeWeather) GetByCity(ctx context.Context, city string) (domain.Weather, error) {
	if city == "error" {
		return domain.Weather{}, errors.New("no weather available")
	}

	return domain.Weather{
		MainCharacter:       "fake character",
		Description:         "light rain",
		IconCode:            "09d",
		Temp:                15,
		Humidity:            80,
		TempMin:             9,
		TempMax:             23,
		WindSpeed:           4.3,
		Cloudiness:          75,
		LocationName:        city,
		LocationCountryCode: "DE",
	}, nil
}
