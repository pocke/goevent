package goevent

import (
	"fmt"
	"sync"
)

type Table interface {
	On(string, interface{}) error
	Trigger(string, ...interface{}) error
}

type table struct {
	events map[string]*Event
	mu     sync.RWMutex
}

func NewTable() Table {
	return &table{
		events: map[string]*Event{},
	}
}

func (t *table) Trigger(name string, args ...interface{}) error {
	t.mu.RLock()
	defer t.mu.RUnlock()

	ev, ok := t.events[name]
	if !ok {
		return fmt.Errorf("%s event has not been defined yet.", name)
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

var _ Table = &table{}
