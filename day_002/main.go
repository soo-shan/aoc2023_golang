package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var input_file string = "example_input.txt"

var input_file string = "input.txt"

func main() {

	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println("Err", err)
	}

	scanner := bufio.NewScanner(file)
	sumGameID := 0
	sumPowerSet := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		parts := strings.Split(line, ":")
		gameID, err := strconv.Atoi(parts[0][5:])

		if err != nil {
			fmt.Println(err)
		}

		if isValidGameP1(parts[1]) {
			sumGameID += gameID
			fmt.Println("sumGameID", sumGameID) //2563
		}

		sumPowerSet += powerOfSet(parts[1])
		fmt.Println("sumPowerSet", sumPowerSet) //70768

	}

}
