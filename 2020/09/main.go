package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("data.xmas")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	records := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Records in Dataset:", len(records))
	fmt.Println()
	fmt.Println("Day 09: PART 1 - Encoding Error")

	numbers := make([]int, len(records))

	for i, record := range records {
		numbers[i], _ = strconv.Atoi(record)
	}

	failedNumber := 0
	// Starting position (after preamble)
	pos := 25
	for n := pos; n < len(records); n++ {

		nextNumber := numbers[n]
		foundMatch := false

		for x := 1; x < 26; x++ {
			for y := 1; y < 26; y++ {
				sum := numbers[n-x] + numbers[n-y]
				if nextNumber == sum {
					foundMatch = true
				}
				//fmt.Printf("%v + %v = %v\n", numbers[n-x], numbers[n-y], sum)

			}
		}

		if !foundMatch {
			//fmt.Println("No match:", nextNumber)
			// Save for use in part 2
			failedNumber = nextNumber
		}
		//break
	}
	fmt.Println("answer for part 1:", failedNumber)

	fmt.Println()
	fmt.Println("Day 09: PART 2 - Series match")

	ixHigh := 0
	ixLow := 0

	for n := pos; n < len(records); n++ {
		//fmt.Println("*", n)
		foundMatch := false
		subSum := numbers[n]
		for y := 1; y < pos+1; y++ {
			subSum += numbers[n-y]
			if subSum == failedNumber {
				foundMatch = true
				// +1 below as slices has index 0
				ixHigh = n
				ixLow = n - y
				break
			}
		}
		if foundMatch {
			break
		}
		//break
	}

	ixLowNum := ixLow
	ixHighNum := ixHigh
	for z := ixLow; z <= ixHigh; z++ {
		if numbers[ixLow] > numbers[z] {
			ixLowNum = z
		} else if numbers[ixHigh] < numbers[z] {
			ixHighNum = z
		}
	}
	fmt.Println("ixLow:", ixLowNum, " ixHigh:", ixHighNum)
	fmt.Printf("%v + %v = %v\n", numbers[ixLowNum], numbers[ixHighNum], numbers[ixLowNum]+numbers[ixHighNum])
}
