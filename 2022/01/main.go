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

	calorieCounts := []int{}
	currentElfCount := 0
	for _, fileRow := range fileLines {

		if fileRow == "" {
			calorieCounts = append(calorieCounts, currentElfCount)
			currentElfCount = 0
			continue
		}

		calories, _ := strconv.Atoi(fileRow)
		currentElfCount += calories
	}

	calorieCounts = append(calorieCounts, currentElfCount)

	topOne := 0
	topTwo := 0
	topThree := 0
	for _, calorieCount := range calorieCounts {
		if calorieCount > topThree {
			topThree = calorieCount
		}

		if calorieCount > topTwo {
			topTwo, topThree = calorieCount, topTwo
		}

		if calorieCount > topOne {
			topOne, topTwo = calorieCount, topOne
		}
	}

	fmt.Println()
	fmt.Println("Day 1: Calorie Counting")
	fmt.Println("How many total Calories is that Elf carrying?")
	fmt.Println(topOne)

	// ---------- Part 2
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("How many Calories are those Elves carrying in total?")
	fmt.Println(topOne + topTwo + topThree)
}
