package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type cell struct {
	row, column int
}

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	var startCell cell
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
		if c := slices.Index(grid[len(grid)-1], '^'); c != -1 {
			startCell = cell{row: len(grid) - 1, column: c}
			grid[startCell.row][startCell.column] = '.'
		}
	}

	directions := []cell{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	maxRow := len(grid)
	maxCol := len(grid[0])

	visited := make(map[cell]bool)

	at := startCell
	facing := 0

	for isValidIndex(at, maxRow, maxCol) {
		if grid[at.row][at.column] == '#' {
			at.row -= directions[facing].row
			at.column -= directions[facing].column
			facing = (facing + 1) % len(directions)
			continue
		}
		visited[at] = true
		grid[at.row][at.column] = 'X'
		at.row += directions[facing].row
		at.column += directions[facing].column
	}

	visitedIndices := make([]cell, 0, len(visited))

	for idx := range visited {
		visitedIndices = append(visitedIndices, idx)
	}

	res := len(visitedIndices)
	fmt.Println("Part One: How many distinct positions will the guard visit before leaving the mapped area?")
	fmt.Println(res)

	// Part 2
	cycleCount := 0
	for _, index := range visitedIndices {
		if (index == startCell) || grid[index.row][index.column] == '#' {
			continue
		}
		grid[index.row][index.column] = '#'
		if hasCycle(grid, startCell) {
			cycleCount++
		}
		grid[index.row][index.column] = '.'
	}
	fmt.Println()
	fmt.Println("Part Two: How many different positions could you choose for this obstruction?")
	fmt.Println(cycleCount)
}

func hasCycle(grid [][]rune, startIndex cell) bool {
	visited := make(map[cell]cell)

	directions := []cell{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	maxRow := len(grid)
	maxCol := len(grid[0])

	at := startIndex
	facing := 0

	for isValidIndex(at, maxRow, maxCol) {
		if visited[at] == directions[facing] {
			return true
		}

		visited[at] = directions[facing]

		if grid[at.row][at.column] == '#' {
			at.row -= directions[facing].row
			at.column -= directions[facing].column
			facing = (facing + 1) % len(directions)
			continue
		}

		grid[at.row][at.column] = 'X'
		at.row += directions[facing].row
		at.column += directions[facing].column
	}

	return false
}

func isValidIndex(idx cell, maxRow, maxCol int) bool {
	return idx.row >= 0 && idx.column >= 0 && idx.row < maxRow && idx.column < maxCol
}
