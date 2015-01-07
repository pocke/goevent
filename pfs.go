package pfs

import (
	"fmt"
	"reflect"
	"sync"
)

type PFS struct {
	// funcs are listener functions.
	funcs []reflect.Value
	mu    *sync.RWMutex
}

func New() *PFS {
	return &PFS{
		funcs: make([]reflect.Value, 0),
		mu:    &sync.RWMutex{},
	}
}

func (p *PFS) Pub(args ...interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()

	arguments := make([]reflect.Value, 0, len(args))
	for _, v := range args {
		arguments = append(arguments, reflect.ValueOf(v))
	}

	wg := sync.WaitGroup{}

	wg.Add(len(p.funcs))
	for _, fn := range p.funcs {
		go func(f reflect.Value) {
			defer wg.Done()
			f.Call(arguments)
		}(fn)
	}

	wg.Wait()
}

func (p *PFS) Sub(f interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	fn := reflect.ValueOf(f)

	if reflect.Func != fn.Kind() {
		return fmt.Errorf("Argument should be a function")
	}

	if len(p.funcs) != 0 {
		// TODO: check fn arguments
	}

	p.funcs = append(p.funcs, fn)
	return nil
}

func (p *PFS) Off() {

}

func (p *PFS) Filter(func() bool) {

}
