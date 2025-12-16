package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type AreaPair struct {
	Area int
	P0   Point
	P1   Point
}

func computeArea(p0, p1 Point) int {
	dx := 1 + abs(p0.X-p1.X)
	dy := 1 + abs(p0.Y-p1.Y)
	return dx * dy
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateStraightLine(p0, p1 Point) []Point {
	var points []Point
	if p0.X == p1.X {
		minY := min(p0.Y, p1.Y)
		maxY := max(p0.Y, p1.Y)
		for y := minY + 1; y < maxY; y++ {
			points = append(points, Point{p0.X, y})
		}
	} else if p0.Y == p1.Y {
		minX := min(p0.X, p1.X)
		maxX := max(p0.X, p1.X)
		for x := minX + 1; x < maxX; x++ {
			points = append(points, Point{x, p0.Y})
		}
	} else {
		panic("Only horizontal or vertical lines are supported")
	}
	return points
}

func removeAdjacent(coords []int) []int {
	var filtered []int
	prev := -1000000 // Sentinel value
	for _, c := range coords {
		if prev == -1000000 || c-prev > 1 {
			filtered = append(filtered, c)
		} else if len(filtered) > 0 {
			filtered = filtered[:len(filtered)-1]
		}
		prev = c
	}
	return filtered
}

func bisectLeft(arr []int, x int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] >= x })
}

func containsRectangle(p0, p1 Point, yLists, xLists map[int][]int) bool {
	xMin := min(p0.X, p1.X) + 1
	xMax := max(p0.X, p1.X) - 1
	yMin := min(p0.Y, p1.Y) + 1
	yMax := max(p0.Y, p1.Y) - 1

	// Test all four edges
	if bisectLeft(yLists[xMin], yMin) != bisectLeft(yLists[xMin], yMax) {
		return false
	}
	if bisectLeft(yLists[xMax], yMin) != bisectLeft(yLists[xMax], yMax) {
		return false
	}
	if bisectLeft(xLists[yMin], xMin) != bisectLeft(xLists[yMin], xMax) {
		return false
	}
	if bisectLeft(xLists[yMax], xMin) != bisectLeft(xLists[yMax], xMax) {
		return false
	}

	// Check all interior rows and columns
	for y := yMin; y <= yMax; y++ {
		if bisectLeft(xLists[y], xMin) != bisectLeft(xLists[y], xMax) {
			return false
		}
	}
	for x := xMin; x <= xMax; x++ {
		if bisectLeft(yLists[x], yMin) != bisectLeft(yLists[x], yMax) {
			return false
		}
	}

	// Check if rectangle is inside the shape
	if bisectLeft(xLists[yMax], xMax)%2 == 0 {
		return false
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reds []Point

	// Read red points
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		reds = append(reds, Point{x, y})
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Generate green points (boundary)
	var greens []Point
	for i := 0; i < len(reds); i++ {
		p0 := reds[i]
		p1 := reds[(i+1)%len(reds)]
		greens = append(greens, generateStraightLine(p0, p1)...)
	}

	// Generate all pairs and compute areas
	var areaPairs []AreaPair
	for i := 0; i < len(reds); i++ {
		for j := i + 1; j < len(reds); j++ {
			area := computeArea(reds[i], reds[j])
			areaPairs = append(areaPairs, AreaPair{area, reds[i], reds[j]})
		}
	}

	// Sort by area in descending order
	sort.Slice(areaPairs, func(i, j int) bool {
		return areaPairs[i].Area > areaPairs[j].Area
	})

	// PART 1
	fmt.Println(areaPairs[0].Area)

	// PART 2
	// Build coordinate lists
	yLists := make(map[int][]int)
	xLists := make(map[int][]int)

	allPoints := append([]Point{}, reds...)
	allPoints = append(allPoints, greens...)

	for _, point := range allPoints {
		yLists[point.X] = append(yLists[point.X], point.Y)
		xLists[point.Y] = append(xLists[point.Y], point.X)
	}

	// Sort all lists
	for x := range yLi
