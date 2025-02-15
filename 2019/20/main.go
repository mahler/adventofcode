package main

import (
	"fmt"
	"os"
	"strings"
)

// Point represents a 3D coordinate
type Point struct {
	x, y, z int
}

// Add returns a new Point that is the sum of the receiver and other
func (p Point) Add(other Point) Point {
	return Point{
		x: p.x + other.x,
		y: p.y + other.y,
		z: p.z + other.z,
	}
}

// Bounds represents a rectangular area
type Bounds struct {
	topLeft     Point
	bottomRight Point
}

// State represents the current position and distance traveled
type State struct {
	pos      Point
	distance int
}

// Portal represents a connection between two points in the maze
type Portal struct {
	label string
	from  Point
	to    Point
	outer bool
}

// Maze represents the complete maze structure
type Maze struct {
	grid    map[Point]bool
	portals map[Point]Portal
	start   Point
	end     Point
}

// Common direction vectors
var directions = []Point{
	{0, -1, 0}, // up
	{1, 0, 0},  // right
	{0, 1, 0},  // down
	{-1, 0, 0}, // left
}

func readInput(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	// Trim carriage returns and empty lines
	var cleaned []string
	for _, line := range lines {
		if trimmed := strings.TrimRight(line, "\r"); trimmed != "" {
			cleaned = append(cleaned, trimmed)
		}
	}
	return cleaned, nil
}

func findInnerBounds(lines []string, outer Bounds) Bounds {
	var inner Bounds
	for y := outer.topLeft.y; y <= outer.bottomRight.y; y++ {
		if idx := strings.Index(lines[y][outer.topLeft.x:outer.bottomRight.x], " "); idx >= 0 {
			inner.topLeft = Point{
				x: idx + outer.topLeft.x - 1,
				y: y - 1,
				z: 0,
			}
			inner.bottomRight = Point{
				x: outer.bottomRight.x - inner.topLeft.x + outer.topLeft.x,
				y: outer.bottomRight.y - inner.topLeft.y + outer.topLeft.y,
				z: 0,
			}
			break
		}
	}
	return inner
}

func parsePortal(lines []string, pos Point, label string, isOuter bool) Portal {
	return Portal{
		label: label,
		from:  pos,
		outer: isOuter,
	}
}

func newMaze(lines []string) *Maze {
	m := &Maze{
		grid:    make(map[Point]bool),
		portals: make(map[Point]Portal),
	}

	outer := Bounds{
		topLeft:     Point{2, 2, 0},
		bottomRight: Point{len(lines[0]) - 3, len(lines) - 3, 0},
	}
	inner := findInnerBounds(lines, outer)

	// Parse the maze grid and portals
	for y := outer.topLeft.y; y <= outer.bottomRight.y; y++ {
		for x := outer.topLeft.x; x <= outer.bottomRight.x; x++ {
			if lines[y][x] != '.' {
				continue
			}

			pos := Point{x, y, 0}
			m.grid[pos] = true

			// Check for portal labels in all directions
			var label string
			var portalPos Point
			isOuter := false

			switch {
			case y == outer.topLeft.y:
				label = lines[y-2][x:x+1] + lines[y-1][x:x+1]
				portalPos = Point{x, y - 1, 0}
				isOuter = true
			case y == outer.bottomRight.y:
				label = lines[y+1][x:x+1] + lines[y+2][x:x+1]
				portalPos = Point{x, y + 1, 0}
				isOuter = true
			case x == outer.topLeft.x:
				label = lines[y][x-2 : x]
				portalPos = Point{x - 1, y, 0}
				isOuter = true
			case x == outer.bottomRight.x:
				label = lines[y][x+1 : x+3]
				portalPos = Point{x + 1, y, 0}
				isOuter = true
			case y == inner.bottomRight.y && x > inner.topLeft.x && x < inner.bottomRight.x:
				label = lines[y-2][x:x+1] + lines[y-1][x:x+1]
				portalPos = Point{x, y - 1, 0}
			case y == inner.topLeft.y && x > inner.topLeft.x && x < inner.bottomRight.x:
				label = lines[y+1][x:x+1] + lines[y+2][x:x+1]
				portalPos = Point{x, y + 1, 0}
			case x == inner.bottomRight.x && y > inner.topLeft.y && y < inner.bottomRight.y:
				label = lines[y][x-2 : x]
				portalPos = Point{x - 1, y, 0}
			case x == inner.topLeft.x && y > inner.topLeft.y && y < inner.bottomRight.y:
				label = lines[y][x+1 : x+3]
				portalPos = Point{x + 1, y, 0}
			}

			if label == "" {
				continue
			}

			switch label {
			case "AA":
				m.start = pos
			case "ZZ":
				m.end = pos
			default:
				portal := parsePortal(lines, pos, label, isOuter)
				m.portals[portalPos] = portal
				m.grid[portalPos] = true
			}
		}
	}

	// Link portals together
	for pos, p1 := range m.portals {
		for _, p2 := range m.portals {
			if p1.label == p2.label && p1.from != p2.from {
				p1.to = p2.from
				m.portals[pos] = p1
			}
		}
	}

	return m
}

func (m *Maze) solve(recursive bool) (int, error) {
	queue := []State{{pos: m.start}}
	visited := map[Point]bool{m.start: true}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			next := current.pos.Add(dir)

			if next == m.end {
				return current.distance + 1, nil
			}

			flatNext := Point{next.x, next.y, 0}
			if !m.grid[flatNext] || visited[next] {
				continue
			}

			visited[next] = true

			if portal, exists := m.portals[flatNext]; exists {
				if recursive && (current.pos.z > 0 || !portal.outer) {
					next = Point{portal.to.x, portal.to.y, current.pos.z}
					if portal.outer {
						next.z--
					} else {
						next.z++
					}
					visited[next] = true
				} else if !recursive {
					next = portal.to
				}
			}

			queue = append(queue, State{next, current.distance + 1})
		}
	}

	return 0, fmt.Errorf("no path found")
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	maze := newMaze(lines)

	part1, err := maze.solve(false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error solving part 1: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Part 1: In your maze, how many steps does it take to get from the open tile marked AA to the open tile marked ZZ?")
	fmt.Println(part1)

	part2, err := maze.solve(true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error solving part 2: %v\n", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println("Part 2: In your maze, when accounting for recursion,")
	fmt.Println("how many steps does it take to get from the open tile")
	fmt.Println("marked AA to the open tile marked ZZ, both at the outermost layer?")
	fmt.Println(part2)
}
