package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Position represents a 2D coordinate
type Position struct {
	x, y int
}

// Queue for BFS
type QueueItem struct {
	pos  Position
	dist int
}

func readInput(filename string) map[Position]rune {
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	file, err := os.Open(filepath.Join(currentDir, filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make(map[Position]rune)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			grid[Position{x, y}] = char
		}
		y++
	}
	return grid
}

func getStart(grid map[Position]rune) Position {
	for pos, tile := range grid {
		if tile == 'S' {
			return pos
		}
	}
	panic("Start position not found")
}

func getGridSize(grid map[Position]rune) int {
	maxX := 0
	for pos := range grid {
		if pos.x > maxX {
			maxX = pos.x
		}
	}
	return maxX + 1
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func walk(maxDist int, start Position, grid map[Position]rune) map[int]int {
	tiles := make(map[int]int)
	seen := make(map[Position]bool)
	queue := []QueueItem{{pos: start, dist: 0}}
	size := getGridSize(grid)

	// Neighbors offsets (up, down, left, right)
	neighbors := []Position{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.dist == maxDist+1 || seen[current.pos] {
			continue
		}

		tiles[current.dist]++
		seen[current.pos] = true

		for _, offset := range neighbors {
			newPos := Position{
				x: current.pos.x + offset.x,
				y: current.pos.y + offset.y,
			}

			// Get wrapped position for grid lookup
			wrappedPos := Position{
				x: mod(newPos.x, size),
				y: mod(newPos.y, size),
			}

			if grid[wrappedPos] != '#' {
				queue = append(queue, QueueItem{pos: newPos, dist: current.dist + 1})
			}
		}
	}

	return tiles
}

func getGardenTiles(steps int, grid map[Position]rune) int {
	tiles := walk(steps, getStart(grid), grid)
	sum := 0
	for distance, amount := range tiles {
		if distance%2 == steps%2 {
			sum += amount
		}
	}
	return sum
}

func main() {
	grid := readInput("input.txt")
	steps := 64
	fmt.Println("Part 1: Starting from the garden plot marked S on your map, how many garden plots could the Elf reach in exactly 64 steps?")
	tiles := walk(steps, getStart(grid), grid)
	sum := 0
	for distance, amount := range tiles {
		if distance%2 == steps%2 {
			sum += amount
		}
	}

	fmt.Println(sum)

	// Part 2
	size := getGridSize(grid)
	edge := size / 2

	// Calculate sequence
	y := make([]int, 3)
	for i := 0; i < 3; i++ {
		y[i] = getGardenTiles(edge+i*size, grid)
	}

	// Calculate quadratic coefficients
	a := (y[2] - 2*y[1] + y[0]) / 2
	b := y[1] - y[0] - a
	c := y[0]

	// Calculate result for target
	target := (26501365 - edge) / size
	result := a*target*target + b*target + c

	fmt.Println()
	fmt.Println("Part 2: However, the step count the Elf needs is much larger!")
	fmt.Println("Starting from the garden plot marked S on your infinite map,")
	fmt.Println("how many garden plots could the Elf reach in exactly 26501365 steps?")
	fmt.Println(result)
}
