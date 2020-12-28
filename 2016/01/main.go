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
		lengthDirection, _ := strconv.Atoi(instruction[1:])

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
			posY += lengthDirection
		case 1:
			posX += lengthDirection
		case 2:
			posY -= lengthDirection
		case 3:
			posX -= lengthDirection
		}
	}
	fmt.Println("End position:", posX, posY)
	fmt.Println("Total distance: ", Abs(posX)+Abs(posY))

	// Part 2 ------------------------------------------
}

// Abs helper function to get Absolute value of integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
