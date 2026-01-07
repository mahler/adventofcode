package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

type coord struct {
	x, y int
}

func readCoords() ([]coord, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var coords []coord
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			continue
		}
		coords = append(coords, coord{x, y})
	}
	return coords, scanner.Err()
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part1() {
	coords, err := readCoords()
	if err != nil {
		log.Fatal(err)
	}

	area := 0
	for i := 0; i < len(coords); i++ {
		x1, y1 := coords[i].x, coords[i].y
		for j := i + 1; j < len(coords); j++ {
			x2, y2 := coords[j].x, coords[j].y
			if x1 != x2 && y1 != y2 {
				a := (abs(x1-x2) + 1) * (abs(y1-y2) + 1)
				area = max(area, a)
			}
		}
	}
	fmt.Println(area)
}

func part2() {
	coords, err := readCoords()
	if err != nil {
		log.Fatal(err)
	}

	maxX, maxY := 0, 0
	for _, c := range coords {
		maxX = max(maxX, c.x)
		maxY = max(maxY, c.y)
	}

	// Vertical sweep - build spans
	type span struct {
		x1, x2 int
	}
	spans := make([]*span, maxY+2)

	// Append first coord to create closed loop
	coords = append(coords, coords[0])

	for i := 1; i < len(coords); i++ {
		x1, y1 := coords[i-1].x, coords[i-1].y
		x2, y2 := coords[i].x, coords[i].y

		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		for y := y1; y <= y2; y++ {
			if spans[y] == nil {
				spans[y] = &span{x1, x2}
			} else {
				spans[y].x1 = min(x1, spans[y].x1)
				spans[y].x2 = max(x2, spans[y].x2)
			}
		}
	}

	// Remove the appended coord
	coords = coords[:len(coords)-1]

	rectOK := func(x1, y1, x2, y2 int) bool {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			if spans[y] == nil {
				return false
			}
			sx1, sx2 := spans[y].x1, spans[y].x2
			if x1 < sx1 || x1 > sx2 || x2 < sx1 || x2 > sx2 {
				return false
			}
		}
		return true
	}

	area := 0
	for i := 0; i < len(coords); i++ {
		x1, y1 := coords[i].x, coords[i].y
		for j := i + 1; j < len(coords); j++ {
			x2, y2 := coords[j].x, coords[j].y
			if x1 != x2 && y1 != y2 {
				a := (abs(x1-x2) + 1) * (abs(y1-y2) + 1)
				if a > area && rectOK(x1, y1, x2, y2) {
					area = a
				}
			}
		}
	}
	fmt.Println(area)
}

func main() {

	// Part1
	coords, err := readCoords()
	if err != nil {
		log.Fatal(err)
	}

	area := 0
	for i := 0; i < len(coords); i++ {
		x1, y1 := coords[i].x, coords[i].y
		for j := i + 1; j < len(coords); j++ {
			x2, y2 := coords[j].x, coords[j].y
			if x1 != x2 && y1 != y2 {
				a := (abs(x1-x2) + 1) * (abs(y1-y2) + 1)
				area = max(area, a)
			}
		}
	}
	fmt.Println("Using two red tiles as opposite corners, what is the largest area of any rectangle you can make?")
	fmt.Println(area)

	// Part2
	maxX, maxY := 0, 0
	for _, c := range coords {
		maxX = max(maxX, c.x)
		maxY = max(maxY, c.y)
	}

	// Vertical sweep - build spans
	type span struct {
		x1, x2 int
	}
	spans := make([]*span, maxY+2)

	// Append first coord to create closed loop
	coords = append(coords, coords[0])

	for i := 1; i < len(coords); i++ {
		x1, y1 := coords[i-1].x, coords[i-1].y
		x2, y2 := coords[i].x, coords[i].y

		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		for y := y1; y <= y2; y++ {
			if spans[y] == nil {
				spans[y] = &span{x1, x2}
			} else {
				spans[y].x1 = min(x1, spans[y].x1)
				spans[y].x2 = max(x2, spans[y].x2)
			}
		}
	}

	// Remove the appended coord
	coords = coords[:len(coords)-1]

	rectOK := func(x1, y1, x2, y2 int) bool {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			if spans[y] == nil {
				return false
			}
			sx1, sx2 := spans[y].x1, spans[y].x2
			if x1 < sx1 || x1 > sx2 || x2 < sx1 || x2 > sx2 {
				return false
			}
		}
		return true
	}

	area = 0
	for i := 0; i < len(coords); i++ {
		x1, y1 := coords[i].x, coords[i].y
		for j := i + 1; j < len(coords); j++ {
			x2, y2 := coords[j].x, coords[j].y
			if x1 != x2 && y1 != y2 {
				a := (abs(x1-x2) + 1) * (abs(y1-y2) + 1)
				if a > area && rectOK(x1, y1, x2, y2) {
					area = a
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("Using two red tiles as opposite corners, what is the largest area of any rectangle you can make using only red and green tiles?")
	fmt.Println(area)

}
