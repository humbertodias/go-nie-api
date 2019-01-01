package fake

import "testing"

func TestNif(t *testing.T) {
	nif := NifRandom()
	if !NifValid(nif) {
		t.Errorf("Nif invalid, got: %s", nif)
	}
	if !NifCheckDigit(nif, 0) {
		t.Errorf("Nif invalid digit, got: %s", nif)
	}
}
