package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// findAndReplaceStart finds the marker in the grid and replaces it with '.'
func findAndReplaceStart(grid [][]rune, marker rune) (int, int) {
	for row, line := range grid {
		for col, cell := range line {
			if cell == marker {
				grid[row][col] = '.'
				return row, col
			}
		}
	}
	return -1, -1
}

// readInput reads the grid from a file and finds the start position
func readInput(filename string) ([][]rune, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, -1, -1, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, -1, -1, err
	}

	sRow, sCol := findAndReplaceStart(grid, 'S')
	return grid, sRow, sCol, nil
}

// simulateBeams simulates beam propagation through the grid
func simulateBeams(filename string, countSplits bool) (int, error) {
	grid, sRow, sCol, err := readInput(filename)
	if err != nil {
		return 0, err
	}

	rows := len(grid)
	cols := len(grid[0])

	// Track beams in each column for the current row
	currentBeams := make([]int, cols)
	currentBeams[sCol] = 1
	splitsEncountered := 0

	// Process each row starting from the starting row
	for row := sRow; row < rows; row++ {
		nextBeams := make([]int, cols)

		for col := 0; col < cols; col++ {
			if currentBeams[col] > 0 {
				if grid[row][col] == '^' {
					// This is a split
					splitsEncountered++

					// Beams go left and right
					if col-1 >= 0 {
						nextBeams[col-1] += currentBeams[col]
					}
					if col+1 < cols {
						nextBeams[col+1] += currentBeams[col]
					}
				} else {
					// No split, beam continues straight down
					nextBeams[col] += currentBeams[col]
				}
			}
		}

		currentBeams = nextBeams
	}

	if countSplits {
		return splitsEncountered, nil
	}

	// Sum all remaining beams
	total := 0
	for _, count := range currentBeams {
		total += count
	}
	return total, nil
}

// solvePart1 counts the number of splits encountered
func solvePart1(filename string) (int, error) {
	return simulateBeams(filename, true)
}

// solvePart2 counts the total beams at the end
func solvePart2(filename string) (int, error) {
	return simulateBeams(filename, false)
}

func main() {
	part1, err := solvePart1("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Analyze your manifold diagram. How many times will the beam be split?")
	fmt.Println(part1)

	part2, err := solvePart2("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println("In total, how many different timelines would a single tachyon particle end up on?")
	fmt.Println(part2)
}
