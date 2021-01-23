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
		//fmt.Println("Round ", i+1)
	}

	fmt.Println("Look&Say 40 rounds:")
	fmt.Println(len(result))

	// ------------ PART 2 ------------------------
	fmt.Println("Part 2 - 50 rounds")

	byteResult := []byte(input)
	for i := 0; i < 50; i++ {
		byteResult = lookAndSayTurbo(byteResult)
		//fmt.Println("Round ", i+1)
	}
	fmt.Println("Look&Say 50 rounds:")
	result = string(byteResult)
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

// Function from https://github.com/mevdschee/AdventOfCode2015/blob/master/day10/part2.go
// Which runs much faster
func lookAndSayTurbo(str []byte) []byte {
	ch := str[0]
	count := 0
	result := []byte{}
	for i := 0; i < len(str); i++ {
		if str[i] != ch {
			result = append(result, byte(48+count), ch)
			ch = str[i]
			count = 0
		}
		count++
	}
	result = append(result, byte(48+count), ch)
	return result
}
