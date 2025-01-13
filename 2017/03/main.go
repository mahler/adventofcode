package main

import (
	"fmt"
	"math"
)

func main() {
	input := 347991

	fmt.Println("Part 1: How many steps are required to carry the data from the square identified in your puzzle input all the way to the access port?")
	fmt.Println(part1(input))

	fmt.Println()
	fmt.Println("Part 2: What is the first value written that is larger than your puzzle input?")
	fmt.Println(part2(input))
}

func part1(input int) int {
	x := 0
	y := 0
	n := 1
	d := 1
	for {
		y += d
		n += d
		if n >= input {
			y -= n - input
			break
		}
		x -= d
		n += d
		if n >= input {
			x += n - input
			break
		}
		d += 1
		y -= d
		n += d
		if n >= input {
			y += n - input
			break
		}
		x += d
		n += d
		if n >= input {
			x -= n - input
			break
		}
		d += 1
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func part2(input int) int {
	grid := make([][]int, 1024)
	for i := range grid {
		grid[i] = make([]int, 1024)
	}
	x := 512
	y := 512
	grid[x][y] = 1
	k := 1
	for {
		for i := 0; i < k; i++ {
			y++
			r := fill(grid, x, y)
			if r > input {
				return r
			}
			grid[x][y] = r
		}
		for i := 0; i < k; i++ {
			x--
			r := fill(grid, x, y)
			if r > input {
				return r
			}
			grid[x][y] = r
		}
		k++
		for i := 0; i < k; i++ {
			y--
			r := fill(grid, x, y)
			if r > input {
				return r
			}
			grid[x][y] = r
		}
		for i := 0; i < k; i++ {
			x++
			r := fill(grid, x, y)
			if r > input {
				return r
			}
			grid[x][y] = r
		}
		k++
	}
}

func fill(grid [][]int, x, y int) int {
	return grid[x-1][y-1] + grid[x][y-1] + grid[x+1][y-1] + grid[x-1][y] +
		grid[x+1][y] + grid[x-1][y+1] + grid[x][y+1] + grid[x+1][y+1]
}
