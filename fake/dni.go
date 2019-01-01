package fake

import (
	"fmt"
	"regexp"
)

func DniRandom() string {
	number := NifRandomNumber(8)
	return fmt.Sprintf("%08d%c", number, NifLetter(number))
}

func DniValid(nif string) bool {
	re := regexp.MustCompile(`(^\d{7}\d?|^\d{8})[A-HJ-NP-TV-Z]$`)
	return re.MatchString(nif)
}

func DniCheckDigit(dni string) bool {
	return NifCheckDigit(dni, 0)
}
