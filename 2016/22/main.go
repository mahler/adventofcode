package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numRe := regexp.MustCompile(`\d+`)

	nodes := 0
	maxX, maxY := 0, 0
	big := make(map[Point]bool)
	var empty Point

	// Skip first two lines
	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {
		nums := numRe.FindAllString(scanner.Text(), -1)
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		size, _ := strconv.Atoi(nums[2])
		used, _ := strconv.Atoi(nums[3])

		nodes++
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if size > 500 {
			big[Point{x, y}] = true
		}
		if used == 0 {
			empty = Point{x, y}
		}
	}

	fmt.Println("part 1: How many viable pairs of nodes are there?")
	fmt.Println(nodes - len(big) - 1)

	queue := []Point{empty}
	steps := 0
	seen := make(map[Point]bool)
	seen[empty] = true

	dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		var nextQueue []Point

		for _, p := range queue {
			if p.x == maxX && p.y == 0 {
				fmt.Println()
				fmt.Println("Part 2: What is the fewest number of steps required to move your goal data to node-x0-y0?")
				fmt.Println(steps + 5*(maxX-1))
				return
			}

			for _, d := range dirs {
				nx := p.x + d.x
				ny := p.y + d.y
				np := Point{nx, ny}

				if nx < 0 || nx > maxX || ny < 0 || ny > maxY || big[np] {
					continue
				}
				if seen[np] {
					continue
				}

				seen[np] = true
				nextQueue = append(nextQueue, np)
			}
		}

		queue = nextQueue
		steps++
	}
}
