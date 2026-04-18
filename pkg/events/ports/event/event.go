package event

type Event interface {
	Name() string
	AggregateType() string
	AggregateID() string
}
