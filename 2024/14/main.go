package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	T    = 100
	XMax = 101
	YMax = 103
	XMid = XMax / 2 // 50
	YMid = YMax / 2 // 51
)

type Point struct {
	x, y int
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pos []Point

	var positions []Point
	var velocities []Point

	// Process each line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, " ")

		// Parse position
		p := strings.Split(strings.Split(parts[0], "=")[1], ",")
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])

		// Parse velocity
		v := strings.Split(strings.Split(parts[1], "=")[1], ",")
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])

		// Calculate final position
		xFinal := (px + vx*T) % XMax
		if xFinal < 0 {
			xFinal += XMax
		}
		yFinal := (py + vy*T) % YMax
		if yFinal < 0 {
			yFinal += YMax
		}

		pos = append(pos, Point{xFinal, yFinal})

		positions = append(positions, Point{px, py})
		velocities = append(velocities, Point{vx, vy})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Count points in each quadrant
	var q1, q2, q3, q4 int
	for _, p := range pos {
		if p.x < XMid && p.y < YMid {
			q1++
		} else if p.x < XMid && p.y > YMid {
			q2++
		} else if p.x > XMid && p.y < YMid {
			q3++
		} else if p.x > XMid && p.y > YMid {
			q4++
		}
	}

	// Print result
	fmt.Println("Part 1: What will the safety factor be after exactly 100 seconds have elapsed?")
	fmt.Println(q1 * q2 * q3 * q4)

	// ----------------------
	T := 0
	for {
		// Check for distinct positions
		posSet := make(map[Point]bool)
		allDistinct := true

		for _, pos := range positions {
			if posSet[pos] {
				allDistinct = false
				break
			}
			posSet[pos] = true
		}

		if allDistinct {
			break
		}

		// Update positions
		for i := range positions {
			px := positions[i].x
			py := positions[i].y
			vx := velocities[i].x
			vy := velocities[i].y

			newX := (px + vx) % XMax
			if newX < 0 {
				newX += XMax
			}
			newY := (py + vy) % YMax
			if newY < 0 {
				newY += YMax
			}

			positions[i] = Point{newX, newY}
		}

		T++
	}

	fmt.Println()
	fmt.Println("Part 2: What is the fewest number of seconds that must elapse for the robots to display the Easter egg?")
	fmt.Println(T)

}
