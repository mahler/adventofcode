package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the input file
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the file content into lines
	lines := strings.Split(string(data), "\n")

	// Remove any empty lines at the end
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// Parse the grid
	H, W := len(lines), len(lines[0])
	grid := make(map[string]rune)
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			grid[fmt.Sprintf("%d,%d", y, x)] = rune(lines[y][x])
		}
	}

	// Part 1 - Find anything that says 'XMAS'
	target := "XMAS"
	deltas := generateDeltas()
	count := 0

	for coord := range grid {
		for _, delta := range deltas {
			candidate := findCandidate(grid, coord, delta, len(target))
			if candidate == target {
				count++
			}
		}
	}
	fmt.Println("Part 1:How many times does XMAS appear?")
	fmt.Println(count)

}

// Generate all possible deltas excluding (0,0)
func generateDeltas() [][]int {
	var deltas [][]int
	for _, dy := range []int{-1, 0, 1} {
		for _, dx := range []int{-1, 0, 1} {
			if dx != 0 || dy != 0 {
				deltas = append(deltas, []int{dy, dx})
			}
		}
	}
	return deltas
}

// Find a candidate string based on given coordinate, delta, and length
func findCandidate(grid map[string]rune, coord string, delta []int, length int) string {
	coords := strings.Split(coord, ",")
	y, x := parseInt(coords[0]), parseInt(coords[1])

	var candidate strings.Builder
	for i := 0; i < length; i++ {
		newCoord := fmt.Sprintf("%d,%d", y+delta[0]*i, x+delta[1]*i)
		if char, exists := grid[newCoord]; exists {
			candidate.WriteRune(char)
		}
	}
	return candidate.String()
}

// Get adjacent characters for Part 2
func getAdjacentChars(grid map[string]rune, coord string, dyRange, dxRange []int) string {
	coords := strings.Split(coord, ",")
	y, x := parseInt(coords[0]), parseInt(coords[1])

	var chars strings.Builder
	for _, dy := range dyRange {
		for _, dx := range dxRange {
			newCoord := fmt.Sprintf("%d,%d", y+dy, x+dx)
			if char, exists := grid[newCoord]; exists {
				chars.WriteRune(char)
			}
		}
	}
	return chars.String()
}

// Helper function to parse integer from string
func parseInt(s string) int {
	val := 0
	for _, r := range s {
		val = val*10 + int(r-'0')
	}
	return val
}
