package fake

import "testing"

func TestDni(t *testing.T) {
	Seed(1234)
	dni := DniRandom()
	if !DniValid(dni) || dni != "75231682Q" {
		t.Errorf("Nie invalid, got: %s", dni)
	}
}
