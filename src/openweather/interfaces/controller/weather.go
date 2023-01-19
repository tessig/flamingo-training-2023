package controller

import (
	"context"

	"flamingo.me/flamingo/v3/core/auth"
	"flamingo.me/flamingo/v3/core/auth/oauth"
	"flamingo.me/flamingo/v3/framework/web"
	"go.opencensus.io/trace"

	"flamingo.me/training/src/openweather/application"
	"flamingo.me/training/src/openweather/domain"
)

type (
	Weather struct {
		responder       *web.Responder
		weatherService  application.Service
		identityService *auth.WebIdentityService
	}

	viewData struct {
		Weather domain.Weather
	}
)

// Inject dependencies
func (c *Weather) Inject(
	responder *web.Responder,
	weatherService application.Service,
	identityService *auth.WebIdentityService,
) *Weather {
	c.responder = responder
	c.weatherService = weatherService
	c.identityService = identityService

	return c
}

func (c *Weather) Weather(ctx context.Context, r *web.Request) web.Result {
	ctx, span := trace.StartSpan(ctx, "weather/city")
	defer span.End()
	city := r.Params["city"]

	weather, err := c.weatherService.GetByCity(ctx, city)
	if err != nil {
		c.responder.ServerError(err)
	}

	return c.responder.Render("weather/weather", &viewData{
		Weather: weather,
	})
}

func (c *Weather) Detail(ctx context.Context, req *web.Request, callParams web.RequestParams) interface{} {
	identity := c.identityService.Identify(ctx, req)

	city := callParams["city"]
	if identity != nil {
		oidcIdentity, ok := identity.(oauth.OpenIDIdentity)
		if ok {
			var claims struct {
				Address struct {
					Locality string
				}
			}
			err := oidcIdentity.IDTokenClaims(&claims)
			if err == nil {
				city = claims.Address.Locality
			}
		}
	}

	weather, err := c.weatherService.GetByCity(ctx, city)
	if err != nil {
		return nil
	}

	return weather
}
