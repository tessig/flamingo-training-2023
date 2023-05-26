package controllers

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	HelloController struct {
		greeting  string
		responder *web.Responder
	}

	viewData struct {
		Greeting string
		Name     string
	}
)

// Inject dependencies
func (h *HelloController) Inject(
	responder *web.Responder,
	cfg *struct {
		Greeting string `inject:"config:helloworld.greeting"`
	},
) *HelloController {
	h.responder = responder
	if cfg != nil {
		h.greeting = cfg.Greeting
	}

	return h
}

func (h *HelloController) HelloAction(ctx context.Context, req *web.Request) web.Result {
	return h.responder.Render("index", viewData{
		Greeting: h.greeting,
	})
}

func (h *HelloController) HelloNameAction(ctx context.Context, req *web.Request) web.Result {
	name, ok := req.Params["name"]
	if !ok {
		name = "Unknown Person"
	}
	return h.responder.Render("index", viewData{
		Greeting: h.greeting,
		Name:     name,
	})
}
