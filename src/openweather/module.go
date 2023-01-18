package openweather

import (
	"fmt"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/core/healthcheck/domain/healthcheck"
	"flamingo.me/flamingo/v3/framework/web"
	flamingoGraphQL "flamingo.me/graphql"
	"github.com/spf13/cobra"

	"flamingo.me/training/src/openweather/domain"
	infrastructure2 "flamingo.me/training/src/openweather/infrastructure"
	"flamingo.me/training/src/openweather/interfaces/controller"
	"flamingo.me/training/src/openweather/interfaces/graphql"
)

type (
	Module struct {
		useFake bool
	}
	routes struct {
		controller *controller.WeatherController
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
	injector.BindMulti(new(cobra.Command)).ToProvider(DefaultCityCmd)
	injector.Bind(new(domain.Service)).To(new(infrastructure2.Adapter))
	injector.BindMap(new(healthcheck.Status), "openweather").To(new(infrastructure2.Adapter))
	if m.useFake {
		injector.Override(new(domain.Service), "").To(new(infrastructure2.Fakeservice))
	}

	injector.Bind(new(cache.HTTPFrontend)).AnnotatedWith("openweather").In(dingo.Singleton)

	injector.BindMulti(new(flamingoGraphQL.Service)).To(new(graphql.Service))
}

// CueConfig for the module
func (m *Module) CueConfig() string {
	// language=cue
	return `
openweather: {
  defaultCity: string | *"Wiesbaden"
  useFake: bool | *false 
  apiURL: string
  apiKey: string
}
`
}

// Inject dependencies
func (r *routes) Inject(
	controller *controller.WeatherController,
) *routes {
	r.controller = controller

	return r
}

// Routes definition for the module
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/weather/:city", "weather")
	registry.HandleGet("weather", r.controller.CurrentCityWeather)
	registry.MustRoute("/weather/special/:city", "weather-special")
	registry.HandleGet("weather-special", r.controller.HandleIfLoggedIn(r.controller.CurrentCityWeatherSpecial))
	registry.HandleData("openweather.detail", r.controller.Data)
}

func DefaultCityCmd(cfg *struct {
	DefaultCity string `inject:"config:openweather.defaultCity"`
}) *cobra.Command {
	return &cobra.Command{
		Use:   "defaultCity",
		Short: "The openweather default city",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cfg.DefaultCity)
		},
	}
}
