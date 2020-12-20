package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	wires := strings.Fields(string(data))
	seen := make([]map[image.Point]struct{}, len(wires))
	min := 0

	for i, w := range wires {
		x, y := 0, 0
		seen[i] = map[image.Point]struct{}{}

		for _, s := range strings.Split(w, ",") {
			for j, _ := strconv.Atoi(s[1:]); j > 0; j-- {
				d := map[byte]image.Point{'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0}}[s[0]]
				x, y = x+d.X, y+d.Y
				seen[i][image.Point{x, y}] = struct{}{}
			}
		}
	}

	for p := range seen[1] {
		if _, ok := seen[0][p]; ok {
			dist := int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
			if min == 0 || dist < min {
				min = dist
			}
		}
	}

	fmt.Println(min)

	fmt.Println()
	fmt.Println("Part2")

	// Reset vars
	seen2 := make([]map[image.Point]int, len(wires))
	min = 0

	for i, w := range wires {
		x, y := 0, 0
		steps := 0
		seen2[i] = map[image.Point]int{}

		for _, s := range strings.Split(w, ",") {
			for j, _ := strconv.Atoi(s[1:]); j > 0; j-- {
				d := map[byte]image.Point{'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0}}[s[0]]
				x, y = x+d.X, y+d.Y
				steps++
				seen2[i][image.Point{x, y}] = steps
			}
		}
	}

	for p := range seen2[1] {
		if _, ok := seen2[0][p]; ok {
			steps := seen2[0][p] + seen2[1][p]
			if min == 0 || steps < min {
				min = steps
			}
		}
	}

	fmt.Println(min)
}
