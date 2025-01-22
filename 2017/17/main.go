package main

import "fmt"

func main() {
	const step = 328 // input

	// Part 1
	buf := []int{0}
	i := 0

	for t := 1; t <= 2017; t++ {
		i = (i+step)%len(buf) + 1
		// Insert at position i
		buf = append(buf[:i], append([]int{t}, buf[i:]...)...)
	}

	// Print surrounding values
	start := max(0, i-5)
	end := min(len(buf), i+5)
	fmt.Println("Part 1: What is the value after 2017 in your completed circular buffer?")
	fmt.Println(buf[start:end])

	// Part 2
	// We don't need to maintain the actual buffer, just track position
	// and value after 0
	i = 0
	valAfter0 := 0

	for t := 1; t <= 50_000_000; t++ {
		i = (i+step)%t + 1
		if i == 1 {
			valAfter0 = t
		}
	}

	fmt.Println()
	fmt.Println("Part 2: What is the value after 0 the moment 50000000 is inserted?")
	fmt.Println(valAfter0)
}

// Helper functions for min/max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
