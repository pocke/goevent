package goevent

import "sync"

// Table is an event table.
type Table interface {
	Trigger(name string, args ...interface{}) error
	// f is a function
	On(name string, f interface{}) error
	Off(name string, f interface{}) error
	// Destroy a event
	Destroy(name string) error
}

type table struct {
	events map[string]Event
	mu     sync.RWMutex
}

// NewTable creates a new event table.
func NewTable() Table {
	return &table{
		events: map[string]Event{},
	}
}

func (t *table) Trigger(name string, args ...interface{}) error {
	t.mu.RLock()
	defer t.mu.RUnlock()

	ev, ok := t.events[name]
	if !ok {
		return newEventNotDefined(name)
	}

	return ev.Trigger(args...)
}

func (t *table) On(name string, f interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	ev, ok := t.events[name]
	if !ok {
		ev = New()
		t.events[name] = ev
	}
	return ev.On(f)
}

func (t *table) Off(name string, f interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	e, ok := t.events[name]
	if !ok {
		return newEventNotDefined(name)
	}

	return e.Off(f)
}

func (t *table) Destroy(name string) error {
	if _, ok := t.events[name]; !ok {
		return newEventNotDefined(name)
	}
	delete(t.events, name)
	return nil
}

var _ Table = &table{}
