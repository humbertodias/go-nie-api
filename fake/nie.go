package fake

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

func NieRandom() string {
	firstLetter := "XYZ"[rand.Intn(3)]
	middleNumber := NifRandomNumber(7)
	return fmt.Sprintf("%c%d%c", firstLetter, middleNumber, NifLetter(middleNumber))
}

func NieValid(nif string) bool {
	re := regexp.MustCompile(`(^[XYZ]\d{7})[A-HJ-NP-TV-Z]$`)
	return re.MatchString(nif)
}

func NieCheckFirstLetter(nie string) bool {
	firstLetter := nie[0:1]
	return strings.ContainsAny(firstLetter, "X & Y & Z")
}

func NieCheckDigit(nie string) bool {
	return NieCheckFirstLetter(nie) && NifCheckDigit(nie, 1)
}
