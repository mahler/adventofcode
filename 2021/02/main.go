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
		direction := fields[1]
		length, _ := strconv.Atoi(fields[2])

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

	// Part 2
	fmt.Println()
	fmt.Println("Part 2")

	// reset start pos
	depth, distance = 0, 0
	aim := 0
	for _, fileLine := range fileLines {
		fields := regexpData.FindStringSubmatch(fileLine)

		direction := fields[1]
		length, _ := strconv.Atoi(fields[2])

		switch direction {
		case "up":
			aim -= length
			break
		case "down":
			aim += length
		case "forward":
			distance += length
			depth += (length * aim)
			break
		}
	}
	fmt.Println("Distance:", distance)
	fmt.Println("Depth:", depth)
	fmt.Println("Answer:", distance*depth)
}
