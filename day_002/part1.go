package main

import (
	"fmt"
	"strconv"
	"strings"
)

var maxRed int = 12
var maxGreen int = 13
var maxBlue int = 14

func isValidGameP1(line string) bool {
	for _, game := range strings.Split(line, ";") {

		// fmt.Println(game)
		for _, cube := range strings.Split(game, ",") {
			// fmt.Println(cube)
			cube = strings.TrimSpace(cube)
			color := strings.Split(cube, " ")[1]
			count, err := strconv.Atoi(strings.Split(cube, " ")[0])
			if err != nil {
				fmt.Println(err)
			}

			switch color {
			case "red":
				if maxRed < count {
					return false
				}

			case "green":
				if maxGreen < count {
					return false
				}

			case "blue":
				if maxBlue < count {
					return false
				}
			}

		}

	}
	return true

}
