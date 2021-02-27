package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println()
	fmt.Println("Day 01, part 1: The Tyranny of the Rocket Equation")
	requiredFuel := 0
	for _, fileLine := range fileLines {
		moduleMass, _ := strconv.Atoi(fileLine)
		requiredFuel += massToFuel(moduleMass)
	}
	fmt.Println("Total fuel need:", requiredFuel)

	fmt.Println()
	fmt.Println("Part 2: Fuel need with fuelweight")
	// reset fuel counter
	requiredFuel = 0
	for _, fileLine := range fileLines {
		moduleMass, _ := strconv.Atoi(fileLine)
		requiredFuel += massToFuelWithFuel(moduleMass)
	}
	fmt.Println()
	fmt.Println("New fuel need:", requiredFuel)
}

func massToFuel(massSize int) int {
	return massSize/3 - 2
}

func massToFuelWithFuel(massSize int) int {
	neededFuel := massToFuel(massSize)
	if neededFuel > 0 {
		return neededFuel + massToFuelWithFuel(neededFuel)
	}
	return 0
}
