package main

import (
	"fmt"
	"strconv"
	"strings"
)

func powerOfSet(line string) int {

	maxRed, maxGreen, maxBlue := 0, 0, 0
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
					maxRed = count
				}

			case "green":
				if maxGreen < count {
					maxGreen = count
				}

			case "blue":
				if maxBlue < count {
					maxBlue = count
				}
			}

		}

	}

	fmt.Println("maxRed:", maxRed, "maxGreen:", maxGreen, "maxBlue:", maxBlue)

	return maxRed * maxGreen * maxBlue

}
