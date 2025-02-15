package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func isKey(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

func isDoor(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

func keyIdx(ch rune) int {
	return int(ch - 'a')
}

func keyBit(ch rune) int {
	return 1 << keyIdx(ch)
}

func doorIdx(ch rune) int {
	if ch < 'A' || ch > 'Z' {
		fmt.Printf("char '%c' is not a door\n", ch)
	}
	return int(ch - 'A')
}

func doorBit(ch rune) int {
	return 1 << doorIdx(ch)
}

func hasKey(keys int, door rune) bool {
	return (keys & doorBit(door)) != 0
}

func step(pos Point, direction int) Point {
	deltas := []Point{
		{0, -1}, // north
		{0, 1},  // south
		{-1, 0}, // west
		{1, 0},  // east
	}
	return Point{pos.x + deltas[direction].x, pos.y + deltas[direction].y}
}

var (
	tiles    = make(map[Point]rune)
	size     Point
	entrance Point
	keys     [26]*Point
	nkeys    int
)

func parseInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	w, y := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > w {
			w = len(line)
		}
		for x, ch := range line {
			switch {
			case ch == '@':
				entrance = Point{x, y}
				tiles[Point{x, y}] = ch
			case isKey(ch):
				keys[keyIdx(ch)] = &Point{x, y}
				nkeys++
				tiles[Point{x, y}] = ch
			case ch == '.' || isDoor(ch):
				tiles[Point{x, y}] = ch
			}
		}
		y++
	}

	size = Point{w, y}
}

func printMap() {
	fmt.Printf("%dx%d map:\n", size.x, size.y)
	for y := 0; y < size.y; y++ {
		for x := 0; x < size.x; x++ {
			if ch, exists := tiles[Point{x, y}]; exists {
				fmt.Print(string(ch))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func bfs() int {
	allKeys := (1 << nkeys) - 1
	state := make(map[[3]int]int) // (x, y, visited) -> distance
	queue := []struct {
		x, y, visited, distance int
	}{{entrance.x, entrance.y, 0, 0}}

	progress := 0

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		pos := Point{n.x, n.y}
		ch := tiles[pos]

		visited := n.visited
		if isKey(ch) {
			visited |= keyBit(ch)
			if visited == allKeys {
				return n.distance
			}
		} else if isDoor(ch) && !hasKey(visited, ch) {
			continue
		}

		stateKey := [3]int{n.x, n.y, visited}
		if val, exists := state[stateKey]; exists && n.distance >= val {
			continue
		}
		state[stateKey] = n.distance

		for d := 0; d < 4; d++ {
			newPos := step(pos, d)
			if _, exists := tiles[newPos]; exists {
				queue = append(queue, struct {
					x, y, visited, distance int
				}{newPos.x, newPos.y, visited, n.distance + 1})
			}
		}

		progress++
		if progress%1000 == 0 {
			fmt.Printf("\r%d", progress/1000)
		}
	}

	return -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	parseInput(os.Args[1])
	printMap()

	fmt.Println()
	shortest := bfs()
	fmt.Printf("shortest path for all keys is %d\n", shortest)
}
