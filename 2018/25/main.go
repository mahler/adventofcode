package main

//
// READY TO COMMIT
//

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
	Z int
	T int
}

func (p Point) distance(pp Point) int {
	deltaX := p.X - pp.X
	if deltaX < 0 {
		deltaX = -deltaX
	}
	deltaY := p.Y - pp.Y
	if deltaY < 0 {
		deltaY = -deltaY
	}
	deltaZ := p.Z - pp.Z
	if deltaZ < 0 {
		deltaZ = -deltaZ
	}
	deltaT := p.T - pp.T
	if deltaT < 0 {
		deltaT = -deltaT
	}

	return deltaX + deltaY + deltaZ + deltaT
}

func main() {
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("DAY25, Part 1: Four-Dimensional Adventure")

	data, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}
	fileContent := strings.Split(string(data), "\n")

	var points []Point
	for _, line := range fileContent {
		coords := strings.Split(line, ",")

		var point Point
		point.X, _ = strconv.Atoi(coords[0])
		point.Y, _ = strconv.Atoi(coords[1])
		point.Z, _ = strconv.Atoi(coords[2])
		point.T, _ = strconv.Atoi(coords[3])

		points = append(points, point)
	}

	inDistance := make(map[Point][]Point)
	for i, point := range points {
		for _, otherPoint := range points[i+1:] {
			if point.distance(otherPoint) <= 3 {
				inDistance[point] = append(inDistance[point], otherPoint)
				inDistance[otherPoint] = append(inDistance[otherPoint], point)
			}
		}
		if _, ok := inDistance[point]; !ok {
			inDistance[point] = nil
		}
	}
	constellations := make(map[Point]int)
	var currentConstellation int

	for len(constellations) != len(points) {
		var thisPoint Point
		for p := range inDistance {
			thisPoint = p
			break
		}

		toVisit := map[Point]bool{thisPoint: true}
		for len(toVisit) > 0 {
			for p := range toVisit {
				thisPoint = p
				break
			}

			for _, p := range inDistance[thisPoint] {
				if _, ok := constellations[p]; !ok {
					toVisit[p] = true
				}
			}
			constellations[thisPoint] = currentConstellation
			// Clean up...
			delete(toVisit, thisPoint)
			delete(inDistance, thisPoint)
		}

		currentConstellation++
	}

	fmt.Println("How many constellations are formed by the fixed points in spacetime?")
	fmt.Println(currentConstellation)
}
