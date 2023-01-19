package openweather

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/core/healthcheck/domain/healthcheck"
	"flamingo.me/flamingo/v3/framework/web"
	flamingoGQL "flamingo.me/graphql"

	"flamingo.me/training/src/openweather/application"
	"flamingo.me/training/src/openweather/infrastructure"
	"flamingo.me/training/src/openweather/interfaces/controller"
	"flamingo.me/training/src/openweather/interfaces/graphql"
)

type (
	Module struct {
		useFake bool
	}

	routes struct {
		weatherController *controller.Weather
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
		injector.Override(new(application.Service), "").To(new(infrastructure.Fakeservice))
	}

	injector.BindMap(new(healthcheck.Status), "openweather").To(new(infrastructure.Adapter))

	injector.Bind(new(cache.HTTPFrontend)).AnnotatedWith("openweather").In(dingo.Singleton)
	injector.Bind(new(cache.Backend)).ToInstance(cache.NewInMemoryCache())

	injector.BindMulti(new(flamingoGQL.Service)).To(new(graphql.Service))
}

// CueConfig for the module
func (m *Module) CueConfig() string {
	return `
openweather: {
	apiURL: string
	apiKey: string | *flamingo.os.env.OPENWEATHER_API_KEY
	useFake: bool | *false
}
`
}

// Inject dependencies
func (r *routes) Inject(
	weatherController *controller.Weather,
) *routes {
	r.weatherController = weatherController

	return r
}

// Routes definition for the module
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/weather/:city", "openweather.city")
	registry.HandleAny("openweather.city", r.weatherController.Weather)
	registry.HandleData("openweather.detail", r.weatherController.Detail)
}
