package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	// Setup
	fileRows := string(fileContent)
	fileLines := strings.Split(fileRows, "\n")

	sum := 0
	for _, line := range fileLines {
		lineParts := strings.Split(line, "|")
		outputValueParts := strings.Split(lineParts[1], " ")

		for _, value := range outputValueParts {
			if len(value) == 2 || len(value) == 3 || len(value) == 4 || len(value) == 7 {
				sum += 1
			}
		}
	}
	fmt.Println("Part 1:")
	fmt.Println(sum)

}
