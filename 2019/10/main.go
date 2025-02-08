package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

type AnglePoint struct {
	angle float64
	point Point
}

func initAsteroids() []Point {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var asteroids []Point
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, a := range line {
			if a == '#' {
				asteroids = append(asteroids, Point{x, y})
			}
		}
	}
	return asteroids
}

func angle(start, end Point) float64 {
	result := math.Atan2(float64(end.x-start.x), float64(start.y-end.y)) * 180 / math.Pi
	if result < 0 {
		return 360 + result
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	asteroids := initAsteroids()

	// Part 1
	var result Point
	maxVisible := 0

	for _, start := range asteroids {
		angles := make(map[float64]bool)
		for _, end := range asteroids {
			if start != end {
				angles[angle(start, end)] = true
			}
		}
		if len(angles) > maxVisible {
			maxVisible = len(angles)
			result = start
		}
	}

	// fmt.Printf("x %d y %d\n", result.x, result.y)
	fmt.Println("Part 1: How many other asteroids can be detected from that location?")
	fmt.Println(maxVisible)

	// Part 2
	// Remove the station asteroid
	for i, a := range asteroids {
		if a == result {
			asteroids = append(asteroids[:i], asteroids[i+1:]...)
			break
		}
	}

	// Create sorted angles list
	angles := make([]AnglePoint, 0, len(asteroids))
	for _, end := range asteroids {
		angles = append(angles, AnglePoint{
			angle: angle(result, end),
			point: end,
		})
	}

	sort.Slice(angles, func(i, j int) bool {
		if angles[i].angle == angles[j].angle {
			dist1 := abs(result.x-angles[i].point.x) + abs(result.y-angles[i].point.y)
			dist2 := abs(result.x-angles[j].point.x) + abs(result.y-angles[j].point.y)
			return dist1 < dist2
		}
		return angles[i].angle < angles[j].angle
	})

	idx := 0
	var last AnglePoint
	lastAngle := -1.0
	cnt := 0

	for cnt < 200 && len(angles) > 0 {
		if idx >= len(angles) {
			idx = 0
			lastAngle = -1
		}
		if lastAngle == angles[idx].angle {
			idx++
			continue
		}
		last = angles[idx]
		angles = append(angles[:idx], angles[idx+1:]...)
		lastAngle = last.angle
		cnt++
	}
	//	fmt.Printf("vaporized %d: %v %d\n", cnt, last.point)

	fmt.Println()
	fmt.Println("Part 2: Win the bet by determining which asteroid that will be;")
	fmt.Println("what do you get if you multiply its X coordinate by 100 and then add its Y coordinate?")
	fmt.Println(last.point.x*100 + last.point.y)
}
