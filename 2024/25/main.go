package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire file
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	// Split into patterns
	rawPatterns := strings.Split(strings.TrimSpace(content), "\n\n")
	patterns := make([]map[[2]int]struct{}, len(rawPatterns))

	// Parse each pattern
	for pIndex, pattern := range rawPatterns {
		lines := strings.Split(pattern, "\n")
		patternSet := make(map[[2]int]struct{})
		for i, line := range lines {
			for j, char := range line {
				if char == '#' {
					patternSet[[2]int{i, j}] = struct{}{}
				}
			}
		}
		patterns[pIndex] = patternSet
	}

	// Count non-overlapping pairs
	count := 0
	for i := 0; i < len(patterns); i++ {
		for j := i + 1; j < len(patterns); j++ {
			if !hasIntersection(patterns[i], patterns[j]) {
				count++
			}
		}
	}

	// Print the result
	fmt.Println("Part 1: How many unique lock/key pairs fit together without overlapping in any column?")
	fmt.Println(count)
}

// Check if two sets intersect
func hasIntersection(set1, set2 map[[2]int]struct{}) bool {
	for key := range set1 {
		if _, exists := set2[key]; exists {
			return true
		}
	}
	return false
}
