package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(grid [][]rune, reps int, solveMode string) int {
	var infect rune
	flip := make(map[rune]rune)

	// Determine infection rules based on solveMode
	if solveMode == "part1" {
		infect = '.'
		flip['.'] = '#'
		flip['#'] = '.'
	} else {
		infect = 'W'
		flip['.'] = 'W'
		flip['W'] = '#'
		flip['#'] = 'F'
		flip['F'] = '.'
	}

	// Initialize variables
	x, y := len(grid)/2, len(grid[0])/2
	mov := [2]int{-1, 0}
	spin := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	infections := 0

	// Main loop for iterations
	for i := 0; i < reps; i++ {
		state := grid[x][y]

		// Check for infection
		if state == infect {
			infections++
		}

		// Flip state
		grid[x][y] = flip[state]

		// Adjust direction based on state
		if state == '.' || state == '#' {
			rot := 1
			if state == '.' {
				rot = -1
			}
			newIndex := (indexOf(spin, mov) + rot + 4) % 4
			mov = spin[newIndex]
		} else if state == 'F' {
			mov[0] = -mov[0]
			mov[1] = -mov[1]
		}

		// Move position
		x += mov[0]
		y += mov[1]

		// Dynamically expand grid if necessary
		if x < 0 {
			x++
			grid = append([][]rune{make([]rune, len(grid[0]))}, grid...)
			for i := 0; i < len(grid[0]); i++ {
				grid[0][i] = '.'
			}
		} else if x >= len(grid) {
			grid = append(grid, make([]rune, len(grid[0])))
			for i := 0; i < len(grid[0]); i++ {
				grid[len(grid)-1][i] = '.'
			}
		}
		if y < 0 {
			y++
			for i := range grid {
				grid[i] = append([]rune{'.'}, grid[i]...)
			}
		} else if y >= len(grid[0]) {
			for i := range grid {
				grid[i] = append(grid[i], '.')
			}
		}
	}

	return infections
}

// Helper function to find the index of a 2D vector in a slice
func indexOf(slice [][2]int, item [2]int) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Solve Part 1
	gridCopy := make([][]rune, len(grid))
	for i := range grid {
		gridCopy[i] = append([]rune{}, grid[i]...)
	}
	fmt.Println("Part 1: Given your actual map, after 10000 bursts of activity, how many bursts cause a node to become infected?")
	fmt.Println(solve(gridCopy, 10000, "part1"))

	// Solve Part 2
	fmt.Println()
	fmt.Println("Part 2: Given your actual map, after 10000000 bursts of activity, how many bursts cause a node to become infected?")
	fmt.Println(solve(grid, 10000000, "part2"))
}
