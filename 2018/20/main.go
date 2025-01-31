package main

import (
	"fmt"
	"os"
	"strings"
)

// Direction represents movement in x,y coordinates
type Direction struct {
	dx, dy int
}

func main() {
	// Read input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	input := strings.TrimSpace(string(data))

	// Direction mapping
	directions := map[rune]Direction{
		'N': {0, -1},
		'E': {1, 0},
		'S': {0, 1},
		'W': {-1, 0},
	}

	// Stack for position tracking
	positions := make([][2]int, 0)

	// Current position
	x, y := 5000, 5000
	prevX, prevY := x, y

	// Map to store connections between rooms
	connections := make(map[[2]int]map[[2]int]bool)

	// Map to store distances
	distances := make(map[[2]int]int)

	// Process input string, skipping first and last character
	for _, c := range input[1 : len(input)-1] {
		//fmt.Printf("%c %d\n", c, len(positions))

		switch c {
		case '(':
			positions = append(positions, [2]int{x, y})
		case ')':
			if len(positions) > 0 {
				x, y = positions[len(positions)-1][0], positions[len(positions)-1][1]
				positions = positions[:len(positions)-1]
			}
		case '|':
			if len(positions) > 0 {
				x, y = positions[len(positions)-1][0], positions[len(positions)-1][1]
			}
		default:
			if dir, ok := directions[c]; ok {
				x += dir.dx
				y += dir.dy

				// Initialize map entry if it doesn't exist
				if _, exists := connections[[2]int{x, y}]; !exists {
					connections[[2]int{x, y}] = make(map[[2]int]bool)
				}

				// Add connection
				connections[[2]int{x, y}][[2]int{prevX, prevY}] = true

				// Update distances
				currentDist := distances[[2]int{prevX, prevY}] + 1
				if existing, exists := distances[[2]int{x, y}]; !exists || currentDist < existing {
					distances[[2]int{x, y}] = currentDist
				}
			}
		}

		prevX, prevY = x, y
	}

	// Find maximum distance
	maxDist := 0
	roomsOver1000 := 0
	for _, dist := range distances {
		if dist > maxDist {
			maxDist = dist
		}
		if dist >= 1000 {
			roomsOver1000++
		}
	}

	fmt.Println("Part 1:What is the largest number of doors you would be required to pass through to reach a room?")
	fmt.Println(maxDist)
	fmt.Println()
	fmt.Println("Part 2: How many rooms have a shortest path from your current location that pass through at least 1000 doors?")
	fmt.Println(roomsOver1000)
}
