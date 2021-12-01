package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	numIncreases := 0
	numDecreases := 0
	number := 0
	tmpNumber := 0

	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println()

	for _, fileRow := range fileLines {
		//	fmt.Println(fileRow)
		if number == 0 {
			number, _ = strconv.Atoi(fileRow)
		} else {
			tmpNumber, _ = strconv.Atoi(fileRow)
			if tmpNumber < number {
				numDecreases++
			} else if tmpNumber > number {
				numIncreases++
			} else {
				fmt.Println("Equal")
			}
			number = tmpNumber
		}

	}

	fmt.Println()
	fmt.Println("Day 1: Sonar Sweep")
	fmt.Println("Increases:", numIncreases)
	// fmt.Println("Decreases:", numDecreases)

	// ---------- Part 2
	numIncreases = 0 // Reset for Part 2

	var theNumbers []int
	for _, fileRow := range fileLines {
		number, _ = strconv.Atoi(fileRow)
		theNumbers = append(theNumbers, number)
	}

	var baseNumber = theNumbers[0] + theNumbers[1] + theNumbers[2]
	var compareNumber = 0

	for x := 2; x < len(theNumbers)-2; x += 1 {
		compareNumber = theNumbers[x] + theNumbers[x+1] + theNumbers[x+2]
		//		fmt.Println(baseNumber, " <-> ", compareNumber)

		if baseNumber < compareNumber {
			numIncreases++
		}
		baseNumber = compareNumber
	}
	fmt.Println()
	fmt.Println("Part 2:", numIncreases)

}
