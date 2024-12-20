package main

import (
	"fmt"
	"os"
	"strings"
)

type Input struct {
	availableStripes map[string]struct{}
	targetDesigns    []string
}

func counts(input *Input) []int {
	known := map[string]int{"": 1}
	results := make([]int, 0, len(input.targetDesigns))

	for _, target := range input.targetDesigns {
		results = append(results, count(target, input.availableStripes, known))
	}

	return results
}

func count(target string, availableStripes map[string]struct{}, known map[string]int) int {
	if val, exists := known[target]; exists {
		return val
	}

	sum := 0
	for prefix := range availableStripes {
		if strings.HasPrefix(target, prefix) {
			sum += count(strings.TrimPrefix(target, prefix), availableStripes, known)
		}
	}

	known[target] = sum
	return sum
}

func filterPositiveCounts(counts []int) []int {
	filtered := []int{}
	for _, count := range counts {
		if count > 0 {
			filtered = append(filtered, count)
		}
	}
	return filtered
}

func main() {
	inputBytes, _ := os.ReadFile("input.txt")
	inputString := string(inputBytes)

	// func parseInput(input string) (Input, error) {
	parts := strings.SplitN(inputString, "\n\n", 2)

	availableStripes := make(map[string]struct{})
	for _, stripe := range strings.Split(parts[0], ",") {
		availableStripes[strings.TrimSpace(stripe)] = struct{}{}
	}
	targetDesigns := strings.Split(strings.TrimSpace(parts[1]), "\n")

	parsedInput := Input{
		availableStripes: availableStripes,
		targetDesigns:    targetDesigns,
	}

	counts := filterPositiveCounts(counts(&parsedInput))

	fmt.Println("Part 1: How many designs are possible?")
	fmt.Println(len(counts))

	// Part 2
	sum := 0
	for _, count := range counts {
		sum += count
	}
	fmt.Println()
	fmt.Println("Part 2: What do you get if you add up the number of different ways you could make each design?")
	fmt.Println(sum)
}
