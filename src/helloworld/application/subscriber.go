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

func (s *Subscriber) Notify(ctx context.Context, event flamingo.Event) {
	if e, ok := event.(*IncomingRequestEvent); ok {
		s.logger.WithContext(ctx).Info("incoming request: ", e.Request)
	}
}
