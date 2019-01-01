package fake

import (
	"fmt"
	"math/rand"
	"regexp"
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
