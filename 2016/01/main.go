package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	instructions := strings.Fields(string(fileContent))
	fmt.Println("Total instructions:", len(instructions))

	// Starting position
	posX := 0
	posY := 0

	// Facing North (0 degrees)
	dir := 0

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 01, Part 1: No Time for a Taxicab")

	for _, instruction := range instructions {
		instruction := strings.TrimRight(instruction, ",")
		turnDirection := instruction[0:1]
		moves, _ := strconv.Atoi(instruction[1:])

		// Turn to new direction
		if turnDirection == "L" {
			dir += -1 + 4
		} else {
			dir++
		}

		dir %= 4

		// move
		switch dir {
		case 0:
			posY += moves
		case 1:
			posX += moves
		case 2:
			posY -= moves
		case 3:
			posX -= moves
		}

	}

	fmt.Println("End position:", posX, posY)
	fmt.Println("Total distance: ", Abs(posX)+Abs(posY))

	// Part 2 ------------------------------------------
	fmt.Println()
	fmt.Println("Part 2")

	// Reset starting position
	dir = 0
	posX = 0
	posY = 0

	visitedLocations := make(map[string]bool)
	xFirst := 0
	yFirst := 0
	foundFirst := false

	for _, instruction := range instructions {
		instruction := strings.TrimRight(instruction, ",")
		turnDirection := instruction[0:1]
		moves, _ := strconv.Atoi(instruction[1:])
		// Turn to new direction
		if turnDirection == "L" {
			dir += -1 + 4
		} else {
			dir++
		}

		dir %= 4

		// move step by step to record all locations
		for step := 0; step < moves; step++ {
			switch dir {
			case 0:
				posY++
			case 1:
				posX++
			case 2:
				posY--
			case 3:
				posX--
			}
			// Part 2 data collection
			locationID := strconv.Itoa(posX) + "," + strconv.Itoa(posY)
			if visitedLocations[locationID] && !foundFirst {
				xFirst = posX
				yFirst = posY
				foundFirst = true
			}
			visitedLocations[locationID] = true
		}
	}

	fmt.Println("Distance to the first location you visit twice?")
	fmt.Println("First position visited twice (x,y):", xFirst, yFirst)
	fmt.Println("Total distance to first: ", Abs(xFirst)+Abs(yFirst))

}

// Abs helper function to get Absolute value of integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
