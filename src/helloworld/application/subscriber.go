package application

import (
	"context"

	"flamingo.me/flamingo/v3/framework/flamingo"
)

type (
	Subscriber struct {
		logger flamingo.Logger
	}
)

// Inject dependencies
func (s *Subscriber) Inject(
	logger flamingo.Logger,
) *Subscriber {
	s.logger = logger

	return s
}

func (s *Subscriber) Notify(ctx context.Context, e flamingo.Event) {
	if event, ok := e.(*IncomingRequestEvent); ok {
		s.logger.WithContext(ctx).Info("Incoming request: ", event.Request)
	}
}
