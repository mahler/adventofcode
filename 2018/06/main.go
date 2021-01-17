package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// Point defines an x,y coordinate
type Point struct {
	x, y int
}

// GetDistance returns the Manhattan distance between points p and other point
func (p Point) GetDistance(other Point) int {
	xDist := AbsInt(p.x - other.x)
	yDist := AbsInt(p.y - other.y)
	return xDist + yDist
}

// FindNearestPoint determines which point in a list of points is closest to point p.
func (p Point) FindNearestPoint(points []Point) (Point, error) {
	distanceMap := make(map[int][]Point)

	for _, pt := range points {
		dist := p.GetDistance(pt)
		distanceMap[dist] = append(distanceMap[dist], pt)
	}

	distances := make([]int, len(distanceMap))
	i := 0
	for d := range distanceMap {
		distances[i] = d
		i++
	}
	sort.Ints(distances)

	nearest := distanceMap[distances[0]]
	if len(nearest) > 1 {
		return Point{}, errors.New("More than one point tied for closest")
	}

	return nearest[0], nil
}

func main() {
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("Day 6: Chronal Coordinates")
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}
	fileRows := strings.Split(string(fileContent), "\n")

	points := make([]Point, 0)
	for _, fileLine := range fileRows {
		coords := strings.Split(fileLine, ", ")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}

		pt := Point{x, y}
		points = append(points, pt)
	}

	n := len(points)
	x := make([]int, n)
	y := make([]int, n)
	for i, pt := range points {
		x[i] = pt.x
		y[i] = pt.y
	}
	sort.Ints(x)
	sort.Ints(y)
	xMin := x[0]
	xMax := x[n-1]
	yMin := y[0]
	yMax := y[n-1]

	blankCanvas := make(map[Point]bool)
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			pt := Point{i, j}
			blankCanvas[pt] = true
		}
	}

	canvas := make([]Point, len(blankCanvas))
	i := 0
	for pt := range blankCanvas {
		canvas[i] = pt
		i++
	}

	areas := make(map[Point]int)

	for _, pt := range canvas {
		nearest, err := pt.FindNearestPoint(points)
		if err == nil {
			areas[nearest]++
		}
	}

	areaVals := make([]int, len(areas))
	i = 0
	for _, val := range areas {
		areaVals[i] = val
		i++
	}
	sort.Ints(areaVals)

	fmt.Println()
	fmt.Println("2018")
	fmt.Println("Day 06, Part 1: Chronal Coordinates")
	fmt.Println("What is the size of the largest area that isn't infinite?")
	fmt.Println(areaVals[len(areaVals)-1])

	// -------------------------------
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("What is the size of the region containing all locations which")
	fmt.Println("have a total distance to all given coordinates of less than 10000?")
	area := 0

	for _, origin := range canvas {
		sum := 0
		for _, pt := range points {
			sum += origin.GetDistance(pt)
		}
		if sum < 10000 {
			area++
		}
	}

	fmt.Println(area)
}

// AbsInt calculates the absolute value of an integer
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
