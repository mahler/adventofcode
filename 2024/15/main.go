package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// Point represents a 2D coordinate
type Point struct {
	x, y int
}

// Grid represents a 2D grid of bytes
type Grid struct {
	data   []byte
	width  int
	height int
}

//go:embed input.txt
var input string

// NewGrid creates a new grid with given width, height, and default value
func NewGrid(width, height int, defaultVal byte) *Grid {
	g := &Grid{
		data:   make([]byte, width*height),
		width:  width,
		height: height,
	}
	for i := range g.data {
		g.data[i] = defaultVal
	}
	return g
}

// Get retrieves the value at a specific point
func (g *Grid) Get(p Point) byte {
	if p.x < 0 || p.x >= g.width || p.y < 0 || p.y >= g.height {
		panic("point out of bounds")
	}
	return g.data[p.y*g.width+p.x]
}

// Set sets the value at a specific point
func (g *Grid) Set(p Point, val byte) {
	if p.x < 0 || p.x >= g.width || p.y < 0 || p.y >= g.height {
		panic("point out of bounds")
	}
	g.data[p.y*g.width+p.x] = val
}

// Clone creates a deep copy of the grid
func (g *Grid) Clone() *Grid {
	newGrid := NewGrid(g.width, g.height, '.')
	copy(newGrid.data, g.data)
	return newGrid
}

// Find locates the first occurrence of a specific byte
func (g *Grid) Find(needle byte) *Point {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			p := Point{x, y}
			if g.Get(p) == needle {
				return &p
			}
		}
	}
	return nil
}

// Predefined directions
var (
	LEFT  = Point{-1, 0}
	RIGHT = Point{1, 0}
	UP    = Point{0, -1}
	DOWN  = Point{0, 1}
	ORIG  = Point{0, 0}
)

// Add adds two points
func (p Point) Add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

// Parse parses the input into a grid and moves
func Parse(input string) (*Grid, string) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		panic("invalid input format")
	}

	// Parse grid
	lines := strings.Split(parts[0], "\n")
	grid := NewGrid(len(lines[0]), len(lines), '.')
	for y, line := range lines {
		for x, ch := range line {
			grid.Set(Point{x, y}, byte(ch))
		}
	}

	return grid, parts[1]
}

// Part1 solves the first part of the puzzle
func Part1(grid *Grid, moves string) int {
	gridCopy := grid.Clone()
	position := gridCopy.Find(byte('@'))
	if position == nil {
		panic("no start position found")
	}
	gridCopy.Set(*position, '.')

	for _, b := range moves {
		switch b {
		case '<':
			narrow(gridCopy, position, LEFT)
		case '>':
			narrow(gridCopy, position, RIGHT)
		case '^':
			narrow(gridCopy, position, UP)
		case 'v':
			narrow(gridCopy, position, DOWN)
		}
	}

	return gps(gridCopy, byte('O'))
}

// Part2 solves the second part of the puzzle
func Part2(grid *Grid, moves string) int {
	stretchedGrid := stretch(grid)
	position := stretchedGrid.Find(byte('@'))
	if position == nil {
		panic("no start position found")
	}
	stretchedGrid.Set(*position, '.')

	todo := make([]Point, 0, 50)

	for _, b := range moves {
		switch b {
		case '<':
			narrow(stretchedGrid, position, LEFT)
		case '>':
			narrow(stretchedGrid, position, RIGHT)
		case '^':
			wide(stretchedGrid, position, UP, &todo)
		case 'v':
			wide(stretchedGrid, position, DOWN, &todo)
		}
	}

	return gps(stretchedGrid, byte('['))
}

// narrow handles pushing in a single direction with constraints
func narrow(grid *Grid, start *Point, direction Point) {
	position := start.Add(direction)
	size := 1

	// Search for next wall or open space
	for grid.Get(position) != '.' && grid.Get(position) != '#' {
		position = position.Add(direction)
		size++
	}

	// Move items if open space found
	if grid.Get(position) == '.' {
		previous := byte('.')
		position = start.Add(direction)

		for i := 0; i < size; i++ {
			// Swap current item with previous
			current := grid.Get(position)
			grid.Set(position, previous)
			previous = current
			position = position.Add(direction)
		}

		// Move robot
		*start = start.Add(direction)
	}
}

// wide handles pushing with more complex rules
func wide(grid *Grid, start *Point, direction Point, todo *[]Point) {
	// Short circuit if path is empty
	if grid.Get(start.Add(direction)) == '.' {
		*start = start.Add(direction)
		return
	}

	// Clear todo list
	*todo = (*todo)[:0]
	// Add dummy item to prevent out of bounds
	*todo = append(*todo, ORIG, *start)
	index := 1

	for index < len(*todo) {
		next := (*todo)[index].Add(direction)
		index++

		// Add boxes strictly left to right
		var first, second Point
		switch grid.Get(next) {
		case '[':
			first, second = next, next.Add(RIGHT)
		case ']':
			first, second = next.Add(LEFT), next
		case '#':
			return // Wall in the way, cancel move
		default:
			continue // Open space, skip
		}

		// Check if box has already been added
		if first != (*todo)[len(*todo)-2] {
			*todo = append(*todo, first, second)
		}
	}

	// Move boxes in reverse order
	for i := len(*todo) - 1; i >= 2; i-- {
		point := (*todo)[i]
		grid.Set(point.Add(direction), grid.Get(point))
		grid.Set(point, '.')
	}

	// Move robot
	*start = start.Add(direction)
}

// stretch doubles the grid width
func stretch(grid *Grid) *Grid {
	next := NewGrid(grid.width*2, grid.height, '.')

	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			// Determine left and right values based on original grid
			var left, right byte
			switch grid.Get(Point{x, y}) {
			case '#':
				left, right = '#', '#'
			case 'O':
				left, right = '[', ']'
			case '@':
				left, right = '@', '.'
			default:
				continue
			}

			next.Set(Point{2 * x, y}, left)
			next.Set(Point{2*x + 1, y}, right)
		}
	}

	return next
}

// gps calculates the GPS value based on specific markers
func gps(grid *Grid, needle byte) int {
	result := 0

	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if grid.Get(Point{x, y}) == needle {
				result += 100*y + x
			}
		}
	}

	return result
}

func main() {
	grid, moves := Parse(input)

	fmt.Println("Part 1: what is the sum of all boxes' GPS coordinates?")
	fmt.Println(Part1(grid, moves))
	fmt.Println()
	fmt.Println("Part 2: What is the sum of all boxes' final GPS coordinates?")
	fmt.Println(Part2(grid, moves))
}
