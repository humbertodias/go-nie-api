package fake

import "testing"

func TestDni(t *testing.T) {
	Seed(1234)
	dni := DniRandom()
	if !DniValid(dni) || dni != "75231682Q" {
		t.Errorf("Dni invalid, got: %s", dni)
	}
	if !DniCheckDigit(dni) {
		t.Errorf("Dni invalid digit, got: %s", dni)
	}

}
