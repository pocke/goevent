package goevent

import "fmt"

// EventNotDefined is an error indicationg that the event has not been defined.
type EventNotDefined struct {
	eventName string
}

func newEventNotDefined(name string) *EventNotDefined {
	return &EventNotDefined{
		eventName: name,
	}
}

func (e *EventNotDefined) Error() string {
	return fmt.Sprintf("%s event has not been defined yet.", e.eventName)
}

// EventName return name of the event.
func (e *EventNotDefined) EventName() string {
	return e.eventName
}

var _ error = newEventNotDefined("foo")
