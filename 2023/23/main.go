package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func getJunctions(grid []string) map[Point]bool {
	junctions := map[Point]bool{}
	for row, line := range grid {
		for col, char := range line {
			if char == '#' {
				continue
			}
			point := Point{col, row}
			neighbours := 0
			for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				next := Point{point.x + dir.x, point.y + dir.y}
				if IsInsideGrid(grid, next) && grid[next.y][next.x] != '#' {
					neighbours++
				}
			}
			if neighbours > 2 {
				junctions[point] = true
			}
		}
	}
	return junctions
}

type PathTo struct {
	end           Point
	length, index int
}

func getPaths(grid []string, junctions map[Point]bool) map[Point][]PathTo {
	paths := map[Point][]PathTo{}
	junctionIndex := 0
	for junctionPoint := range junctions {
		for _, startDir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			currentPoint := Point{junctionPoint.x + startDir.x, junctionPoint.y + startDir.y}
			if IsInsideGrid(grid, currentPoint) && grid[currentPoint.y][currentPoint.x] != '#' {
				path := getPath(grid, junctionPoint, currentPoint, startDir, 1, junctions)
				path.index = junctionIndex
				paths[junctionPoint] = append(paths[junctionPoint], path)
			}
		}
		junctionIndex++
	}
	return paths
}

func getPath(grid []string, pathStart, currentPoint, currentDir Point, pathLength int, junctions map[Point]bool) PathTo {
	for _, dir := range [3]Point{currentDir, dirLeft(currentDir), dirRight(currentDir)} {
		next := Point{currentPoint.x + dir.x, currentPoint.y + dir.y}
		if grid[next.y][next.x] != '#' {
			if _, found := junctions[next]; found {
				return PathTo{next, pathLength + 1, 0}
			} else {
				return getPath(grid, pathStart, next, dir, pathLength+1, junctions)
			}
		}
	}
	return PathTo{Point{-1, -1}, 0, 0}
}

func findLongestPath(grid []string, paths map[Point][]PathTo, start, end Point, step int, visited []bool) int {
	maxStep := 0
	for _, path := range paths[start] {
		index := paths[path.end][0].index
		if !visited[index] {
			if path.end == end {
				return step + path.length
			}
			visited[index] = true
			maxStep = max(maxStep, findLongestPath(grid, paths, path.end, end, step+path.length, visited))
			visited[index] = false
		}
	}
	return maxStep
}

func walkTrail(grid []string, start, currentDir Point, visited map[Point]int) {
	current := start
	currentStep := visited[current]

	for _, dir := range [3]Point{currentDir, dirLeft(currentDir), dirRight(currentDir)} {
		next := Point{current.x + dir.x, current.y + dir.y}
		if IsInsideGrid(grid, next) && grid[next.y][next.x] != '#' {
			char := grid[next.y][next.x]
			oppositeChar := map[Point]byte{{1, 0}: '<', {-1, 0}: '>', {0, 1}: '^', {0, -1}: 'v'}
			if oppositeChar[dir] == char {
				continue
			}

			if val, found := visited[next]; !found || val < currentStep+1 {
				visited[next] = currentStep + 1
				walkTrail(grid, next, dir, visited)
			}
		}
	}
}

func readGridFromFile(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func IsInsideGrid(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
}

func dirLeft(p Point) Point {
	return Point{p.y, -p.x}
}

func dirRight(p Point) Point {
	return Point{-p.y, p.x}
}

func main() {
	grid := readGridFromFile("input.txt")
	start, end := Point{1, 0}, Point{len(grid[0]) - 2, len(grid) - 1}

	visited := map[Point]int{start: 0}
	currentDir := Point{0, 1}
	walkTrail(grid, start, currentDir, visited)

	var part1Result = visited[end]
	fmt.Println("Part 1: Find the longest hike you can take through the hiking trails listed on your map.")
	fmt.Println("How many steps long is the longest hike?")
	fmt.Println(part1Result)

	// Part 2
	junctions := getJunctions(grid)
	junctions[start] = true
	junctions[end] = true

	paths := getPaths(grid, junctions)
	visited2 := make([]bool, len(junctions))
	visited2[paths[start][0].index] = true
	var part2Result = findLongestPath(grid, paths, start, end, 0, visited2)
	fmt.Println()
	fmt.Println("Part 2: Find the longest hike you can take through the surprisingly dry hiking trails listed on your map.")
	fmt.Println("How many steps long is the longest hike?")
	fmt.Println(part2Result)
}
