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
	return h.responder.Render("index", h.greeting)
}
