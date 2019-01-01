package fake

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

func NieRandom() string {
	firstLetterIndex := rand.Intn(3)
	firstLetter := "XYZ"[firstLetterIndex]
	middleNumber := NifRandomNumber(7)
	middleNumberWithFirstLetter := fmt.Sprintf("%d%d", firstLetterIndex, middleNumber)
	return fmt.Sprintf("%c%07d%c", firstLetter, middleNumber, NifLetterFromString(middleNumberWithFirstLetter))
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
