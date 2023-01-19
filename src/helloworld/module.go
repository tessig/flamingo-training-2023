package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/training/src/helloworld/application"
	"flamingo.me/training/src/helloworld/interfaces/controller"
)

type (
	Module struct {
	}

	routes struct {
		helloController *controller.Hello
	}
)

// Inject dependencies
func (r *routes) Inject(
	helloController *controller.Hello,
) *routes {
	r.helloController = helloController

	return r
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/", "helloworld.hello")
	registry.HandleAny("helloworld.hello", r.helloController.HelloAction)
	registry.MustRoute("/hello/:name", "helloworld.name")
	registry.HandleAny(`helloworld.name`, r.helloController.HelloName)
}

func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))

	injector.BindMulti(new(web.Filter)).To(new(application.Filter))
	flamingo.BindEventSubscriber(injector).To(new(application.Subscriber))
}

// CueConfig for the module
func (m *Module) CueConfig() string {
	//language=cue
	return `
{
	helloworld: {
		greeting: string | *"Hello Word"
	}
}
`
}
