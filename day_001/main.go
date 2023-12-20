package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Err")
	}

	scanner := bufio.NewScanner(file)

	caliberationValue := 0

	caliberationSumP1 := 0
	caliberationSumP2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line", line)
		caliberationValue, err = GetCaliberationValueP1([]rune(line))

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("caliberation value p1", caliberationValue)
		caliberationSumP1 += caliberationValue

		caliberationValue, err = GetCaliberationValueP2([]rune(line))

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("caliberation value p2", caliberationValue)
		caliberationSumP2 += caliberationValue

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1 caliberation Sum Value:", caliberationSumP1) //56042
	// caliberateSumP2 := day1_2()
	fmt.Println("Part 2 caliberation Sum Value:", caliberationSumP2) //55358

}
