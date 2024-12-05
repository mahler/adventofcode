package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := make(map[struct{ x, y int }]rune)
	var xh, yh int

	data, _ := os.Open("puzzle.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, c := range line {
			grid[struct{ x, y int }{x, y}] = c
			xh = max(xh, x)
			yh = max(yh, y)
		}
	}

	part1 := 0
	for y := 0; y <= yh; y++ {
		for x := 0; x <= xh; x++ {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					word := getWord(grid, x, y, dx, dy)
					if word == "XMAS" {
						part1++
					}
				}
			}
		}
	}
	fmt.Println("Part1: How many times does XMAS appear?")
	fmt.Println(part1)

	part2 := 0
	for y := 0; y <= yh; y++ {
		for x := 0; x <= xh; x++ {
			diag1 := []rune{
				grid[struct{ x, y int }{x - 1, y - 1}],
				grid[struct{ x, y int }{x, y}],
				grid[struct{ x, y int }{x + 1, y + 1}],
			}
			diag2 := []rune{
				grid[struct{ x, y int }{x - 1, y + 1}],
				grid[struct{ x, y int }{x, y}],
				grid[struct{ x, y int }{x + 1, y - 1}],
			}
			word1 := string(diag1)
			word2 := string(diag2)
			if (word1 == "MAS" || word1 == "SAM") &&
				(word2 == "MAS" || word2 == "SAM") {
				part2++
			}
		}
	}
	fmt.Println()
	fmt.Println("Part 2: How many times does an X-MAS appear?")
	fmt.Println(part2)
}

func getWord(grid map[struct{ x, y int }]rune, x, y, dx, dy int) string {
	word := []rune{}
	for n := 0; n < 4; n++ {
		word = append(word, grid[struct{ x, y int }{x + dx*n, y + dy*n}])
	}
	return string(word)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
