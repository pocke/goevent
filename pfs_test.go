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

	ok := p.Pub(2)
	if i != 3 {
		t.Errorf("Expected i == 3, Got i == %d", i)
	}
	if !ok {
		t.Error("should return true When not reject. But got false.")
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
	pfs := pfs.New()
	err := pfs.Sub("foobar")
	if err == nil {
		t.Error("should return error When recieve not function. But got nil.")
	}
}

func TestFilter(t *testing.T) {
	p := pfs.New()

	i := 1
	err := p.Sub(func(j int) {
		i += j
	})
	if err != nil {
		t.Fatal(err)
	}

	err = p.Filter(func(j int) bool {
		return j != 3
	})
	if err != nil {
		t.Fatal(err)
	}

	ok := p.Pub(2)
	if i != 3 {
		t.Errorf("Expected i == 3, Got i == %d", i)
	}
	if !ok {
		t.Error("should return true When not reject. But got false.")
	}

	ok = p.Pub(3)
	if i != 3 {
		t.Errorf("Expected i == 3, Got i == %d", i)
	}
	if ok {
		t.Error("should return false When reject. But got true.")
	}
}

func TestFilterWhenNotFunction(t *testing.T) {
	pfs := pfs.New()
	err := pfs.Filter("foobar")
	if err == nil {
		t.Error("should return error When recieve not function. But got nil.")
	}
}
