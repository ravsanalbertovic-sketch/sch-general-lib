package events

import "context"

type EventPublisher interface {
	Publish(ctx context.Context, event Event) error
}
