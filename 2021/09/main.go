package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("2021-09 Part 1")
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	// Setup
	fileRows := string(fileContent)
	fileLines := strings.Split(fileRows, "\n")

	caves := make(map[int]map[int]int)
	for i, row := range fileLines {
		caveRow := make(map[int]int)
		for xPos := 0; xPos < len(row); xPos++ {
			value, _ := strconv.Atoi(string(row[xPos]))
			caveRow[xPos] = value
		}
		caves[i] = caveRow
	}

	sum := 0
	for i := 0; i < len(caves); i++ {
		for j := 0; j < len(caves[i]); j++ {
			if isLowPoint(caves, i, j) {
				sum += (1 + caves[i][j])
				//fmt.Println("Sum", sum, "- Cave value:", caves[i][j], " ", i, "x", j)
			}
		}
	}
	fmt.Println("What is the sum of the risk levels of all low points on your heightmap?")
	fmt.Println(sum)

	// -----------------

}

func isLowPoint(caves map[int]map[int]int, i int, j int) bool {
	pointValue := caves[i][j]
	value, ok := caves[i+1][j]
	if ok && value <= pointValue {
		return false
	}
	value, ok = caves[i-1][j]
	if ok && value <= pointValue {
		return false
	}
	value, ok = caves[i][j-1]
	if ok && value <= pointValue {
		return false
	}
	value, ok = caves[i][j+1]
	if ok && value <= pointValue {
		return false
	}

	return true
}
