package handler

import "testing"

func TestNie(t *testing.T) {
	nie := NieRandom()
	if !NieValid(nie) {
		t.Errorf("Nie invalid, got: %s", nie)
	}
}
