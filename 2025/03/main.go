package main

import (
	"bufio"
	"fmt"
	"os"
)

func num(m int, ns []int) int {
	result := m
	for _, n := range ns {
		result *= 10
		result += n
	}
	return result
}

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, n := range nums[1:] {
		if n > maxVal {
			maxVal = n
		}
	}
	return maxVal
}

func indexOf(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	p1, p2 := 0, 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		xs := make([]int, 0, len(line))
		for _, c := range line {
			if c >= '0' && c <= '9' {
				xs = append(xs, int(c-'0'))
			}
		}

		// Part 1
		m := max(xs[:len(xs)-1])
		mi := indexOf(xs, m)
		p1 += m*10 + max(xs[mi+1:])

		// Part 2
		m = max(xs[:len(xs)-11])
		mi = indexOf(xs, m)
		ns := make([]int, 11)
		copy(ns, xs[mi+1:mi+12])

		for _, x := range xs[mi+12:] {
			n := num(m, ns)
			for i := 0; i < 11; i++ {
				// Create ns2: ns[:i] + ns[i+1:] + [x]
				ns2 := make([]int, 0, 11)
				ns2 = append(ns2, ns[:i]...)
				ns2 = append(ns2, ns[i+1:]...)
				ns2 = append(ns2, x)

				if n < num(m, ns2) {
					ns = ns2
					break
				}
			}
		}
		p2 += num(m, ns)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("what is the total output joltage?")
	fmt.Println(p1)
	fmt.Println()
	fmt.Println("Part 2: What is the new total output joltage?")
	fmt.Println(p2
        )

}
