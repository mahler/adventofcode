package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	total := 0
	for _, rucksack := range fileLines {
		set := make(map[rune]bool)
		l := len(rucksack)
		first, last := rucksack[:l/2], rucksack[l/2:]
		for _, char := range first {
			set[char] = true
		}
		for _, char := range last {
			if _, ok := set[char]; ok {
				total += getPrio(char)
				break
			}
		}
	}

	fmt.Println()
	fmt.Println("Day 3: Rucksack Reorganization")
	fmt.Println("What is the sum of the priorities of those item types?")
	fmt.Println(total)

	// Part 2
	total = 0

	for i := 0; i < len(fileLines); i += 3 {
		a, b, c := fileLines[i], fileLines[i+1], fileLines[i+2]
		seenA := make(map[rune]bool)
		seenB := make(map[rune]bool)
		for _, char := range a {
			seenA[char] = true
		}
		for _, char := range b {
			seenB[char] = true
		}
		for _, char := range c {
			if v1, v2 := seenA[char], seenB[char]; v1 && v2 {
				total += getPrio(char)
				break
			}
		}
	}

	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("What is the sum of the priorities of those item types?")
	fmt.Println(total)
}

func getPrio(common rune) int {
	if common >= 97 {
		return int(common) - 96
	} else if common >= 65 {
		return int(common) - 38
	}
	return 0
}
