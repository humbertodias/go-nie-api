package handler

import "testing"

func TestNif(t *testing.T) {
	nif := NifRandom()
	if !NifValid(nif) {
		t.Errorf("Nif invalid, got: %s", nif)
	}
}
