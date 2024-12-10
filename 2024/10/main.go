package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readInput reads the input from a file and converts it to a 2D slice of integers
func readInput(filename string) [][]int {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	mat := make([][]int, len(lines))

	for i, line := range lines {
		if line == "" {
			continue
		}
		row := make([]int, len(line))
		for j, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			row[j] = num
		}
		mat[i] = row
	}

	return mat
}

// trailheadLocations finds all starting points with value 0
func trailheadLocations(mat [][]int) [][2]int {
	var trailheads [][2]int
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[0]); y++ {
			if mat[x][y] == 0 {
				trailheads = append(trailheads, [2]int{x, y})
			}
		}
	}
	return trailheads
}

// findAllPaths finds paths from a starting point
func findAllPaths(mat [][]int, start [2]int, byRating bool) int {
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make(map[string]bool)
	rows, cols := len(mat), len(mat[0])

	var iterate func(int, int, string) int
	iterate = func(x, y int, path string) int {
		// Check bounds
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return 0
		}

		// Create cache key
		var cacheKey string
		if byRating {
			cacheKey = fmt.Sprintf("%d,%d,%s", x, y, path)
		} else {
			cacheKey = fmt.Sprintf("%d,%d", x, y)
		}

		// Check visited
		if visited[cacheKey] {
			return 0
		}
		visited[cacheKey] = true

		// Check destination
		if mat[x][y] == 9 {
			return 1
		}

		pathCount := 0
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && mat[nx][ny]-mat[x][y] == 1 {
				newPath := path + fmt.Sprintf("(%d,%d)", nx, ny)
				pathCount += iterate(nx, ny, newPath)
			}
		}

		return pathCount
	}

	return iterate(start[0], start[1], "")
}

func main() {
	// Replace with your input file path
	mat := readInput("input.txt")
	trailheads := trailheadLocations(mat)

	part1 := 0
	part2 := 0

	for _, th := range trailheads {
		part1 += findAllPaths(mat, th, false)
		part2 += findAllPaths(mat, th, true)
	}

	fmt.Println("Part 1: The reindeer gleefully carries over a protractor and adds it to the pile.")
	fmt.Println("What is the sum of the scores of all trailheads on your topographic map?")
	fmt.Println(part1)
	fmt.Println()
	fmt.Println("Part 2: You're not sure how, but the reindeer seems to have crafted some tiny flags")
	fmt.Println("out of toothpicks and bits of paper and is using them to mark trailheads on your topographic map.")
	fmt.Println("What is the sum of the ratings of all trailheads?")
	fmt.Println(part2)
}
