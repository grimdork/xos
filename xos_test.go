package xos_test

import (
	"testing"

	"github.com/grimdork/xos"
)

func TestAppPaths(t *testing.T) {
	t.Log("Testing CLI/desktop paths.")
	cp, err := xos.NewConfig("test")
	if err != nil {
		t.Errorf("Error building path: %s", err.Error())
		t.FailNow()
	}

	t.Logf("Resulting path: %s", cp.Path())
}

func TestServerPaths(t *testing.T) {
	t.Log("Testing server paths.")
	cp, err := xos.NewServerConfig("test")
	if err != nil {
		t.Errorf("Error building path: %s", err.Error())
		t.FailNow()
	}

	t.Logf("Resulting path: %s", cp.Path())
}
