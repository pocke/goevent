package pfs

import (
	"fmt"
	"reflect"
	"sync"
)

type PSF struct {
	mu *sync.RWMutex
	// events are listener functions.
	events []reflect.Value
}

func (p *PSF) Pub() {

}

func (p *PSF) Sub(f interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	fn := reflect.ValueOf(f)

	if reflect.Func != fn.Kind() {
		return fmt.Errorf("Argument should be a function")
	}

	if len(p.events) != 0 {
		// TODO: check fn arguments
	}

	p.events = append(p.events, fn)
	return nil
}

func (p *PSF) Off() {

}

func (p *PSF) Filter(func() bool) {

}
