package goevent

import (
	"fmt"
	"reflect"
	"sync"
)

type Event struct {
	// listeners are listener functions.
	listeners []reflect.Value
	lmu       sync.RWMutex

	argTypes []reflect.Type
	tmu      sync.RWMutex
}

func New() *Event {
	return &Event{}
}

func (p *Event) Trigger(args ...interface{}) error {
	p.lmu.Lock()
	defer p.lmu.Unlock()

	arguments := make([]reflect.Value, 0, len(args))
	argTypes := make([]reflect.Type, 0, len(args))
	for _, v := range args {
		arguments = append(arguments, reflect.ValueOf(v))
		argTypes = append(argTypes, reflect.TypeOf(v))
	}

	err := p.validateArgs(argTypes)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(p.listeners))
	for _, fn := range p.listeners {
		go func(f reflect.Value) {
			defer wg.Done()
			f.Call(arguments)
		}(fn)
	}

	wg.Wait()
	return nil
}

func (p *Event) On(f interface{}) error {
	fn, err := p.checkFuncSignature(f)
	if err != nil {
		return err
	}

	p.lmu.Lock()
	defer p.lmu.Unlock()
	p.listeners = append(p.listeners, *fn)

	return nil
}

func (p *Event) Off() {
	panic(fmt.Errorf("Off() has not been implemented yet."))
}

func (p *Event) checkFuncSignature(f interface{}) (*reflect.Value, error) {
	fn := reflect.ValueOf(f)
	if fn.Kind() != reflect.Func {
		return nil, fmt.Errorf("Argument should be a function")
	}

	types := fnArgTypes(fn)

	p.lmu.RLock()
	defer p.lmu.RUnlock()
	if len(p.listeners) == 0 {
		p.tmu.Lock()
		defer p.tmu.Unlock()
		p.argTypes = types
		return &fn, nil
	}

	err := p.validateArgs(types)
	if err != nil {
		return nil, err
	}

	return &fn, nil
}

func (p *Event) validateArgs(types []reflect.Type) error {
	p.tmu.RLock()
	defer p.tmu.RUnlock()
	if len(types) != len(p.argTypes) {
		return fmt.Errorf("Argument length expected %d, but got %d", len(p.argTypes), len(types))
	}
	for i, t := range types {
		if t != p.argTypes[i] {
			return fmt.Errorf("Argument Error. Args[%d] expected %s, but got %s", i, p.argTypes[i], t)
		}
	}

	return nil
}

func fnArgTypes(fn reflect.Value) []reflect.Type {
	fnType := fn.Type()
	fnNum := fnType.NumIn()

	types := make([]reflect.Type, 0, fnNum)

	for i := 0; i < fnNum; i++ {
		types = append(types, fnType.In(i))
	}

	return types
}
