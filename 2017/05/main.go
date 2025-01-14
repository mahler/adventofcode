package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	instructions2 := make([]int, len(instructions))
	copy(instructions1, instructions)
	copy(instructions2, instructions)

	fmt.Printf("star 1: %d\n", process1(instructions1))
	fmt.Printf("star 2: %d\n", process2(instructions2))
}

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
