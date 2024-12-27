package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hasNonZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []int
		for _, num := range strings.Fields(scanner.Text()) {
			n, _ := strconv.Atoi(num)
			line = append(line, n)
		}
		lines = append(lines, line)
	}

	r1, r2 := 0, 0
	for _, line := range lines {
		h := [][]int{line}
		for hasNonZero(h[len(h)-1]) {
			diff := make([]int, len(h[len(h)-1])-1)
			for i := 0; i < len(h[len(h)-1])-1; i++ {
				diff[i] = h[len(h)-1][i+1] - h[len(h)-1][i]
			}
			h = append(h, diff)
		}

		for i := range h {
			r1 += h[i][len(h[i])-1]
			if i%2 == 0 {
				r2 += h[i][0]
			} else {
				r2 -= h[i][0]
			}
		}
	}
	fmt.Println("Part 1: Analyze your OASIS report and extrapolate the next value for each history.")
	fmt.Println("What is the sum of these extrapolated values?")
	fmt.Println(r1)
	fmt.Println()
	fmt.Println("Part 2: Analyze your OASIS report again, this time extrapolating the previous value for each history.")
	fmt.Println("What is the sum of these extrapolated values?")
	fmt.Println(r2)
}
