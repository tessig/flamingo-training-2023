package controllers

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	HelloController struct {
		responder *web.Responder
	}
)

// Inject dependencies
func (h *HelloController) Inject(
	responder *web.Responder,
) *HelloController {
	h.responder = responder

	return h
}

func (h *HelloController) HelloAction(ctx context.Context, req *web.Request) web.Result {
	return h.responder.Render("index", []string{"hello", "world"})
}
