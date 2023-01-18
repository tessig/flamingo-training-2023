package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/training/src/helloworld/application"
	"flamingo.me/training/src/helloworld/interfaces/controller"
)

type (
	Module struct{}
	routes struct {
		helloController *controller.Hello
	}
)

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
	injector.BindMulti(new(web.Filter)).To(new(application.Filter))
	flamingo.BindEventSubscriber(injector).To(new(application.Subscriber))
}

// Inject dependencies
func (r *routes) Inject(
	helloController *controller.Hello,
) *routes {
	r.helloController = helloController

	return r
}

// CueConfig for the module
func (m *Module) CueConfig() string {
	// language=cue
	return `
helloworld: {
	greeting: string | *"Hello Flamingo"
}
`
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/", "hello")
	registry.HandleGet("hello", r.helloController.Hello)

	registry.MustRoute("/hello/:name", "hello.name")
	registry.HandleGet("hello.name", r.helloController.HelloName)
}
