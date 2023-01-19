package application

import (
	"context"
	"net/http"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	Filter struct {
		eventRouter flamingo.EventRouter
	}

	IncomingRequestEvent struct {
		Request *http.Request
	}
)

// Inject dependencies
func (f *Filter) Inject(
	eventRouter flamingo.EventRouter,
) *Filter {
	f.eventRouter = eventRouter

	return f
}

func (f *Filter) Filter(ctx context.Context, r *web.Request, w http.ResponseWriter, chain *web.FilterChain) web.Result {
	f.eventRouter.Dispatch(ctx, &IncomingRequestEvent{Request: r.Request()})

	return chain.Next(ctx, r, w)
}

