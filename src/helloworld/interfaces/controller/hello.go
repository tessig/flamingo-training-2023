package controller

import (
	"context"
	"fmt"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	Hello struct {
		responder *web.Responder
		greeting  string
	}
)

// Inject dependencies
func (h *Hello) Inject(
	responder *web.Responder,
	cfg *struct {
		Greeting string `inject:"config:helloworld.greeting"`
	},
) *Hello {
	h.responder = responder
	if cfg != nil {
		h.greeting = cfg.Greeting
	}

	return h
}

func (h *Hello) Hello(ctx context.Context, r *web.Request) web.Result {
	return h.responder.Render("index", h.greeting)
}

func (h *Hello) HelloName(ctx context.Context, r *web.Request) web.Result {
	name := r.Params["name"]
	return h.responder.Render("index", fmt.Sprintf("Hello %s!", name))
}
