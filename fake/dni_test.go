package fake

import "testing"

func TestDni(t *testing.T) {
	dni := DniRandom()
	if !DniValid(dni) {
		t.Errorf("Dni invalid, got: %s", dni)
	}
	if !DniCheckDigit(dni) {
		t.Errorf("Dni invalid digit, got: %s", dni)
	}

}
