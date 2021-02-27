package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 3: Squares With Three Sides")
	fileContent, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileRows := strings.Split(string(fileContent), "\n")

	validTriangles := 0

	for _, fileRow := range fileRows {
		fields := strings.Fields(fileRow)

		sideOne, _ := strconv.Atoi(fields[0])
		sideTwo, _ := strconv.Atoi(fields[1])
		sideThree, _ := strconv.Atoi(fields[2])

		if sideOne+sideTwo > sideThree && sideOne+sideThree > sideTwo && sideTwo+sideThree > sideOne {
			validTriangles++
		}
	}

	fmt.Println("how many of the listed triangles are possible?")
	fmt.Println(validTriangles)
	// ----------------------------------------
	fmt.Println()
	fmt.Println("Part 2: Vertical Triangles")

	// Reset reused vars
	validTriangles = 0

	// Build triangle data 2D slice
	triData := make([][]int, len(fileRows))
	for i, fileRow := range fileRows {
		fields := strings.Fields(fileRow)

		sideOne, _ := strconv.Atoi(fields[0])
		sideTwo, _ := strconv.Atoi(fields[1])
		sideThree, _ := strconv.Atoi(fields[2])

		dataLine := []int{sideOne, sideTwo, sideThree}
		triData[i] = dataLine
	}

	// Walk the triangle data for vertical triangles
	for t := 0; t < len(fileRows); t += 3 {
		for u := 0; u < 3; u++ {
			if triData[t][u]+triData[t+1][u] > triData[t+2][u] && triData[t+1][u]+triData[t+2][u] > triData[t][u] && triData[t][u]+triData[t+2][u] > triData[t+1][u] {
				validTriangles++
			}
			//fmt.Println(t+u, "#", triData[t][u], "-", triData[t+1][u], "-", triData[t+2][u])
		}
	}
	fmt.Println("How many of the listed vertical triangles are possible?")
	fmt.Println(validTriangles)
}
