package fake

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
)

func NifLetterFromString(sNumber string) byte {
	number, err := strconv.Atoi(sNumber)
	if err != nil {
		fmt.Errorf("NieLetter invalid: %v", err)
	}
	return NifLetter(number)
}

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
	return fmt.Sprintf("%08d%c", number, NifLetter(number))
}

func NifValid(nif string) bool {
	re := regexp.MustCompile(`(^\d{7}\d?|^\d{8})[A-HJ-NP-TV-Z]$`)
	return re.MatchString(nif)
}

func NifCheckDigit(nif string, startAt int) bool {
	middleNumberString := nif[startAt : len(nif)-1]
	middleNumber, _ := strconv.Atoi(middleNumberString)
	currentDigit := nif[len(nif)-1:][0]
	expectedDigit := NifLetter(middleNumber)
	return currentDigit == expectedDigit
}
