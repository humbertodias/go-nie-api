package handler

import "testing"

func TestDni(t *testing.T) {
	dni := DniRandom()
	if !DniValid(dni) {
		t.Errorf("Nie invalid, got: %s", dni)
	}
}
