package fake

import (
	"testing"
)

func TestNie(t *testing.T) {
	Seed(1234)
	nie := NieRandom()
	if !NieValid(nie) || nie != "Z1458315T" {
		t.Errorf("Nie invalid, got: %s", nie)
	}
}
