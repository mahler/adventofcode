package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read instructions
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	fmt.Println()
	fmt.Println("Day 02, part 1: Dive!")

	regexpData := regexp.MustCompile(`(\w+) (\d+)`)

	depth, distance := 0, 0
	for _, fileLine := range fileLines {
		fields := regexpData.FindStringSubmatch(fileLine)

		// fmt.Println(fields)

		direction := fields[1]
		length, _ := strconv.Atoi(fields[2])
		//fmt.Println(direction, "-", length)

		switch direction {
		case "up":
			depth -= length
			break
		case "down":
			depth += length
		case "forward":
			distance += length
			break
		}
	}
	fmt.Println("Distance:", distance)
	fmt.Println("Depth:", depth)
	fmt.Println("Answer:", distance*depth)
}
