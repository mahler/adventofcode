package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	scanner.Scan()
	input := scanner.Text()

	// Split input into numbers
	strNums := strings.Fields(input)
	banks := make([]int, len(strNums))
	for i, str := range strNums {
		num, _ := strconv.Atoi(str)
		banks[i] = num
	}

	// Track seen states using string representation as key
	seen := make(map[string]int)
	cycles := 0

	for {
		// Convert current state to string for map key
		state := arrayToString(banks)

		// Check if we've seen this state before
		if firstSeen, exists := seen[state]; exists {

			fmt.Println("Part 1: Given the initial block counts in your puzzle input, how many redistribution")
			fmt.Println("cycles must be completed before a configuration is produced that has been seen before?")
			fmt.Println(cycles)

			fmt.Println()
			fmt.Println("Part 2: How many cycles are in the infinite loop that arises from the configuration in your puzzle input?")
			fmt.Println(cycles - firstSeen)
			break
		}

		// Record current state
		seen[state] = cycles
		cycles++

		// Find bank with most blocks
		maxIdx := 0
		for i := 1; i < len(banks); i++ {
			if banks[i] > banks[maxIdx] {
				maxIdx = i
			}
		}

		// Redistribute blocks
		blocks := banks[maxIdx]
		banks[maxIdx] = 0
		for i := 0; i < blocks; i++ {
			banks[(maxIdx+1+i)%len(banks)]++
		}
	}
}

// Helper function to convert array to string representation
func arrayToString(arr []int) string {
	strs := make([]string, len(arr))
	for i, num := range arr {
		strs[i] = strconv.Itoa(num)
	}
	return strings.Join(strs, " ")
}
