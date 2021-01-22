package main

import (
	"fmt"
)

// Row represents every tile
type Row []bool

func (r *Row) getSafeCount() int {
	sum := 0
	for i := 1; i < len(*r)-1; i++ {
		if !(*r)[i] {
			sum++
		}
	}
	return sum
}

func main() {
	input := "^^^^......^...^..^....^^^.^^^.^.^^^^^^..^...^^...^^^.^^....^..^^^.^.^^...^.^...^^.^^^.^^^^.^^.^..^.^"

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 18, Part 1: Like a Rogue")
	p1Result := solve(input, 40)
	fmt.Println("Starting with the map in your puzzle input, in a total of 40 rows (including the starting row),")
	fmt.Println("how many safe tiles are there?")
	fmt.Println(p1Result)

	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("How many safe tiles are there in a total of 400000 rows?")
	p2Result := solve(input, 400000)
	fmt.Println(p2Result)
}

func solve(input string, size int) int {
	firstRow := Row(make([]bool, len(input)+2))
	firstRow[0], firstRow[len(firstRow)-1] = false, false
	for i := 0; i < len(input); i++ {
		if input[i] == '^' {
			firstRow[i+1] = true
		} else {
			firstRow[i+1] = false
		}
	}

	rows := make([]Row, size)
	rows[0] = firstRow
	width := len(firstRow)
	for i := 1; i < size; i++ {
		r := make([]bool, width)
		r[0], r[width-1] = false, false
		for j := 1; j < width-1; j++ {
			r[j] = rows[i-1][j-1] && rows[i-1][j] && !rows[i-1][j+1] ||
				!rows[i-1][j-1] && rows[i-1][j] && rows[i-1][j+1] ||
				rows[i-1][j-1] && !rows[i-1][j] && !rows[i-1][j+1] ||
				!rows[i-1][j-1] && !rows[i-1][j] && rows[i-1][j+1]
		}
		rows[i] = Row(r)
	}

	sum := 0
	for i := 0; i < size; i++ {
		sum = sum + rows[i].getSafeCount()
	}
	return sum
}
