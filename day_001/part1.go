package main

import (
	"strconv"
)

func checkDigit(character rune) bool {
	if character >= '0' && character <= '9' {
		return true
	} else {
		return false
	}
}

func GetCaliberationValueP1(line []rune) (int, error) {
	var firstDigit, secondDigit rune
	foundFirstDigit := false
	for _, character := range line {
		if !foundFirstDigit && checkDigit(character) {
			foundFirstDigit = true
			firstDigit = character
			secondDigit = character
		} else if checkDigit(character) {
			secondDigit = character
		}
	}

	return strconv.Atoi(string(firstDigit) + string(secondDigit))

}
