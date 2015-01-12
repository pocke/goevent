package goevent

import "fmt"

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

func (e *EventNotDefined) EventName() string {
	return e.eventName
}
