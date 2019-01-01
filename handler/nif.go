package handler

import (
	"fmt"
	"math/rand"
	"regexp"
)

func NifLetter(number int) byte {
	nifSeq := "TRWAGMYFPDXBNJZSQVHLCKE"
	index := int(number % len(nifSeq))
	return nifSeq[index]
}

func NifRandomNumber() int {
	return rand.Intn(100000000)
}

func NifRandom() string {
	number := NifRandomNumber()
	return fmt.Sprintf("%d%c", number, NifLetter(number))
}

func NifValid(nif string) bool {
	re := regexp.MustCompile(`(^\d{7}\d?|^\d{8})[A-HJ-NP-TV-Z]$`)
	return re.MatchString(nif)
}
