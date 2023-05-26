package openweather

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/core/healthcheck/domain/healthcheck"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/training/src/openweather/application"
	"flamingo.me/training/src/openweather/infrastructure"
	"flamingo.me/training/src/openweather/interfaces/controllers"
)

type (
	Module struct {
		useFake bool
	}

	routes struct {
		weatherController *controllers.WeatherController
	}
)

// Inject dependencies
func (m *Module) Inject(
	cfg *struct {
		UseFake bool `inject:"config:openweather.useFake"`
	},
) *Module {
	if cfg != nil {
		m.useFake = cfg.UseFake
	}

	return m
}

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))

	injector.Bind(new(application.Service)).To(new(infrastructure.Adapter))
	if m.useFake {
		injector.Override(new(application.Service), "").To(new(infrastructure.FakeWeather))
	}

	injector.BindMap(new(healthcheck.Status), "openweather").To(infrastructure.Adapter{})

	injector.Bind(new(cache.HTTPFrontend)).AnnotatedWith("openweather").In(dingo.Singleton)
	injector.Bind(new(cache.Backend)).ToInstance(cache.NewInMemoryCache())
}

// CueConfig for the module
func (m *Module) CueConfig() string {
	// language=cue
	return `
flamingo: {
	os: {
		env: {
			OPENWEATHER_API_KEY: string | * ""
		}
	}
}

openweather: {
	useFake: bool | *false
	apiURL: string | *"http://api.openweathermap.org/data/2.5/"
	apiKey: flamingo.os.env.OPENWEATHER_API_KEY
}
`
}

// Inject dependencies
func (r *routes) Inject(
	weatherController *controllers.WeatherController,
) *routes {
	r.weatherController = weatherController

	return r
}

// Routes definition for the module
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/weather/:city", "weather.city")
	registry.HandleGet("weather.city", r.weatherController.Weather)

}
