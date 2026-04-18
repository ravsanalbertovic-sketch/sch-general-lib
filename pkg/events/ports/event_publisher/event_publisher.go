package event_publisher

import (
	"context"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/ports/event"
)

type EventPublisher interface {
	Publish(ctx context.Context, event event.Event) error
}
