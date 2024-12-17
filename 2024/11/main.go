package main

import (
	"fmt"
	"math"
	"strconv"
)

// solve computes the number of stones resulting from a starting set of stones
// and a number of blink iterations
func solve(n int, input []int) int {
	// Create initial stone count map
	stones := make(map[int]int)
	for _, i := range input {
		stones[i]++
	}

	// Perform n blink iterations
	for i := 0; i < n; i++ {
		stones = blinks(stones)
	}

	// Sum up total stones
	total := 0
	for _, count := range stones {
		total += count
	}
	return total
}

// blinks transforms all stones in the map according to their blink rule
func blinks(stones map[int]int) map[int]int {
	newStones := make(map[int]int)

	for stone, count := range stones {
		newStoneVals := blink(stone)
		for _, newStone := range newStoneVals {
			newStones[newStone] += count
		}
	}

	return newStones
}

// blink determines what stone(s) a given stone turns into
func blink(n int) []int {
	if n == 0 {
		return []int{1}
	}

	// Convert number to string to check length
	nStr := strconv.Itoa(n)
	w := len(nStr) / 2

	// If even length, split in half
	if w > 0 && len(nStr)%2 == 0 {
		// Calculate split point
		divisor := int(math.Pow10(w))
		l := n / divisor
		r := n % divisor
		return []int{l, r}
	}

	// Otherwise multiply by 2024
	return []int{n * 2024}
}

func main() {
	input := []int{5, 62914, 65, 972, 0, 805922, 6521, 1639064}

	fmt.Println("Part 1: How many stones will you have after blinking 25 times?")
	fmt.Println(solve(25, input))
	fmt.Println()
	fmt.Println("How many stones would you have after blinking a total of 75 times?")
	fmt.Println(solve(75, input))
}
