package main

// https://regoio.herokuapp.com/
// (\w+) (\d+),(\d+) through (\d+),(\d+)

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Read instructions
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	fmt.Println("Instructions in dataset:", len(fileLines))

	fmt.Println()
	fmt.Println("Day 06, part 1: Probably a Fire Hazard")

	// Setup lightgrid
	lightGrid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		lightGrid[i] = make([]int, 1000)
	}

	regexpData := regexp.MustCompile(`(\w+) (\d+),(\d+) through (\d+),(\d+)`)

	for _, fileLine := range fileLines {
		if strings.Contains(fileLine, "turn ") {
			fileLine = fileLine[5:]
		}

		fields := regexpData.FindStringSubmatch(fileLine)

		// fmt.Println(fields)
		instruction := fields[1]
		rowStart, _ := strconv.Atoi(fields[2])
		rowEnd, _ := strconv.Atoi(fields[4])
		colStart, _ := strconv.Atoi(fields[3])
		colEnd, _ := strconv.Atoi(fields[5])

		for col := colStart; col <= colEnd; col++ {
			for row := rowStart; row <= rowEnd; row++ {
				if instruction == "on" {
					lightGrid[col][row] = 1
				} else if instruction == "off" {
					lightGrid[col][row] = 0
				} else if instruction == "toggle" {
					if lightGrid[col][row] == 1 {
						lightGrid[col][row] = 0
					} else {
						lightGrid[col][row] = 1
					}
				} else {
					fmt.Println("ERRRRRRRRRROR!", instruction)
				}

			}
		}
		fmt.Println("Lights on:", countLightsInGrid(lightGrid))
	}

	fmt.Println("how many lights are lit?", countLightsInGrid(lightGrid))

	// ------------ PART 2 ------------------------
	fmt.Println()
}

func countLightsInGrid(lights [][]int) int {
	counter := 0

	for column := 0; column < len(lights); column++ {
		for row := 0; row < len(lights[0]); row++ {
			if lights[column][row] == 1 {
				counter++
			}
		}
	}

	return counter
}
