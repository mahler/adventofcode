package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// flow simulates water flowing through the grid
func flow(grid [][]byte, x, y int, d int) int {
	if grid[y][x] == '.' {
		grid[y][x] = '|'
	}
	if y == len(grid)-1 {
		return x
	}
	if grid[y][x] == '#' {
		return x
	}
	if grid[y+1][x] == '.' {
		flow(grid, x, y+1, 0)
	}
	if grid[y+1][x] == '~' || grid[y+1][x] == '#' {
		if d != 0 {
			return flow(grid, x+d, y, d)
		} else {
			leftX := flow(grid, x-1, y, -1)
			rightX := flow(grid, x+1, y, 1)
			if grid[y][leftX] == '#' && grid[y][rightX] == '#' {
				for fillX := leftX + 1; fillX < rightX; fillX++ {
					grid[y][fillX] = '~'
				}
			}
		}
	}
	return x
}

// Coordinate represents a point in the grid
type Coordinate struct {
	x1, x2, y1, y2 int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data []Coordinate
	re := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)

	// Parse input
	minX, maxX := 500, 500
	minY, maxY := 100000, -1

	for scanner.Scan() {
		line := scanner.Text()
		nums := re.FindAllString(line, -1)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		c, _ := strconv.Atoi(nums[2])

		var coord Coordinate
		if line[0] == 'x' {
			coord = Coordinate{a, a, b, c}
		} else {
			coord = Coordinate{b, c, a, a}
		}

		// Track boundaries
		if coord.x1 < minX {
			minX = coord.x1
		}
		if coord.x2 > maxX {
			maxX = coord.x2
		}
		if coord.y1 < minY {
			minY = coord.y1
		}
		if coord.y2 > maxY {
			maxY = coord.y2
		}

		data = append(data, coord)
	}

	// Create grid
	grid := make([][]byte, maxY+1)
	width := maxX - minX + 3
	for i := range grid {
		grid[i] = make([]byte, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// Fill walls
	for _, coord := range data {
		for x := coord.x1; x <= coord.x2; x++ {
			for y := coord.y1; y <= coord.y2; y++ {
				grid[y][x-minX+1] = '#'
			}
		}
	}

	// Set spring
	springX := 500 - minX + 1
	grid[0][springX] = '+'

	// Simulate flow
	flow(grid, springX, 0, 0)

	// Count water
	still, flowing := 0, 0
	for y := minY; y <= maxY; y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '|' {
				flowing++
			} else if grid[y][x] == '~' {
				still++
			}
		}
	}

	fmt.Println("Part 1: How many tiles can the water reach within the range of y values in your scan?")
	fmt.Println(still + flowing)

	fmt.Println()
	fmt.Println("Part 2: How many water tiles are left after the water spring stops producing water and all remaining water not at rest has drained?")
	fmt.Println(still)
}
