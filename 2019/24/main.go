package main

import (
	"fmt"
	"os"
	"strings"
)

// Utility function to count adjacent bugs
func countNeighbours(grid [][]rune, x, y int) int {
	count := 0

	// Check left
	if x > 0 && grid[y][x-1] == '#' {
		count++
	}

	// Check up
	if y > 0 && grid[y-1][x] == '#' {
		count++
	}

	// Check right
	if x < len(grid[0])-1 && grid[y][x+1] == '#' {
		count++
	}

	// Check down
	if y < len(grid)-1 && grid[y+1][x] == '#' {
		count++
	}

	return count
}

// Determine the new state of a cell
func stepCoord(grid [][]rune, x, y int) rune {
	if grid[y][x] == '#' {
		if countNeighbours(grid, x, y) == 1 {
			return '#'
		}
		return '.'
	} else {
		neighbours := countNeighbours(grid, x, y)
		if neighbours == 1 || neighbours == 2 {
			return '#'
		}
		return '.'
	}
}

// Advance the entire grid one step
func stepGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for y := range grid {
		newGrid[y] = make([]rune, len(grid[y]))
		for x := range grid[y] {
			newGrid[y][x] = stepCoord(grid, x, y)
		}
	}
	return newGrid
}

// Calculate biodiversity rating
func biodiversity(grid [][]rune) int {
	score := 0
	power := 1

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '#' {
				score += power
			}
			power *= 2
		}
	}

	return score
}

// Hash the grid for comparison
func gridHash(grid [][]rune) string {
	var b strings.Builder
	for y := range grid {
		for x := range grid[y] {
			b.WriteRune(grid[y][x])
		}
	}
	return b.String()
}

// Define a 3D coordinate
type Coord struct {
	X, Y, Z int
}

// Get adjacent cells for part 2
func adjacentP2(x, y, z int) []Coord {
	// Skip the center cell
	if x == 2 && y == 2 {
		return []Coord{}
	}

	adjacent := []Coord{}

	// Look left
	if x == 0 {
		// Looking out
		adjacent = append(adjacent, Coord{1, 2, z - 1})
	} else if x == 3 && y == 2 {
		// Looking in
		for yy := 0; yy < 5; yy++ {
			adjacent = append(adjacent, Coord{4, yy, z + 1})
		}
	} else {
		adjacent = append(adjacent, Coord{x - 1, y, z})
	}

	// Look right
	if x == 4 {
		// Looking out
		adjacent = append(adjacent, Coord{3, 2, z - 1})
	} else if x == 1 && y == 2 {
		// Looking in
		for yy := 0; yy < 5; yy++ {
			adjacent = append(adjacent, Coord{0, yy, z + 1})
		}
	} else {
		adjacent = append(adjacent, Coord{x + 1, y, z})
	}

	// Look up
	if y == 0 {
		// Looking out
		adjacent = append(adjacent, Coord{2, 1, z - 1})
	} else if x == 2 && y == 3 {
		// Looking in
		for xx := 0; xx < 5; xx++ {
			adjacent = append(adjacent, Coord{xx, 4, z + 1})
		}
	} else {
		adjacent = append(adjacent, Coord{x, y - 1, z})
	}

	// Look down
	if y == 4 {
		// Looking out
		adjacent = append(adjacent, Coord{2, 3, z - 1})
	} else if x == 2 && y == 1 {
		// Looking in
		for xx := 0; xx < 5; xx++ {
			adjacent = append(adjacent, Coord{xx, 0, z + 1})
		}
	} else {
		adjacent = append(adjacent, Coord{x, y + 1, z})
	}

	return adjacent
}

// Count neighbors for part 2
func countNeighboursP2(grids map[int][][]rune, x, y, z int) int {
	adjacent := adjacentP2(x, y, z)
	ch := grids[z][y][x]
	nNeighbours := 0

	for _, coord := range adjacent {
		// Skip creating new levels if not adjacent to bugs
		if _, exists := grids[coord.Z]; !exists && ch != '#' {
			continue
		}

		// Ensure the level exists in our map
		if _, exists := grids[coord.Z]; !exists {
			grids[coord.Z] = emptyGrid()
		}

		if grids[coord.Z][coord.Y][coord.X] == '#' {
			nNeighbours++
		}
	}

	return nNeighbours
}

// Determine new state for a cell in part 2
func stepCoordP2(grids map[int][][]rune, x, y, z int) rune {
	if grids[z][y][x] == '#' {
		if countNeighboursP2(grids, x, y, z) == 1 {
			return '#'
		}
		return '.'
	} else {
		neighbours := countNeighboursP2(grids, x, y, z)
		if neighbours == 1 || neighbours == 2 {
			return '#'
		}
		return '.'
	}
}

// Advance an entire level in part 2
func stepGridP2(grids map[int][][]rune, z int) [][]rune {
	newGrid := make([][]rune, 5)
	for y := range grids[z] {
		newGrid[y] = make([]rune, 5)
		for x := range grids[z][y] {
			newGrid[y][x] = stepCoordP2(grids, x, y, z)
		}
	}
	return newGrid
}

// Create an empty 5x5 grid
func emptyGrid() [][]rune {
	grid := make([][]rune, 5)
	for i := range grid {
		grid[i] = []rune(".....")
	}
	return grid
}

// Iterate all levels for part 2
func iterate(grids map[int][][]rune) map[int][][]rune {
	newGrids := make(map[int][][]rune)
	gridsToScan := make(map[int]bool)

	// Mark existing levels to scan
	for z := range grids {
		gridsToScan[z] = true
	}

	// Step existing levels
	for z := range gridsToScan {
		newGrids[z] = stepGridP2(grids, z)
	}

	// Step any new levels that were implicitly created
	for z := range grids {
		if !gridsToScan[z] {
			newGrids[z] = stepGridP2(grids, z)
		}
	}

	return newGrids
}

// Count total bugs across all levels
func countBugs(grids map[int][][]rune) int {
	count := 0
	for _, grid := range grids {
		for _, row := range grid {
			for _, cell := range row {
				if cell == '#' {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	mapStr := string(data)

	biodiversityRating := 0
	lines := strings.Split(strings.TrimSpace(mapStr), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	seen := make(map[string]bool)
	for {
		gridKey := gridHash(grid)
		if seen[gridKey] {
			biodiversityRating = biodiversity(grid)
			break
		}
		seen[gridKey] = true
		grid = stepGrid(grid)
	}

	fmt.Println("Part 1: What is the biodiversity rating for the first layout that appears twice?")
	fmt.Println(biodiversityRating)

	// Reset grid
	grid = make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	grids := make(map[int][][]rune)
	grids[0] = grid

	for s := 0; s < 200; s++ {
		grids = iterate(grids)
	}

	fmt.Println()
	fmt.Println("Part 2: Starting with your scan, how many bugs are present after 200 minutes?")
	fmt.Println(countBugs(grids))
}
