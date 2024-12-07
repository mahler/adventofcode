package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type lineWindow struct {
	above *string
	cur   string
	below *string
}

func windows(input string) []lineWindow {
	lines := strings.Split(input, "\n")
	var windows []lineWindow

	for i := range lines {
		var above, below *string
		if i > 0 {
			above = &lines[i-1]
		}
		if i < len(lines)-1 {
			below = &lines[i+1]
		}
		windows = append(windows, lineWindow{
			above: above,
			cur:   lines[i],
			below: below,
		})
	}

	return windows
}

func isSymbol(ch rune) bool {
	return !unicode.IsDigit(ch) && ch != '.'
}

func hasAdjSymbol(w lineWindow, m []int) bool {
	start, end := m[0], m[1]
	start = max(1, start) - 1
	end = min(len(w.cur), end+1)

	// Check current line boundaries
	if isSymbol(rune(w.cur[start])) || isSymbol(rune(w.cur[end-1])) {
		return true
	}

	// Check above and below lines
	lines := []string{}
	if w.above != nil {
		lines = append(lines, *w.above)
	}
	if w.below != nil {
		lines = append(lines, *w.below)
	}

	for _, line := range lines {
		for _, ch := range line[start:end] {
			if isSymbol(ch) {
				return true
			}
		}
	}

	return false
}

func part1(input string) int {
	numRE := regexp.MustCompile(`\d+`)
	total := 0

	for _, w := range windows(input) {
		nums := numRE.FindAllStringIndex(w.cur, -1)
		for _, m := range nums {
			if hasAdjSymbol(w, m) {
				num, _ := strconv.Atoi(w.cur[m[0]:m[1]])
				total += num
			}
		}
	}

	return total
}

func part2(input string) int {
	numRE := regexp.MustCompile(`\d+`)
	total := 0

	for _, w := range windows(input) {
		for i, ch := range w.cur {
			if ch == '*' {
				ratio := findGearRatio(w, i, numRE)
				total += ratio
			}
		}
	}

	return total
}

func findGearRatio(w lineWindow, gearIndex int, numRE *regexp.Regexp) int {
	var adjacentNums []int
	lines := []string{w.cur}
	if w.above != nil {
		lines = append(lines, *w.above)
	}
	if w.below != nil {
		lines = append(lines, *w.below)
	}

	for _, line := range lines {
		matches := numRE.FindAllStringIndex(line, -1)
		for _, m := range matches {
			if isAdjacent(gearIndex, m[0], m[1]) {
				num, _ := strconv.Atoi(line[m[0]:m[1]])
				adjacentNums = append(adjacentNums, num)
			}
		}
	}

	if len(adjacentNums) == 2 {
		return adjacentNums[0] * adjacentNums[1]
	}
	return 0
}

func isAdjacent(gearIndex, numStart, numEnd int) bool {
	return gearIndex >= numStart-1 && gearIndex < numEnd+1
}

func main() {
	// Read input file
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	input := strings.TrimSpace(string(inputBytes))

	// Solve and print
	part1Result := part1(input)
	part2Result := part2(input)
	fmt.Printf("Day 03 Part 1: %d\n", part1Result)
	fmt.Printf("Day 03 Part 2: %d\n", part2Result)
}
