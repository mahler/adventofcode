package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func process1(instructions []int) uint32 {
	currentLocation := 0
	var numSteps uint32 = 0

	for currentLocation < len(instructions) {
		nextStep := instructions[currentLocation]
		instructions[currentLocation]++
		currentLocation += nextStep
		numSteps++
	}
	return numSteps
}

func process2(instructions []int) uint32 {
	currentLocation := 0
	var numSteps uint32 = 0

	for currentLocation < len(instructions) {
		nextStep := instructions[currentLocation]
		if nextStep >= 3 {
			instructions[currentLocation]--
		} else {
			instructions[currentLocation]++
		}
		currentLocation += nextStep
		numSteps++
	}
	return numSteps
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var instructions []int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		instructions = append(instructions, num)
	}

	// Create copies of instructions for each process
	instructions1 := make([]int, len(instructions))
	copy(instructions1, instructions)

	fmt.Println("Part 1: How many steps does it take to reach the exit?")
	fmt.Println(process1(instructions1))

	// Part 2
	instructions2 := make([]int, len(instructions))
	copy(instructions2, instructions)

	fmt.Println()
	fmt.Println("Part 2: How many steps does it now take to reach the exit?")
	fmt.Println(process2(instructions2))
}
