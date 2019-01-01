package fake

import "testing"

func TestNif(t *testing.T) {
	Seed(1234)
	nif := NifRandom()
	if !NifValid(nif) || nif != "75231682Q" {
		t.Errorf("Nif invalid, got: %s", nif)
	}
}
