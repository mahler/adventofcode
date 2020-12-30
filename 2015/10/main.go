package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "1321131112"

	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 01: Elves Look, Elves Say")
	fmt.Println("Patience, this may be slow...")
	// Copy to result to save original input
	result := input
	for i := 0; i < 40; i++ {
		result = lookAndSay(result)
	}

	fmt.Println("Look&Say 40 rounds:")
	fmt.Println(len(result))
}

func lookAndSay(inputstring string) string {
	returnString := ""
	count := 0

	var current rune
	for _, c := range inputstring {
		if c != current {
			if count > 0 {
				returnString += strconv.Itoa(count) + string(current)
			}

			current = c
			count = 1
		} else {
			count++
		}
	}

	returnString += strconv.Itoa(count) + string(current)

	return returnString
}
