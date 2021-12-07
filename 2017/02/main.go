package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	// Setup
	source := string(fileContent)

	sum := 0
	for _, line := range strings.Split(source, "\n") {
		intSlice := []int{}
		s := strings.Split(line, "\t")
		for _, r := range s {
			number, _ := strconv.Atoi(string(r))
			intSlice = append(intSlice, number)
		}

		sort.Ints(intSlice)
		sum += intSlice[len(intSlice)-1] - intSlice[0]
	}

	fmt.Println()
	fmt.Println("2017 - Day 02 part 1")
	fmt.Println("checksum is", sum)

	//  -----
	fmt.Println()
	fmt.Println("Part 2")

	// reset sum
	sum = 0
	for _, line := range strings.Split(source, "\n") {
		intSlice := []int{}
		s := strings.Split(line, "\t")
		for _, r := range s {
			number, _ := strconv.Atoi(string(r))
			intSlice = append(intSlice, number)
		}

		sum += divisibleValues(intSlice)
	}

	fmt.Println("sum of divisibles is", sum)
}

func divisibleValues(ints []int) int {
	sort.Ints(ints)
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[j]%ints[i] == 0 {
				return ints[j] / ints[i]
			}
		}
	}
	// Should never go here
	return -1
}
