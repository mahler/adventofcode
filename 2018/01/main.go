package main

import (
	"fmt"
	"log"
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
	fmt.Println("DAY01, Part 1: Chronal Calibration")
	fmt.Println("Total instructions in dataset: ", len(fileContent))

	// Starting frequency
	frequency := 0

	for _, calibration := range fileContent {
		cali, _ := strconv.Atoi(calibration)
		frequency += cali
	}

	fmt.Println("End frequency:", frequency)

	// -------------------------------------------------------------
	fmt.Println()
	fmt.Println("Part 2")

	// Reset starting frequency
	frequency = 0
	seenFrq := make(map[int]bool)
	found := false

	// Make sure we loop over calibration list
	for !found {
		for _, calibration := range fileContent {
			cali, _ := strconv.Atoi(calibration)
			frequency += cali

			if _, ok := seenFrq[frequency]; ok {
				fmt.Println("First Duplicate Frequency:", frequency)
				found = true
				break
			} else {
				seenFrq[frequency] = true
			}

		}
	}
}
