package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	dx, dy int
}

var (
	UP    = Direction{-1, 0}
	DOWN  = Direction{1, 0}
	LEFT  = Direction{0, -1}
	RIGHT = Direction{0, 1}
)

var directions = []Direction{UP, DOWN, LEFT, RIGHT}

func main() {
	// Read input
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	// Create grid
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// Find start position
	var pos Direction
	for j, ch := range grid[0] {
		if ch == '|' {
			pos = Direction{0, j}
			break
		}
	}

	dir := DOWN
	seen := []rune{}
	steps := 0

	for {
		// Move
		pos.dx += dir.dx
		pos.dy += dir.dy
		steps++

		// Check if out of bounds
		if pos.dx < 0 || pos.dx >= len(grid) || pos.dy < 0 || pos.dy >= len(grid[pos.dx]) {
			break
		}

		curr := grid[pos.dx][pos.dy]

		switch {
		case curr == '|' || curr == '-':
			continue
		case curr == '+':
			// Find new direction
			for _, check := range directions {
				if check.dx == -dir.dx && check.dy == -dir.dy {
					continue
				}

				newX := pos.dx + check.dx
				newY := pos.dy + check.dy

				if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[newX]) {
					newCurr := grid[newX][newY]
					if newCurr == '|' || newCurr == '-' {
						pos.dx = newX
						pos.dy = newY
						dir = check
						steps++
						break
					}
				}
			}
		case curr >= 'A' && curr <= 'Z':
			seen = append(seen, curr)
		case curr == ' ':
			goto end
		}
	}

end:

	fmt.Println("Part 1: What letters will it see (in the order it would see them) if it follows the path?")
	fmt.Println(string(seen))

	fmt.Println()
	fmt.Println("Part 2: How many steps does the packet need to go?")
	fmt.Println(steps)
}
