package event

import "github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/vo/id"

type Event interface {
	GetID() id.ID
	Name() string
	AggregateType() string
	AggregateID() string
}
