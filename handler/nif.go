package handler

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
)

func NifLetter(number int) byte {
	nifSeq := "TRWAGMYFPDXBNJZSQVHLCKE"
	index := int(number % len(nifSeq))
	return nifSeq[index]
}

func NifRandomNumber(length int) int {
	number := int(math.Pow10(length))
	return rand.Intn(number)
}

func NifRandom() string {
	number := NifRandomNumber(8)
	return fmt.Sprintf("%d%c", number, NifLetter(number))
}

func NifValid(nif string) bool {
	re := regexp.MustCompile(`(^\d{7}\d?|^\d{8})[A-HJ-NP-TV-Z]$`)
	return re.MatchString(nif)
}
