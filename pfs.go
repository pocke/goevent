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
}

func New() *PFS {
	return &PFS{
		listeners: make([]reflect.Value, 0),
		lmu:       &sync.RWMutex{},
	}
}

func (p *PFS) Pub(args ...interface{}) bool {
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
	return true
}

func (p *PFS) Sub(f interface{}) error {
	fn, err := p.checkFuncSignature(f)
	if err != nil {
		return err
	}

	p.lmu.Lock()
	defer p.lmu.Unlock()
	p.listeners = append(p.listeners, *fn)

	return nil
}

func (p *PFS) Off() {
	panic(fmt.Errorf("Off() has not been implemented yet."))
}

func (p *PFS) checkFuncSignature(f interface{}) (*reflect.Value, error) {
	fn := reflect.ValueOf(f)
	if fn.Kind() != reflect.Func {
		return nil, fmt.Errorf("Argument should be a function")
	}

	p.lmu.RLock()
	defer p.lmu.RUnlock()
	if len(p.listeners) != 0 {
		// TODO: check fn arguments
	}

	return &fn, nil
}
