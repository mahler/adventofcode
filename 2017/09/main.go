package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input file
	content, _ := os.ReadFile("input.txt")
	line := strings.TrimSpace(string(content))

	score := 0
	garbageScore := 0
	currentDepth := 0
	insideGarbage := false
	skipChar := false

	for _, char := range line {
		if insideGarbage {
			if skipChar {
				skipChar = false
			} else if char == '!' {
				skipChar = true
			} else if char == '>' {
				insideGarbage = false
			} else {
				garbageScore++
			}
		} else { // when inside group, not garbage
			if char == '{' {
				currentDepth++
			} else if char == '}' {
				score += currentDepth
				currentDepth--
			} else if char == '<' {
				insideGarbage = true
			}
		}
	}

	fmt.Println("Part 1: What is the total score for all groups in your input?")
	fmt.Println(score)

	fmt.Println()
	fmt.Println("Part 2: How many non-canceled characters are within the garbage in your puzzle input?")
	fmt.Println(garbageScore)
}
