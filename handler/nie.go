package handler

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

func NieRandom() string {
	first_letter := "XYZ"[rand.Intn(3)]
	number_part := NifRandomNumber()
	number_for_calculation := fmt.Sprintf("%d%d", int(first_letter), number_part)
	number_for_calculation_int, _ := strconv.Atoi(number_for_calculation)
	return fmt.Sprintf("%c%d%c", first_letter, number_part, NifLetter(number_for_calculation_int))
}

func NieValid(nif string) bool {
	re := regexp.MustCompile(`(^[XYZ]\d{7}\d?|^\d{8})[A-HJ-NP-TV-Z]$`)
	return re.MatchString(nif)
}
