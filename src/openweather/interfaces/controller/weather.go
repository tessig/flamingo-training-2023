package controller

import (
	"context"

	"flamingo.me/flamingo/v3/core/auth"
	"flamingo.me/flamingo/v3/core/auth/oauth"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/training/src/openweather/domain"
)

type (
	WeatherController struct {
		responder       *web.Responder
		service         domain.Service
		identityService *auth.WebIdentityService
	}

	viewData struct {
		City    string
		Weather domain.Weather
	}
)

// Inject dependencies
func (w *WeatherController) Inject(
	responder *web.Responder,
	service domain.Service,
	identityService *auth.WebIdentityService,
) *WeatherController {
	w.responder = responder
	w.service = service
	w.identityService = identityService

	return w
}

func (w *WeatherController) CurrentCityWeather(ctx context.Context, r *web.Request) web.Result {
	city := r.Params["city"]
	weather, err := w.service.GetByCity(ctx, city)
	if err != nil {
		return w.responder.ServerError(err)
	}

	return w.responder.Render("weather/weather", viewData{
		City:    city,
		Weather: weather,
	})
}

func (w *WeatherController) CurrentCityWeatherSpecial(ctx context.Context, r *web.Request) web.Result {
	identity := w.identityService.Identify(ctx, r)

	city := r.Params["city"]
	weather, err := w.service.GetByCity(ctx, city)
	if err != nil {
		return w.responder.ServerError(err)
	}

	return w.responder.Data(&struct {
		User    string
		Weather domain.Weather
	}{
		User:    identity.Subject(),
		Weather: weather,
	})
}

// HandleIfLoggedIn allows a controller to be used for logged-in users
func (w *WeatherController) HandleIfLoggedIn(action web.Action) web.Action {
	return func(ctx context.Context, req *web.Request) web.Result {
		if identity := w.identityService.Identify(ctx, req); identity != nil {
			return action(ctx, req)
		}
		redirectURL := req.Request().URL.String()
		return w.responder.RouteRedirect(
			"core.auth.login",
			map[string]string{
				"broker":      "keycloak",
				"redirecturl": redirectURL,
			})
	}
}

func (w *WeatherController) Data(ctx context.Context, req *web.Request, callParams web.RequestParams) interface{} {
	city, ok := callParams["city"]
	if !ok {
		return domain.Weather{}
	}

	if identity := w.identityService.Identify(ctx, req); identity != nil {
		c := &struct {
			Address struct {
				Locality string
			}
		}{}
		err := identity.(oauth.OpenIDIdentity).IDTokenClaims(&c)
		if err == nil {
			city = c.Address.Locality
		}
	}

	weather, err := w.service.GetByCity(ctx, city)
	if err != nil {
		return domain.Weather{}
	}

	return weather
}
