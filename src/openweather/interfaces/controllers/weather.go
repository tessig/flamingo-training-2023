package controllers

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/training/src/openweather/application"
	"flamingo.me/training/src/openweather/domain"
)

type (
	WeatherController struct {
		responder      *web.Responder
		weatherService application.Service
	}
	viewData struct {
		Weather domain.Weather
	}
)

// Inject dependencies
func (w *WeatherController) Inject(
	responder *web.Responder,
	weatherServcice application.Service,
) *WeatherController {
	w.responder = responder
	w.weatherService = weatherServcice

	return w
}

// Weather for city
func (w *WeatherController) Weather(ctx context.Context, r *web.Request) web.Result {
	city := r.Params["city"]

	weather, err := w.weatherService.GetByCity(ctx, city)
	if err != nil {
		return w.responder.ServerError(err)
	}

	return w.responder.Render("weather/weather", &viewData{Weather: weather})
}
