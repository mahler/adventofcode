package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("DAY03, Part 1: No Matter How You Slice It")
	fmt.Println("Total puzzle input: ", len(fileContent))

	// ----------------------------------------------------------
	squareSize := 1000
	fabric := make([][]int, squareSize)
	for i := 0; i < squareSize; i++ {
		fabric[i] = make([]int, squareSize)
	}

	// ----------------------------------------------------------
	// RegExp tester online: https://regoio.herokuapp.com/
	claimRegExp := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

	for _, claim := range fileContent {
		fields := claimRegExp.FindStringSubmatch(claim)
		//fmt.Println(fields)
		//claimID := fields[1]
		xStart, _ := strconv.Atoi(fields[2])
		yStart, _ := strconv.Atoi(fields[3])
		xWidth, _ := strconv.Atoi(fields[4])
		yHeight, _ := strconv.Atoi(fields[5])

		xEnd := xStart + xWidth
		yEnd := yStart + yHeight

		// fmt.Println("Claim", claimID, "X (", xStart, ",", xEnd, ") - Y(", yStart, ",", yEnd, ")")

		for x := xStart; x < xEnd; x++ {
			for y := yStart; y < yEnd; y++ {
				fabric[x][y]++
			}
		}
	}

	// Count squares with 2 or more claims
	multiClaimSquares := 0
	for xCount := 0; xCount < squareSize; xCount++ {
		for yCount := 0; yCount < squareSize; yCount++ {
			if fabric[xCount][yCount] > 1 {
				multiClaimSquares++
			}
		}
	}

	fmt.Println("Squares claimed by at least two:", multiClaimSquares)
	// ----------------------------------------------------------------------------------------------------

	fmt.Println()
	fmt.Println("Part 2")

	for _, claim := range fileContent {
		checkFailed := false
		fields := claimRegExp.FindStringSubmatch(claim)
		//fmt.Println(fields)
		claimID := fields[1]
		xStart, _ := strconv.Atoi(fields[2])
		yStart, _ := strconv.Atoi(fields[3])
		xWidth, _ := strconv.Atoi(fields[4])
		yHeight, _ := strconv.Atoi(fields[5])

		xEnd := xStart + xWidth
		yEnd := yStart + yHeight

		// fmt.Println("Claim", claimID, "X (", xStart, ",", xEnd, ") - Y(", yStart, ",", yEnd, ")")

		for x := xStart; x < xEnd; x++ {
			for y := yStart; y < yEnd; y++ {
				if fabric[x][y] > 1 {
					checkFailed = true
				}
			}
		}

		if !checkFailed {
			fmt.Println("Claim", claimID, "is not claimed by others")
			break
		}
	}

}
