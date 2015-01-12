package goevent_test

import (
	"sync"
	"testing"

	"github.com/pocke/goevent"
)

func TestEventNew(t *testing.T) {
	p := goevent.New()
	t.Log("Event: %+v", p)
}

func TestOnTrigger(t *testing.T) {
	p := goevent.New()

	i := 1
	err := p.On(func(j int) {
		i += j
	})
	if err != nil {
		t.Fatal(err)
	}

	err = p.Trigger(2)
	if i != 3 {
		t.Errorf("Expected i == 3, Got i == %d", i)
	}
	if err != nil {
		t.Error("should not return error When not reject. But got %s.", err)
	}

	err = p.Trigger("2")
	if err == nil {
		t.Error("should return error when invalid type. But got nil")
	}
}

func TestManyTrigger(t *testing.T) {
	p := goevent.New()
	i := 0
	p.On(func(j int) {
		i += j
	})

	for j := 0; j < 1000; j++ {
		p.Trigger(1)
	}

	if i != 1000 {
		t.Errorf("i should be 1000, but got %d", i)
	}
}

func TestManyOn(t *testing.T) {
	p := goevent.New()
	i := 0
	m := sync.Mutex{}
	for j := 0; j < 1000; j++ {
		p.On(func(j int) {
			m.Lock()
			defer m.Unlock()
			i += j
		})
	}
	p.Trigger(1)
	if i != 1000 {
		t.Errorf("i should be 1000, but got %d", i)
	}
}

func TestOnWhenNotFunction(t *testing.T) {
	p := goevent.New()
	err := p.On("foobar")
	if err == nil {
		t.Error("should return error When recieve not function. But got nil.")
	}
}

func TestOnWhenInvalidArgs(t *testing.T) {
	p := goevent.New()
	p.On(func(i int) {})

	err := p.On(func() {})
	if err == nil {
		t.Error("Should return error when different argument num. But got nil")
	}

	err = p.On(func(s string) {})
	if err == nil {
		t.Error("Should return error when different args type. But got nil")
	}
}

func TestOff(t *testing.T) {
	p := goevent.New()
	i := 0
	j := 0
	k := 0

	p.On(func() { j++ })
	f := func() { i++ }
	p.On(f)
	p.On(func() { k++ })

	p.Trigger()
	if i != 1 {
		t.Errorf("i expected 1, but got %d", i)
	}
	if j != 1 {
		t.Errorf("j expected 1, but got %d", j)
	}
	if k != 1 {
		t.Errorf("k expected 1, but got %d", k)
	}

	err := p.Off(f)
	if err != nil {
		t.Error(err)
	}
	p.Trigger()
	if i != 1 {
		t.Errorf("i expected 1, but got %d", i)
	}
	if j != 2 {
		t.Errorf("j expected 2, but got %d", j)
	}
	if k != 2 {
		t.Errorf("k expected 2, but got %d", k)
	}

	err = p.Off(f)
	if err == nil {
		t.Errorf("should return error when Listener doesn't exists. but got nil")
	}
}
