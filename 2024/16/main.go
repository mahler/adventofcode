package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y, dx, dy int
}

type PriorityItem struct {
	coord Coord
	time  int
	index int
}

type PriorityQueue []*PriorityItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].time < pq[j].time }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

var directions = [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func readGrid(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func findSpecialCoords(grid []string, char rune) [][2]int {
	var coords [][2]int
	for i, line := range grid {
		for j, c := range line {
			if c == char {
				coords = append(coords, [2]int{i, j})
			}
		}
	}
	return coords
}

func buildWalls(grid []string) map[Coord]bool {
	walls := make(map[Coord]bool)
	for i, line := range grid {
		for j, char := range line {
			if char == '#' {
				walls[Coord{x: i, y: j}] = true
			}
		}
	}
	return walls
}

func distanceFunction(loc Coord, walls map[Coord]bool) [][2]interface{} {
	var output [][2]interface{}
	newLoc := Coord{x: loc.x + loc.dx, y: loc.y + loc.dy, dx: loc.dx, dy: loc.dy}

	if !walls[Coord{x: newLoc.x, y: newLoc.y}] {
		output = append(output, [2]interface{}{newLoc, 1})
	}

	for _, dir := range directions {
		if dir[0] != loc.dx || dir[1] != loc.dy {
			newCoord := Coord{x: loc.x, y: loc.y, dx: dir[0], dy: dir[1]}
			output = append(output, [2]interface{}{newCoord, 1000})
		}
	}
	return output
}

func dijkstraFunction(start Coord, ends []Coord, distFunc func(Coord) [][2]interface{}, walls map[Coord]bool) int {
	visited := make(map[Coord]bool)
	pq := make(PriorityQueue, 1)
	pq[0] = &PriorityItem{
		coord: start,
		time:  0,
		index: 0,
	}
	heap.Init(&pq)

	timeDict := make(map[Coord]int)
	timeDict[start] = 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityItem)
		loc := item.coord
		time := item.time

		if visited[loc] {
			continue
		}
		visited[loc] = true

		for _, end := range ends {
			if loc.x == end.x && loc.y == end.y {
				return time
			}
		}

		neighbours := distFunc(loc)
		for _, neighbour := range neighbours {
			coord := neighbour[0].(Coord)
			distance := neighbour[1].(int)
			newTime := time + distance

			if !visited[coord] && (timeDict[coord] == 0 || newTime < timeDict[coord]) {
				timeDict[coord] = newTime
				heap.Push(&pq, &PriorityItem{
					coord: coord,
					time:  newTime,
				})
			}
		}
	}
	return -1
}

func dijkstraFunction2(starts []Coord, distFunc func(Coord) [][2]interface{}, walls map[Coord]bool) map[Coord]int {
	visited := make(map[Coord]bool)
	pq := make(PriorityQueue, len(starts))
	output := make(map[Coord]int)
	timeDict := make(map[Coord]int)

	for i, start := range starts {
		pq[i] = &PriorityItem{
			coord: start,
			time:  0,
			index: i,
		}
		timeDict[start] = 0
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityItem)
		loc := item.coord
		time := item.time

		if visited[loc] {
			continue
		}
		output[loc] = time
		visited[loc] = true

		neighbours := distFunc(loc)
		for _, neighbour := range neighbours {
			coord := neighbour[0].(Coord)
			distance := neighbour[1].(int)
			newTime := time + distance

			if !visited[coord] && (timeDict[coord] == 0 || newTime < timeDict[coord]) {
				timeDict[coord] = newTime
				heap.Push(&pq, &PriorityItem{
					coord: coord,
					time:  newTime,
				})
			}
		}
	}
	return output
}

func main() {
	grid := readGrid("input.txt")
	walls := buildWalls(grid)

	startCoords := findSpecialCoords(grid, 'S')
	endCoords := findSpecialCoords(grid, 'E')

	if len(startCoords) != 1 {
		panic("Multiple or no start coordinates")
	}
	if len(endCoords) != 1 {
		panic("Multiple or no end coordinates")
	}

	start := Coord{x: startCoords[0][0], y: startCoords[0][1], dx: 0, dy: 1}

	var ends []Coord
	for _, dir := range directions {
		ends = append(ends, Coord{
			x:  endCoords[0][0],
			y:  endCoords[0][1],
			dx: dir[0],
			dy: dir[1],
		})
	}

	distFunc := func(loc Coord) [][2]interface{} {
		return distanceFunction(loc, walls)
	}

	answer := dijkstraFunction(start, ends, distFunc, walls)
	fmt.Println("Part 1: What is the lowest score a Reindeer could possibly get?")
	fmt.Println(answer)

	distanceFromStart := dijkstraFunction2([]Coord{start}, distFunc, walls)
	distanceFromEnd := dijkstraFunction2(ends, distFunc, walls)

	coords := make(map[Coord]bool)
	for loc, startDist := range distanceFromStart {
		for _, dir := range directions {
			endLoc := Coord{x: loc.x, y: loc.y, dx: -dir[0], dy: -dir[1]}
			if startDist+distanceFromEnd[endLoc] == answer {
				coords[Coord{x: loc.x, y: loc.y}] = true
			}
		}
	}
	fmt.Println()
	fmt.Println("Part 2:  How many tiles are part of at least one of the best paths through the maze?")
	fmt.Println(len(coords))
}
