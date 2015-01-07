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
	pfs := pfs.New()

	i := 1
	err := pfs.Sub(func(j int) {
		i += j
	})
	if err != nil {
		t.Error(err)
	}

	pfs.Pub(2)

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
