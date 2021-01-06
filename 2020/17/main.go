package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type point struct {
	X, Y, Z, W int
}

func (p point) String() string {
	return fmt.Sprintf("%d/%d/%d", p.X, p.Y, p.Z)
}

func main() {
	activePoints := make(map[point]bool, 40)
	data, err := ioutil.ReadFile("data.set")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}
	dataInput := strings.Split(strings.TrimSpace(string(data)), "\n")

	x := 0
	for _, dataRow := range dataInput {
		for i, c := range dataRow {
			if c == '#' {
				activePoints[point{x, i, 0, 0}] = true
			}
		}
		x++
		if err != nil {
			log.Fatal(err)
		}
	}

	// save copy for part 2
	p2Points := activePoints

	fmt.Println()
	fmt.Println("Day 17 - Part 1: Conway Cubes")
	for i := 0; i < 6; i++ {
		activePoints = apply(activePoints, false)
	}
	fmt.Println("Active Points:", len(activePoints))

	fmt.Println()
	fmt.Println("Part 2: Cubes 4D")

	for i := 0; i < 6; i++ {
		p2Points = apply(p2Points, true)
	}
	fmt.Println("Active Points:", len(p2Points))
}

func process(activePoints map[point]bool, n int, fourD bool) int {
	for i := 0; i < n; i++ {
		activePoints = apply(activePoints, fourD)
	}

	return len(activePoints)
}

// Applied utlized from https://github.com/thlacroix/goadvent/blob/master/2020/day17/main.go
func apply(activePoints map[point]bool, fourD bool) map[point]bool {
	l := 26
	if fourD {
		l = 80
	}
	affectedPoints := make(map[point]int, l)

	for p := range activePoints {
		for x := p.X - 1; x <= p.X+1; x++ {
			for y := p.Y - 1; y <= p.Y+1; y++ {
				for z := p.Z - 1; z <= p.Z+1; z++ {
					if fourD {
						for w := p.W - 1; w <= p.W+1; w++ {
							np := point{x, y, z, w}
							if np == p {
								continue
							}
							affectedPoints[np]++
						}
					} else {
						np := point{x, y, z, 0}
						if np == p {
							continue
						}
						affectedPoints[np]++
					}

				}
			}
		}
	}

	newActives := make(map[point]bool, len(activePoints))

	for p, c := range affectedPoints {
		if activePoints[p] {
			if c == 2 || c == 3 {
				newActives[p] = true
			}
		} else {
			if c == 3 {
				newActives[p] = true
			}
		}
	}
	return newActives
}
