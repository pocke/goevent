package pfs_test

import (
	"sync"
	"testing"

	"github.com/pocke/pfs"
)

func TestPFSNew(t *testing.T) {
	p := pfs.New()
	t.Log("PFS: %+v", p)
}

func TestPubSub(t *testing.T) {
	p := pfs.New()

	i := 1
	err := p.Sub(func(j int) {
		i += j
	})
	if err != nil {
		t.Fatal(err)
	}

	err = p.Pub(2)
	if i != 3 {
		t.Errorf("Expected i == 3, Got i == %d", i)
	}
	if err != nil {
		t.Error("should not return error When not reject. But got %s.", err)
	}

	err = p.Pub("2")
	if err == nil {
		t.Error("should return error when invalid type. But got nil")
	}
}

func TestManyPub(t *testing.T) {
	p := pfs.New()
	i := 0
	p.Sub(func(j int) {
		i += j
	})

	for j := 0; j < 1000; j++ {
		p.Pub(1)
	}

	if i != 1000 {
		t.Errorf("i should be 1000, but got %d", i)
	}
}

func TestManySub(t *testing.T) {
	p := pfs.New()
	i := 0
	m := sync.Mutex{}
	for j := 0; j < 1000; j++ {
		p.Sub(func(j int) {
			m.Lock()
			defer m.Unlock()
			i += j
		})
	}
	p.Pub(1)
	if i != 1000 {
		t.Errorf("i should be 1000, but got %d", i)
	}
}

func TestSubWhenNotFunction(t *testing.T) {
	p := pfs.New()
	err := p.Sub("foobar")
	if err == nil {
		t.Error("should return error When recieve not function. But got nil.")
	}
}

func TestSubWhenInvalidArgs(t *testing.T) {
	p := pfs.New()
	p.Sub(func(i int) {})

	err := p.Sub(func() {})
	if err == nil {
		t.Error("Should return error when different argument num. But got nil")
	}

	err = p.Sub(func(s string) {})
	if err == nil {
		t.Error("Should return error when different args type. But got nil")
	}
}
