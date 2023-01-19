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
func (c *Hello) Inject(
	responder *web.Responder,
	cfg *struct {
		Greeting string `inject:"config:helloworld.greeting"`
	},
) *Hello {
	c.responder = responder
	if cfg != nil {
		c.greeting = cfg.Greeting
	}

	return c
}

func (c *Hello) HelloAction(ctx context.Context, r *web.Request) web.Result {
	return c.responder.Render("index", c.greeting)
}

func (c *Hello) HelloName(ctx context.Context, r *web.Request) web.Result {
	name := r.Params["name"]
	return c.responder.Render("index", fmt.Sprintf("Hello %s!", name))
}
