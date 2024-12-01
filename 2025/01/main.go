package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var part1list, part2list []int
	counts2 := map[int]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		var n1, n2 int
		fmt.Sscanf(s, "%d   %d", &n1, &n2)
		part1list, part2list = append(part1list, n1), append(part2list, n2)
		counts2[n2]++
	}

	slices.Sort(part1list)
	slices.Sort(part2list)

	part1, part2 := 0, 0
	for i := range part1list {
		part1 += absoluteValue(part2list[i] - part1list[i])
		part2 += part1list[i] * counts2[part1list[i]]
	}

	fmt.Println("What is the total distance between your lists?")
	fmt.Println(part1)

	fmt.Println()
	fmt.Println("Part 2: Once again consider your left and right lists. What is their similarity score?")

	fmt.Println(part2)
}

func absoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
