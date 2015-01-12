package goevent_test

import (
	"testing"

	"github.com/pocke/goevent"
)

func TestNewTable(t *testing.T) {
	ta := goevent.NewTable()
	t.Logf("%#v", ta)
}

func TestTableOnTrigger(t *testing.T) {
	ta := goevent.NewTable()
	i := 0
	err := ta.On("foo", func(j int) { i += j })
	if err != nil {
		t.Error(err)
	}

	err = ta.Trigger("foo", 1)
	if err != nil {
		t.Error(err)
	}
	if i != 1 {
		t.Errorf("i expected 1, but got %d", i)
	}
}

func TestTableTriggerFail(t *testing.T) {
	ta := goevent.NewTable()
	err := ta.Trigger("foo", 1)
	if err == nil {
		t.Error("should return error when event has not been defined yet. But got nil")
	}
}

func TestTableOff(t *testing.T) {
	ta := goevent.NewTable()
	ta.Off("foo", func() {})

	i := 0
	f := func() { i++ }
	ta.On("foo", f)
	ta.Trigger("foo")
	err := ta.Off("foo", f)
	if err != nil {
		t.Error(err)
	}
	ta.Trigger("foo")
	if i != 1 {
		t.Errorf("i expected 1, but got %d", i)
	}
}

func TestTableDestroy(t *testing.T) {
	ta := goevent.NewTable()
	i := 0
	ta.On("foo", func(j int) { i += j })

	ta.Trigger("foo", 1)
	err := ta.Destroy("foo")
	if err != nil {
		t.Error(err)
	}
	err = ta.Trigger("foo", 1)
	if err == nil {
		t.Errorf("should destroy event. but not destroy.")
	}
	if i != 1 {
		t.Errorf("i expected 1, but got %d", i)
	}

	err = ta.Destroy("foo")
	if err == nil {
		t.Errorf("should return error when event has not been defined yet. but got nil")
	}
}
