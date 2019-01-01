package handler

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

func NieRandom() string {
	available_letters := "XYZ"
	index := rand.Intn(len(available_letters))
	first_letter := available_letters[index]
	first_letter_value := int(first_letter)
	number_part := NifRandomNumber()
	number_for_calculation := fmt.Sprintf("%d%d", first_letter_value, number_part)
	number_for_calculation_int, _ := strconv.Atoi(number_for_calculation)
	return fmt.Sprintf("%c%d%c", first_letter, number_part, NifLetter(number_for_calculation_int))
}

func NieValid(nif string) bool {
	re := regexp.MustCompile(`(^[XYZ]\d{7}\d?|^\d{8})[A-HJ-NP-TV-Z]$`)
	return re.MatchString(nif)
}
