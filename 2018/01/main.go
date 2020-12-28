package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("DAY01, Part 1: Chronal Calibration")
	fmt.Println("Total passwords in dataset: ", len(fileContent))

	// Starting frequency
	frequency := 0

	for _, calibration := range fileContent {
		cali, _ := strconv.Atoi(calibration)

		frequency += cali
	}

	fmt.Println("End frequency:", frequency)

	fmt.Println()
	fmt.Println("Part 2")

}
