package pfs

import (
	"fmt"
	"reflect"
	"sync"
)

type PFS struct {
	// listeners are listener functions.
	listeners []reflect.Value
	lmu       *sync.RWMutex

	filters []reflect.Value
	fmu     *sync.RWMutex
}

func New() *PFS {
	return &PFS{
		listeners: make([]reflect.Value, 0),
		lmu:       &sync.RWMutex{},
		filters:   make([]reflect.Value, 0),
		fmu:       &sync.RWMutex{},
	}
}

func (p *PFS) Pub(args ...interface{}) {
	p.lmu.Lock()
	defer p.lmu.Unlock()

	arguments := make([]reflect.Value, 0, len(args))
	for _, v := range args {
		arguments = append(arguments, reflect.ValueOf(v))
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
}

func (p *PFS) Sub(f interface{}) error {
	p.lmu.Lock()
	defer p.lmu.Unlock()

	fn := reflect.ValueOf(f)

	if reflect.Func != fn.Kind() {
		return fmt.Errorf("Argument should be a function")
	}

	if len(p.listeners) != 0 {
		// TODO: check fn arguments
	}

	p.fmu.RLock()
	if len(p.filters) != 0 {
		// TODO: check fn arguments
	}
	p.fmu.RUnlock()

	p.listeners = append(p.listeners, fn)
	return nil
}

func (p *PFS) Off() {

}

func (p *PFS) Filter(f interface{}) error {
	p.fmu.Lock()
	defer p.fmu.Unlock()

	// TODO: DRY
	fn := reflect.ValueOf(f)

	if reflect.Func != fn.Kind() {
		return fmt.Errorf("Argument should be a function")
	}

	p.lmu.RLock()
	if len(p.listeners) != 0 {
		// TODO: check fn arguments
	}
	p.lmu.RUnlock()

	if len(p.filters) != 0 {
		// TODO: check fn arguments
	}

	p.filters = append(p.filters, fn)

	return nil
}
