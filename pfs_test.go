package pfs_test

import (
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

	p.Pub(2)

	if i != 3 {
		t.Errorf("Expected i == 3, Got i == %d", i)
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
		return j == 3
	})
	if err != nil {
		t.Fatal(err)
	}

	p.Pub(2)
	if i != 3 {
		t.Errorf("Expected i == 3, Got i == %d", i)
	}

	p.Pub(3)
	if i != 3 {
		t.Errorf("Expected i == 3, Got i == %d", i)
	}
}
