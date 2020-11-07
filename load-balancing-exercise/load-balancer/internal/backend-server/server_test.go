package server

import (
	"testing"

	"github.com/agronskiy/golang-learning/load-balancing-exercise/consts"
)

func TestInit(t *testing.T) {

	if len(freePorts) != consts.MaxListeners || cap(freePorts) != len(freePorts) {
		t.Errorf("Wrong length of `freePorts`.")
	}

	if freePorts[consts.MaxListeners-1] != "50010" {
		t.Errorf("freePorts initialization failed")
	}
}

func TestGetNextFreePort(t *testing.T) {

	for i := 0; i < consts.MaxListeners; i++ {

		if _, err := getNextFreePort(); err != nil {
			t.Errorf("getNextFreePort() returned error when no needed")
		}
	}

	if _, err := getNextFreePort(); err == nil {
		t.Errorf("getNextFreePort() returned no error when needed")
	}
}
