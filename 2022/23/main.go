package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	day      = 23
	testData = `
....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
`
)

// Point represents a 2D coordinate
type Point struct {
	x, y int
}

// Map is a mapping of positions to lists of previous positions
type Map map[Point][]Point

func countEmpty(mapData Map) int {
	if len(mapData) == 0 {
		return 0
	}

	minX, maxX := 1<<31-1, -1<<31
	minY, maxY := 1<<31-1, -1<<31

	for k := range mapData {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}

	h := maxX - minX + 1
	w := maxY - minY + 1
	return h*w - len(mapData)
}

func makeMoves(mapData Map, ruleOrder string) Map {
	proposedMoves := make(Map)

	// Direction offsets for each possible move
	directions := map[string]Point{
		"N":  {-1, 0},
		"NE": {-1, 1},
		"E":  {0, 1},
		"SE": {1, 1},
		"S":  {1, 0},
		"SW": {1, -1},
		"W":  {0, -1},
		"NW": {-1, -1},
	}

	// First phase: propose moves
	for pos := range mapData {
		proposedMoves[pos] = []Point{pos}
		surroundings := make(map[string]struct {
			occupied bool
			pos      Point
		})

		// Check all surrounding positions
		for dir, offset := range directions {
			newPos := Point{pos.x + offset.x, pos.y + offset.y}
			_, exists := mapData[newPos]
			surroundings[dir] = struct {
				occupied bool
				pos      Point
			}{exists, newPos}
		}

		// Check if elf needs to move
		hasNeighbor := false
		for _, s := range surroundings {
			if s.occupied {
				hasNeighbor = true
				break
			}
		}

		if hasNeighbor {
			// Try each direction in order
			for _, direction := range strings.Split(ruleOrder, "") {
				canMove := true
				// Check the three positions in the direction
				for dir := range surroundings {
					if strings.Contains(dir, direction) && surroundings[dir].occupied {
						canMove = false
						break
					}
				}

				if canMove {
					newPos := surroundings[direction].pos
					proposedMoves[newPos] = append(proposedMoves[newPos], pos)
					delete(proposedMoves, pos)
					break
				}
			}
		}
	}

	// Second phase: resolve conflicts
	finalMoves := make(Map)
	for dest, sources := range proposedMoves {
		if len(sources) == 1 {
			finalMoves[dest] = sources
		} else {
			for _, source := range sources {
				finalMoves[source] = []Point{source}
			}
		}
	}

	return finalMoves
}

func mapsEqual(m1, m2 Map) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, exists := m2[k]
		if !exists || len(v1) != len(v2) {
			return false
		}
		for i := range v1 {
			if v1[i] != v2[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	// Read input file...
	fileData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	dataStr := string(fileData)
	lines := strings.Split(strings.TrimSpace(dataStr), "\n")

	// Setup map
	mapData := make(Map)

	for x, line := range lines {
		for y, c := range line {
			if c == '#' {
				pos := Point{x, y}
				mapData[pos] = []Point{pos}
			}
		}
	}

	// Set ruleorder and run rules...
	ruleOrder := "NSWE"

	for i := 0; i < 10; i++ {
		mapData = makeMoves(mapData, ruleOrder)
		ruleOrder = ruleOrder[1:] + ruleOrder[:1]
	}
	fmt.Println("Part 1: How many empty ground tiles does that rectangle contain?")
	fmt.Println(countEmpty(mapData))

	// Reset map data for part2
	mapData = make(Map)
	for x, line := range lines {
		for y, c := range line {
			if c == '#' {
				pos := Point{x, y}
				mapData[pos] = []Point{pos}
			}
		}
	}

	ruleOrder = "NSWE"
	round := 0

	for {
		newMapData := makeMoves(mapData, ruleOrder)
		if mapsEqual(mapData, newMapData) {
			break
		} else {
			round++
		}
		mapData = newMapData
		ruleOrder = ruleOrder[1:] + ruleOrder[:1]
	}
	fmt.Println()
	fmt.Println("Part 2: What is the number of the first round where no Elf moves?")
	fmt.Println(round)
}
