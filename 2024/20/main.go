package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Position struct {
	Row, Col int
}

func main() {
	// Read input from file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	// Locate start ('S') and end ('E') positions
	var start, end Position
	for i, line := range grid {
		for j, char := range line {
			if char == 'S' {
				start = Position{i, j}
			} else if char == 'E' {
				end = Position{i, j}
			}
		}
	}

	// BFS-like traversal
	track := make(map[Position]int)
	track[start] = 0
	cur := start
	curStep := 0

	for cur != end {
		curStep++
		i, j := cur.Row, cur.Col
		moved := false

		for _, d := range []Position{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
			newPos := Position{i + d.Row, j + d.Col}
			if _, visited := track[newPos]; !visited {
				if newPos.Row >= 0 && newPos.Row < len(grid) &&
					newPos.Col >= 0 && newPos.Col < len(grid[0]) &&
					(grid[newPos.Row][newPos.Col] == 'S' || grid[newPos.Row][newPos.Col] == 'E' || grid[newPos.Row][newPos.Col] == '.') {
					cur = newPos
					track[cur] = curStep
					moved = true
					break
				}
			}
		}

		if !moved {
			panic("No valid move found, check input for correctness.")
		}
	}

	// Count special conditions
	count := 0
	for pos := range track {
		i, j := pos.Row, pos.Col
		for _, d := range []Position{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
			neighbor := Position{i + d.Row, j + d.Col}
			doubleNeighbor := Position{i + 2*d.Row, j + 2*d.Col}

			if _, visited := track[neighbor]; !visited {
				if step, exists := track[doubleNeighbor]; exists {
					if step-track[pos] >= 102 {
						count++
					}
				}
			}
		}
	}

	fmt.Println("Part 1: How many cheats would save you at least 100 picoseconds?")
	fmt.Println(count)

	// Part 2
	// Function to find cheat endpoints
	cheatEndpoints := func(coords Position) map[Position]struct{} {
		output := make(map[Position]struct{})
		i, j := coords.Row, coords.Col
		for di := -20; di <= 20; di++ {
			djMax := 20 - int(math.Abs(float64(di)))
			for dj := -djMax; dj <= djMax; dj++ {
				newPos := Position{i + di, j + dj}
				if _, exists := track[newPos]; exists {
					output[newPos] = struct{}{}
				}
			}
		}
		return output
	}

	// Function to calculate Manhattan distance
	manhattanDistance := func(coord1, coord2 Position) int {
		return int(math.Abs(float64(coord1.Row-coord2.Row)) + math.Abs(float64(coord1.Col-coord2.Col)))
	}

	// Count special conditions
	count = 0
	for coords := range track {
		potentials := cheatEndpoints(coords)
		for otherCoords := range potentials {
			if track[otherCoords]-track[coords]-manhattanDistance(coords, otherCoords) >= 100 {
				count++
			}
		}
	}

	fmt.Println()
	fmt.Println("Part 2: How many cheats would save you at least 100 picoseconds?")
	fmt.Println(count)
}
