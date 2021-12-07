package main

import (
	"fmt"
	"log"
	"os"
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
	strCrabSubmarines := strings.Split(source, ",")

	intCrabSubmarines := []int{}
	for i := 0; i < len(strCrabSubmarines); i++ {
		cSubmarine, _ := strconv.Atoi(strCrabSubmarines[i])
		intCrabSubmarines = append(intCrabSubmarines, cSubmarine)
	}

	// Find min and max pos for testing...
	minPos := 9999
	maxPos := -1000

	for _, crabPos := range intCrabSubmarines {
		if crabPos < minPos {
			minPos = crabPos
		}
		if crabPos > maxPos {
			maxPos = crabPos
		}
	}

	//	fmt.Println("Min/max pos:", minPos, "/", maxPos)

	// Test positions
	fuelUsage := 99999999
	superPos := 0

	for pos := minPos; pos < maxPos+1; pos++ {
		crabFuel := 0
		for _, crabSubPos := range intCrabSubmarines {
			crabFuel += diff(crabSubPos, pos)
		}

		if crabFuel < fuelUsage {
			superPos = pos
			fuelUsage = crabFuel
		}
	}

	fmt.Println("Part 01/")
	fmt.Println("Position:", superPos)
	fmt.Println("Fuel usage", fuelUsage)

	// Part 2 ------

	// reset variables
	fuelUsage = 99999999
	superPos = 0

	for pos := minPos; pos < maxPos+1; pos++ {
		crabFuel := 0
		for _, crabSubPos := range intCrabSubmarines {
			crabFuel += p2diff(crabSubPos, pos)
		}

		if crabFuel < fuelUsage {
			superPos = pos
			fuelUsage = crabFuel
		}
	}

	fmt.Println("Part 02/")
	fmt.Println("Position:", superPos)
	fmt.Println("Fuel usage", fuelUsage)

}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func p2diff(a, b int) int {
	delta := diff(a, b)

	if delta > 1 {
		fuelCost := 0
		for fuelSteps := 1; fuelSteps <= delta; fuelSteps++ {
			fuelCost += fuelSteps
		}
		delta = fuelCost
	}
	return delta
}
