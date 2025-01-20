package main

import (
	"fmt"
)

func reverse(numbers []int, start, length int) {
	for pos := 0; pos < length/2; pos++ {
		a := (start + pos) % 256
		b := (start + length - 1 - pos) % 256
		numbers[a], numbers[b] = numbers[b], numbers[a]
	}
}

func writeGrid(g []int, row, col, hashchar int) {
	offset := 128*row + col*8
	for i := 0; i < 8; i++ {
		if (hashchar & (1 << (7 - i))) != 0 {
			g[offset+i] = -1
		} else {
			g[offset+i] = 0
		}
	}
}

func floodFill(g []int, i, c int) {
	g[i] = c
	// Check up
	if i/128 > 0 && g[i-128] == -1 {
		floodFill(g, i-128, c)
	}
	// Check down
	if i/128 < 127 && g[i+128] == -1 {
		floodFill(g, i+128, c)
	}
	// Check left
	if i%128 > 0 && g[i-1] == -1 {
		floodFill(g, i-1, c)
	}
	// Check right
	if i%128 < 127 && g[i+1] == -1 {
		floodFill(g, i+1, c)
	}
}

func countBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func main() {
	seed := "hfdlxzhv"
	trailing := []int{17, 31, 73, 47, 23}
	grid := make([]int, 128*128)
	bitcount := 0

	// Part 1
	for row := 0; row < 128; row++ {
		// Initialize numbers 0-255
		numbers := make([]int, 256)
		for i := range numbers {
			numbers[i] = i
		}

		// Create input string and convert to lengths
		input := fmt.Sprintf("%s-%d", seed, row)
		lengths := make([]int, len(input)+len(trailing))
		for i, ch := range input {
			lengths[i] = int(ch)
		}
		copy(lengths[len(input):], trailing)

		// Perform 64 rounds of knot hash
		start := 0
		skip := 0
		for r := 0; r < 64; r++ {
			for l := 0; l < len(lengths); l++ {
				reverse(numbers, start%256, lengths[l])
				start += lengths[l] + skip
				skip++
			}
		}

		// Calculate dense hash and update bitcount
		for i := 0; i < 256; i++ {
			if i%16 == 0 {
				hashchar := numbers[i]
				if i%16 == 15 {
					bitcount += countBits(hashchar)
					writeGrid(grid, row, i/16, hashchar)
				}
			} else {
				if i%16 == 15 {
					hashchar := 0
					for j := i - 15; j <= i; j++ {
						hashchar ^= numbers[j]
					}
					bitcount += countBits(hashchar)
					writeGrid(grid, row, i/16, hashchar)
				}
			}
		}
	}
	fmt.Println("Part 1: Given your actual key string, how many squares are used?")
	fmt.Println(bitcount)

	// Part 2: Count regions using flood fill
	regions := 0
	for i := 0; i < 128*128; i++ {
		if grid[i] == -1 {
			regions++
			floodFill(grid, i, regions)
		}
	}
	fmt.Println()
	fmt.Println("Part 2: How many regions are present given your key string?")
	fmt.Println(regions)
}
