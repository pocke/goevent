package pfs_test

import (
	"testing"

	"github.com/pocke/pfs"
)

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
