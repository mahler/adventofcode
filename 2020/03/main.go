package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("map.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	strSlice := strings.Split(strings.TrimSpace(string(data)), "\n")
	mapLength := len(strSlice)
	mapRows := make([][]rune, mapLength)

	//	matrix := make([][]byte, mapLength)
	for i, strVal := range strSlice {
		ultraMapRowWidth := ""

		// Expand mapWidth (3x for part one, but part two will need 7x the width):
		for len(ultraMapRowWidth) < (7 * mapLength) {
			ultraMapRowWidth += strVal
		}

		row := make([]rune, len(ultraMapRowWidth))
		//fmt.Println(i, ":", ultraMapRowWidth)
		for j, rowVal := range ultraMapRowWidth {
			row[j] = rowVal
		}
		mapRows[i] = row
	}

	// Building map Matrix in mapRows done, time to walk the map.
	// fmt.Println(mapRows)
	fmt.Println()
	fmt.Println("Day 3, Part 1: Toboggan Trajectory")
	trees := 0
	noTrees := 0
	xPos := 0

	for yPos := 0; yPos < mapLength; yPos++ {

		//fmt.Println("Checking:", string(mapRows[yPos][xPos]))
		if string(mapRows[yPos][xPos]) == "." {
			noTrees++
		} else if string(mapRows[yPos][xPos]) == "#" {
			trees++
		} else {
			fmt.Println("Panic!")
		}
		// move xPos here.
		xPos += 3
		//fmt.Println("Current Xpos:", xPos, "- Current yPos:", yPos)
	}

	fmt.Println("Found", trees, "trees")
	fmt.Println("Found", noTrees, " without trees")
	fmt.Println("Checked", trees+noTrees, "positions")
	fmt.Println()

	fmt.Println("Day 3, Part 2: Alternate lope check")
	slopeTwoTrees := trees

	// Slope 1: Right 1, down 1.
	xPos = 0
	trees = 0
	for yPos := 0; yPos < mapLength; yPos++ {
		if string(mapRows[yPos][xPos]) == "#" {
			trees++
		}
		xPos += 1
	}
	slopeOneTrees := trees

	// Slope 3: Right 5, down 1.
	xPos = 0
	trees = 0
	for yPos := 0; yPos < mapLength; yPos++ {
		if string(mapRows[yPos][xPos]) == "#" {
			trees++
		}
		xPos += 5
	}

	slopeThreeTrees := trees

	// Slope 4: Right 7, down 1.
	xPos = 0
	trees = 0
	for yPos := 0; yPos < mapLength; yPos++ {
		if string(mapRows[yPos][xPos]) == "#" {
			trees++
		}
		xPos += 7
	}

	slopeFourTrees := trees

	// Slope 5: RRight 1, down 2.
	xPos = 0
	trees = 0
	for yPos := 0; yPos < mapLength; yPos++ {
		if string(mapRows[yPos][xPos]) == "#" {
			trees++
		}
		xPos += 1
		yPos += 1 // Force the 2 down.
	}

	slopeFiveTrees := trees

	// ------ Reporting

	fmt.Println("Slope one: ", slopeOneTrees, "trees")
	fmt.Println("Slope two: ", slopeTwoTrees, "trees")
	fmt.Println("Slope three: ", slopeThreeTrees, "trees")
	fmt.Println("Slope four: ", slopeFourTrees, "trees")
	fmt.Println("Slope Five: ", slopeFiveTrees, "trees")

	fmt.Println("Total Trees (so far):", slopeOneTrees*slopeTwoTrees*slopeThreeTrees*slopeFourTrees*slopeFiveTrees)

	fmt.Println()

}
