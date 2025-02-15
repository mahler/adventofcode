package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

// Constants for direction and tile types
const (
	North = iota
	South
	West
	East
)

// Point represents a position in the maze
type Point struct {
	x, y int
}

// Direction represents a movement delta
type Direction struct {
	dx, dy int
}

// MazeState represents the current state of maze traversal
type MazeState struct {
	position Point
	visited  int
	distance int
}

// Maze represents the maze structure and state
type Maze struct {
	tiles     map[Point]rune
	size      Point
	entrance  Point
	entrances []Point
	keys      [26]*Point
	nkeys     int
}

// Directions for movement
var directions = []Direction{
	{0, -1}, // North
	{0, 1},  // South
	{-1, 0}, // West
	{1, 0},  // East
}

// NewMaze creates a new maze instance
func NewMaze() *Maze {
	return &Maze{
		tiles: make(map[Point]rune),
	}
}

// isKey checks if a character represents a key
func isKey(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

// isDoor checks if a character represents a door
func isDoor(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

// keyIdx returns the index for a key
func keyIdx(ch rune) int {
	return int(ch - 'a')
}

// keyBit returns the bit mask for a key
func keyBit(ch rune) int {
	return 1 << keyIdx(ch)
}

// doorIdx returns the index for a door
func doorIdx(ch rune) int {
	return int(ch - 'A')
}

// doorBit returns the bit mask for a door
func doorBit(ch rune) int {
	return 1 << doorIdx(ch)
}

// hasKey checks if the keys collection has a specific door's key
func hasKey(keys int, door rune) bool {
	return (keys & doorBit(door)) != 0
}

// step moves a point in the specified direction
func (p Point) step(dir int) Point {
	delta := directions[dir]
	return Point{p.x + delta.dx, p.y + delta.dy}
}

// ParseInput reads and parses the maze from a file
func (m *Maze) ParseInput(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var width, height int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > width {
			width = len(line)
		}

		for x, ch := range line {
			pos := Point{x, height}
			switch {
			case ch == '@':
				m.entrance = pos
				m.tiles[pos] = ch
			case isKey(ch):
				m.keys[keyIdx(ch)] = &pos
				m.nkeys++
				m.tiles[pos] = ch
			case ch == '.' || isDoor(ch):
				m.tiles[pos] = ch
			}
		}
		height++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	m.size = Point{width, height}
	return nil
}

// SolvePart1 solves the first part of the maze puzzle
func (m *Maze) SolvePart1() (int, error) {
	return m.bfs(m.entrance, 0)
}

// SolvePart2 solves the second part of the maze puzzle
func (m *Maze) SolvePart2() (int, error) {
	// Split the maze into quadrants
	m.splitMazeIntoQuadrants()

	keysByQuadrant := m.getKeysByQuadrant()
	initialVisited := m.calculateInitialVisited(keysByQuadrant)

	total := 0
	for i, entrance := range m.entrances {
		distance, err := m.bfs(entrance, initialVisited[i])
		if err != nil {
			return 0, fmt.Errorf("error solving quadrant %d: %w", i, err)
		}
		total += distance
	}
	return total, nil
}

// bfs performs a breadth-first search through the maze
func (m *Maze) bfs(start Point, initialVisited int) (int, error) {
	allKeys := (1 << m.nkeys) - 1
	state := make(map[[3]int]int)
	queue := list.New()
	queue.PushBack(MazeState{start, initialVisited, 0})

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(MazeState)
		ch := m.tiles[current.position]

		visited := current.visited
		if isKey(ch) {
			visited |= keyBit(ch)
			if visited == allKeys {
				return current.distance, nil
			}
		} else if isDoor(ch) && !hasKey(visited, ch) {
			continue
		}

		stateKey := [3]int{current.position.x, current.position.y, visited}
		if val, exists := state[stateKey]; exists && current.distance >= val {
			continue
		}
		state[stateKey] = current.distance

		for d := 0; d < 4; d++ {
			newPos := current.position.step(d)
			if _, exists := m.tiles[newPos]; exists {
				queue.PushBack(MazeState{newPos, visited, current.distance + 1})
			}
		}
	}

	return -1, fmt.Errorf("no solution found")
}

// splitMazeIntoQuadrants modifies the maze for part 2
func (m *Maze) splitMazeIntoQuadrants() {
	delete(m.tiles, m.entrance)
	for d := 0; d < 4; d++ {
		delete(m.tiles, m.entrance.step(d))
	}

	m.entrances = []Point{
		m.entrance.step(North).step(West),
		m.entrance.step(North).step(East),
		m.entrance.step(South).step(West),
		m.entrance.step(South).step(East),
	}

	for _, e := range m.entrances {
		m.tiles[e] = '@'
	}
}

// getKeysByQuadrant groups keys by their quadrant
func (m *Maze) getKeysByQuadrant() [4]int {
	var keysByQuadrant [4]int
	ex, ey := m.entrance.x, m.entrance.y

	for i := 0; i < m.nkeys; i++ {
		if m.keys[i] == nil {
			continue
		}
		key := m.keys[i]
		idx := 0
		if key.y > ey {
			idx |= 2
		}
		if key.x > ex {
			idx |= 1
		}
		keysByQuadrant[idx] |= (1 << i)
	}
	return keysByQuadrant
}

// calculateInitialVisited calculates the initial visited state for each quadrant
func (m *Maze) calculateInitialVisited(keysByQuadrant [4]int) [4]int {
	var initialVisited [4]int
	allKeys := (1 << m.nkeys) - 1
	for i := 0; i < 4; i++ {
		initialVisited[i] = allKeys ^ keysByQuadrant[i]
	}
	return initialVisited
}

func main() {
	maze := NewMaze()
	if err := maze.ParseInput("input.txt"); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	// Solve Part 1
	if shortest, err := maze.SolvePart1(); err != nil {
		fmt.Fprintf(os.Stderr, "Error solving part 1: %v\n", err)
	} else {
		fmt.Println("Part 1: How many steps is the shortest path that collects all of the keys?")
		fmt.Println(shortest)
	}

	// Solve Part 2
	maze = NewMaze() // Reset maze for part 2
	if err := maze.ParseInput("input.txt"); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	if shortest, err := maze.SolvePart2(); err != nil {
		fmt.Fprintf(os.Stderr, "Error solving part 2: %v\n", err)
	} else {
		fmt.Println()
		fmt.Println("Part 2: After updating your map and using the remote-controlled robots,")
		fmt.Println("what is the fewest steps necessary to collect all of the keys?")
		fmt.Println(shortest)
	}
}
