package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println()
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	// Setup
	strTemplate := ""
	mapPolymer := make(map[string]string)

	fileRows := string(fileContent)
	fileLines := strings.Split(strings.TrimSpace(fileRows), "\n")

	for _, line := range fileLines {
		// empty line
		if len(line) <= 0 {
			continue
		}

		// pair insertion rule
		if strings.Contains(line, "->") {
			strTmp := strings.Split(line, " -> ")
			key := strTmp[0]
			val := strTmp[1]

			mapPolymer[key] = val
		} else {
			// polymer template
			strTemplate = line
		}
	}

	fmt.Println("Day 2021-14/")
	fmt.Println("Part 1 - 10 steps")
	fmt.Println("What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?")

	part1answer := polymerSteps(strTemplate, mapPolymer, 10)
	fmt.Println(part1answer)

	fmt.Println()
	fmt.Println("Part 2 - 40 Steps")
	fmt.Println("What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?")

	part2answer := polymerSteps(strTemplate, mapPolymer, 40)
	fmt.Println(part2answer)

}

func polymerSteps(strTemplate string, mapPolymer map[string]string, steps int) int {
	charMap := make(map[string]int)

	for i := 0; i < len(strTemplate)-1; i++ {
		charMap[strTemplate[i:i+2]] += 1
	}
	charMap[strTemplate[len(strTemplate)-1:]] += 1

	for i := 0; i < steps; i++ {
		newCharMap := make(map[string]int)

		for k, v := range charMap {
			if mapPolymer[k] != "" {
				newCharMap[k[0:1]+mapPolymer[k]] += v
				newCharMap[mapPolymer[k]+k[1:2]] += v
			} else {
				newCharMap[k] += v
			}
		}
		charMap = newCharMap
	}

	freq := make(map[string]int)
	for k, v := range charMap {
		freq[k[0:1]] += v
	}

	minVal := math.MaxInt64
	maxVal := 0
	for _, val := range freq {
		if val < minVal {
			minVal = val
		}

		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal - minVal
}
