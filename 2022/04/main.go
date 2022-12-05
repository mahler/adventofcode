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
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	pairs := 0

	for _, row := range fileLines {
		sections := strings.Split(row, ",")

		a, b := sections[0], sections[1]

		partsA := strings.Split(a, "-")
		partsB := strings.Split(b, "-")

		a1, _ := strconv.Atoi(partsA[0])
		a2, _ := strconv.Atoi(partsA[1])

		b1, _ := strconv.Atoi(partsB[0])
		b2, _ := strconv.Atoi(partsB[1])

		if (b1 >= a1 && b2 <= a2) || (a1 >= b1 && a2 <= b2) {
			pairs++
			continue
		}

	}

	fmt.Println()
	fmt.Println("Day 4:Camp Cleanup")
	fmt.Println("In how many assignment pairs does one range fully contain the other?")
	fmt.Println(pairs)

	// Part 2 ...........
	pairs = 0

	for _, row := range fileLines {
		sections := strings.Split(row, ",")

		a, b := sections[0], sections[1]

		partsA := strings.Split(a, "-")
		partsB := strings.Split(b, "-")

		a1, _ := strconv.Atoi(partsA[0])
		a2, _ := strconv.Atoi(partsA[1])

		b1, _ := strconv.Atoi(partsB[0])
		b2, _ := strconv.Atoi(partsB[1])

		if (a1 >= b1 && a1 <= b2) || (a2 >= b1 && a2 <= b2) || (b1 >= a1 && b1 <= a2) || (b2 >= a1 && b2 <= a2) {
			pairs++
			continue
		}

	}

	fmt.Println()
	fmt.Println("Part 2:")
	fmt.Println("In how many assignment pairs do the ranges overlap?")
	fmt.Println(pairs)

}
