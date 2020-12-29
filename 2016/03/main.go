package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 3: Squares With Three Sides")
	fileContent, err := ioutil.ReadFile("puzzle.txt")
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
}
