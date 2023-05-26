package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/training/src/helloworld/interfaces/controllers"
)

type (
	Module struct {
	}

	routes struct {
		helloController *controllers.HelloController
	}
)

func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

func (r *routes) Inject(helloController *controllers.HelloController) *routes {
	r.helloController = helloController

	return r
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/", "helloworld")
	registry.HandleAny("helloworld", r.helloController.HelloAction)
}
