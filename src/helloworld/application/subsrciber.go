package application

import (
	"context"

	"flamingo.me/flamingo/v3/framework/flamingo"
)

type (
	Subscriber struct {
	}
)

func (s *Subscriber) Notify(ctx context.Context, event flamingo.Event) {
	// if e, ok := event.(*IncomingRequestEvent); ok {
	// 	fmt.Println(e.Request)
	// }
}
