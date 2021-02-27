package main

import (
	"fmt"
	"log"
)

func main() {

	data, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	dataStr := string(data)

	fmt.Println()
	fmt.Println("2019 - Day 01 Part 1")
	xPos, yPos := 0, 0

	locations := make(map[string]int)
	for c := 0; c < len(dataStr); c++ {
		// move santa
		switch string(dataStr[c]) {
		case "<":
			xPos--
		case ">":
			xPos++
		case "^":
			yPos++
		case "v":
			yPos--
		}

		//fmt.Println("xpos:", xPos, "yPos", yPos)
		// make virtual address
		virtualAddress := fmt.Sprintf("%d#%d", xPos, yPos)
		// Been here?
		if _, ok := locations[virtualAddress]; ok {
			locations[virtualAddress]++
		} else {
			locations[virtualAddress] = 1
		}
	}
	fmt.Println("How many houses receive at least one present?")
	// Adding one for starting location.
	fmt.Println(len(locations) + 1)
	// ------------ PART 2 ------------------------

	fmt.Println()
	fmt.Print("Part 2: RoboSanta asssist")

	xSanta, ySanta, xRobo, yRobo := 0, 0, 0, 0

	newLocations := make(map[string]int)
	for c := 0; c < len(dataStr); c++ {
		// move santa
		switch string(dataStr[c]) {
		case "<":
			xSanta--
		case ">":
			xSanta++
		case "^":
			ySanta++
		case "v":
			ySanta--
		}
		// make virtual address
		virtualAddress := fmt.Sprintf("%d#%d", xSanta, ySanta)
		// Been here?
		if _, ok := newLocations[virtualAddress]; ok {
			newLocations[virtualAddress]++
		} else {
			newLocations[virtualAddress] = 1
			fmt.Println("NewSanta")
		}

		c++
		// move roboSanta
		switch string(dataStr[c]) {
		case "<":
			xRobo--
		case ">":
			xRobo++
		case "^":
			yRobo++
		case "v":
			yRobo--
		}
		// make virtual address
		virtualAddress = fmt.Sprintf("%d#%d", xRobo, yRobo)
		// Been here?
		if _, ok := newLocations[virtualAddress]; ok {
			newLocations[virtualAddress]++
		} else {
			newLocations[virtualAddress] = 1
			fmt.Println("NewRobo")
		}
	}
	fmt.Println("How many houses receive at least one present?")
	fmt.Println(len(newLocations))
}
