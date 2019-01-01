package fake

import (
	"testing"
)

func TestNie(t *testing.T) {
	nie := NieRandom()
	if !NieValid(nie) {
		t.Errorf("Nie invalid, got: %s", nie)
	}
	if !NieCheckDigit(nie) {
		t.Errorf("Nie invalid digit, got: %s", nie)
	}
}
