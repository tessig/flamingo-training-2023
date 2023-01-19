package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"flamingo.me/flamingo/v3/core/healthcheck/domain/healthcheck"
	"go.opencensus.io/trace"

	"flamingo.me/training/src/openweather/application"
	"flamingo.me/training/src/openweather/domain"
)

type (
	// Adapter for openweather
	Adapter struct {
		apiClient *APIClient
	}

	weatherDto struct {
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Base string `json:"base"`
		Main struct {
			Temp     float64 `json:"temp"`
			Pressure int     `json:"pressure"`
			Humidity int     `json:"humidity"`
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
		} `json:"main"`
		Visibility int `json:"visibility"`
		Wind       struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Dt  int `json:"dt"`
		Sys struct {
			Type    int     `json:"type"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Country string  `json:"country"`
			Sunrise int     `json:"sunrise"`
			Sunset  int     `json:"sunset"`
		} `json:"sys"`
		ID   int    `json:"id"`
		Name string `json:"name"`
		Cod  int    `json:"cod"`
	}
)

var (
	// Check if we really implement the interface during compilation
	_ application.Service = new(Adapter)
	_ healthcheck.Status  = new(Adapter)
	// ErrNoWeather is returned if no weather data is available
	ErrNoWeather = errors.New("no weather data")
)

// Inject dependencies
func (a *Adapter) Inject(
	client *APIClient,
) *Adapter {
	a.apiClient = client

	return a
}

func (a *Adapter) Status() (alive bool, details string) {
	resp, err := a.apiClient.request(context.Background(), http.MethodGet, "weather?q=London", nil)
	if err != nil {
		return false, err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Sprintf("openweather API HTTP status is %d", resp.StatusCode)
	}

	return true, ""
}

// GetByCity returns the weather for the given city
func (a *Adapter) GetByCity(ctx context.Context, city string) (domain.Weather, error) {
	ctx, span := trace.StartSpan(ctx, "openweather/adapter/city")
	defer span.End()

	resp, err := a.apiClient.request(ctx, http.MethodGet, "weather?q="+city, nil)
	if err != nil {
		return domain.Weather{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Weather{}, err
	}

	dto := new(weatherDto)
	err = json.Unmarshal(body, dto)
	if err != nil {
		return domain.Weather{}, err
	}
	return mapDto(dto)
}

func mapDto(dto *weatherDto) (domain.Weather, error) {
	if len(dto.Weather) < 1 {
		return domain.Weather{}, ErrNoWeather
	}
	return domain.Weather{
		MainCharacter: dto.Weather[0].Main,
		Description:   dto.Weather[0].Description,
		IconCode:      dto.Weather[0].Icon,
		Temp:          int(dto.Main.Temp),
		Humidity:      dto.Main.Humidity,
		TempMin:       int(dto.Main.TempMin),
		TempMax:       int(dto.Main.TempMax),
		WindSpeed:     dto.Wind.Speed,
		Cloudiness:    dto.Clouds.All,
		LocationName:  dto.Name,
	}, nil
}
