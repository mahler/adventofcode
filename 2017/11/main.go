package main

import (
	"fmt"
	"os"
	"strings"
)

// getHexDistance calculates the hex grid distance from origin
func getHexDistance(x, y, z int) int {
	// In a hex grid using cubic coordinates, the sum of absolute values
	// will always be even, so this division will always result in an integer
	return (abs(x) + abs(y) + abs(z)) / 2
}

// abs returns the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Read input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split into directions
	directions := strings.Split(strings.TrimSpace(string(data)), ",")

	var x, y, z int
	distances := make([]int, 0, len(directions))

	// Process each direction
	for _, d := range directions {
		switch d {
		case "n":
			y++
			z--
		case "s":
			y--
			z++
		case "ne":
			x++
			z--
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		case "se":
			x++
			y--
		}

		distances = append(distances, getHexDistance(x, y, z))
	}

	// Calculate final distance
	finalDist := getHexDistance(x, y, z)

	// Find maximum distance reached
	maxDist := distances[0]
	for _, d := range distances {
		if d > maxDist {
			maxDist = d
		}
	}
	fmt.Println("Part 1: You need to determine the fewest number of steps required to reach him.")
	fmt.Println(finalDist)
	fmt.Println()
	fmt.Println("Part 2: How many steps away is the furthest he ever got from his starting position?")
	fmt.Println(maxDist)
}
