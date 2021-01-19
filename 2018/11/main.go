package main

import (
	"fmt"
	"strconv"
)

const gridSerial = 9424

func main() {

	fuelCell := make(map[int]map[int]int)

	for row := 1; row <= 300; row++ {
		fuelCell[row] = make(map[int]int)
		for col := 1; col <= 300; col++ {
			fuelCell[row][col] = calcFuel(col, row, gridSerial)

		}
	}

	colMax, rowMax, topScore := 0, 0, 0
	for row := 1; row <= 298; row++ {
		for col := 1; col <= 298; col++ {
			cellScore := calcScore(fuelCell, row, col, 3)
			if cellScore > topScore {
				topScore = cellScore
				colMax = col
				rowMax = row
			}

		}
	}
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("Day 11, Part 1: Chronal Charge")
	fmt.Println("TopScore:", topScore)
	fmt.Printf("X,Y: %d,%d\n", colMax, rowMax)
	// --------------------------

	fmt.Println("Part 2")
	fmt.Println("Slow... patience please")
	colMax, rowMax, topScore = 0, 0, 0
	gridMax := 0
	for row := 1; row <= 298; row++ {
		for col := 1; col <= 298; col++ {
			// min grid 3x3.
			for gSize := 3; gSize < len(fuelCell); gSize++ {

				cellScore := calcScore(fuelCell, row, col, gSize)
				// fmt.Println(row, col, gSize, cellScore)
				if cellScore > topScore {
					topScore = cellScore
					colMax = col
					rowMax = row
					gridMax = gSize
				}
			}

		}
	}
	fmt.Printf("%d,%d,%d\n", colMax, rowMax, gridMax)
}

func calcScore(grid map[int]map[int]int, x, y, gridSize int) int {
	theScore := 0

	for a := 0; a < gridSize; a++ {
		for b := 0; b < gridSize; b++ {
			theScore += grid[x+a][y+b]
		}
	}

	return theScore
}

func calcFuel(x, y, grid int) int {
	powerLevel := 0
	//	Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackID := x + 10
	// Begin with a power level of the rack ID times the Y coordinate.
	powerLevel = rackID * y
	// Increase the power level by the value of the grid serial number (your puzzle input).
	powerLevel += grid
	// Set the power level to itself multiplied by the rack ID.
	powerLevel = powerLevel * rackID
	// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	if powerLevel > 100 {

		strPower := strconv.Itoa(powerLevel)
		strPower = strPower[len(strPower)-3 : len(strPower)-2]
		powerLevel, _ = strconv.Atoi(strPower)
	} else {
		powerLevel = 0
	}

	// Subtract 5 from the power level.
	powerLevel -= 5

	return powerLevel
}
