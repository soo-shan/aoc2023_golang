package main

import (
	"strconv"
	"strings"
)

func GetCaliberationValueP2(line []rune) (int, error) {

	digitWordMap := map[string]rune{"zero": '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9'}

	var firstDigit, secondDigit rune
	foundFirstDigt, foundSecondDigit := false, false

	for i, character := range line {

		if foundFirstDigt {
			break
		}

		if checkDigit(character) {
			firstDigit = character
			break
		}

		stringToCheck := string(line[i:])

		for word := range digitWordMap {
			if strings.HasPrefix(stringToCheck, word) {
				firstDigit = digitWordMap[word]
				foundFirstDigt = true

				break
			}
		}
	}

	for i := len(line); i > 0; i-- {

		if foundSecondDigit {
			break
		}

		if checkDigit(line[i-1]) {
			secondDigit = line[i-1]
			break
		}

		stringToCheck := string(line[:i])

		for word := range digitWordMap {
			if strings.HasSuffix(stringToCheck, word) {
				secondDigit = digitWordMap[word]
				foundSecondDigit = true

				break
			}
		}

	}
	return strconv.Atoi(string(firstDigit) + string(secondDigit))
}
