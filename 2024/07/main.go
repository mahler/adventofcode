package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var decimalRep []int

func parseData(inputData string) [][]int {
	var parsedData [][]int
	lines := strings.Split(strings.TrimSpace(inputData), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		goal, _ := strconv.Atoi(parts[0])
		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))

		for i, numStr := range numberStrings {
			numbers[i], _ = strconv.Atoi(numStr)
		}

		parsedData = append(parsedData, append([]int{goal}, numbers...))
	}
	return parsedData
}

func concatenate(leftSide, rightSide int) int {
	for _, element := range decimalRep {
		if element > rightSide {
			return leftSide*element + rightSide
		}
	}
	return 0
}

func isValid(goal, result int, remaining []int, concatenation bool) bool {
	if len(remaining) == 0 {
		return goal == result
	}

	current := remaining[0]

	// Try addition
	if isValid(goal, result+current, remaining[1:], concatenation) {
		return true
	}

	// Try multiplication
	if isValid(goal, result*current, remaining[1:], concatenation) {
		return true
	}

	// Try concatenation if allowed
	if concatenation && isValid(goal, concatenate(result, current), remaining[1:], concatenation) {
		return true
	}

	return false
}

func isValidEquation(equation []int, concatenation bool) bool {
	goal := equation[0]
	numbers := equation[1:]
	return isValid(goal, 0, numbers, concatenation)
}

func main() {
	fileName := "puzzle.txt"

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	equations := parseData(string(content))

	var remaining [][]int
	validSum := 0

	for _, equation := range equations {
		if isValidEquation(equation, false) {
			validSum += equation[0]
		} else {
			remaining = append(remaining, equation)
		}
	}

	fmt.Println("Part 1: What is their total calibration result?")
	fmt.Println(validSum)

	// part 2
	withConcatenation := 0
	for _, equation := range remaining {
		if isValidEquation(equation, true) {
			withConcatenation += equation[0]
		}
	}

	fmt.Println()
	fmt.Println("Part 2: Using your new knowledge of elephant hiding spots, determine which equations could possibly be true.")
	fmt.Println("What is their total calibration result?")
	fmt.Println(validSum + withConcatenation)
}
