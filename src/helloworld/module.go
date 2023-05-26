package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/training/src/helloworld/application"
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

	injector.BindMulti(new(web.Filter)).To(new(application.Filter))
	flamingo.BindEventSubscriber(injector).To(new(application.Subscriber))
}

// CueConfig for the module
func (m *Module) CueConfig() string {
	// language=cue
	return `
helloworld: {
	greeting: string | *"hello from default config"
}
`
}
func (r *routes) Inject(helloController *controllers.HelloController) *routes {
	r.helloController = helloController

	return r
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/", "helloworld")
	registry.HandleAny("helloworld", r.helloController.HelloAction)
	registry.MustRoute("/hello/:name", "hello.name")
	registry.HandleAny("hello.name", r.helloController.HelloNameAction)
}
