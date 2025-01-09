package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(part int, items [][]int, operations []string, tests []int, conditions [][]int, modulo int) int {
	monkeyItems := make([][]int, len(items))
	for i := range items {
		monkeyItems[i] = make([]int, len(items[i]))
		copy(monkeyItems[i], items[i])
	}

	rounds := 20
	if part == 2 {
		rounds = 10000
	}

	inspections := make([]int, len(monkeyItems))

	for round := 0; round < rounds; round++ {
		for i := range monkeyItems {
			for _, item := range monkeyItems[i] {
				current := item

				if operations[i] == "* old" {
					current *= current
				} else if strings.HasPrefix(operations[i], "* ") {
					val, _ := strconv.Atoi(operations[i][2:])
					current *= val
				} else if strings.HasPrefix(operations[i], "+ ") {
					val, _ := strconv.Atoi(operations[i][2:])
					current += val
				}

				if part == 1 {
					current /= 3
				} else {
					current %= modulo
				}

				if current%tests[i] == 0 {
					monkeyItems[conditions[i][0]] = append(monkeyItems[conditions[i][0]], current)
				} else {
					monkeyItems[conditions[i][1]] = append(monkeyItems[conditions[i][1]], current)
				}
				inspections[i]++
			}
			monkeyItems[i] = nil
		}
	}

	largest := 0
	secondLargest := 0
	for _, count := range inspections {
		if count > largest {
			secondLargest = largest
			largest = count
		} else if count > secondLargest {
			secondLargest = count
		}
	}

	return largest * secondLargest
}

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	var monkeyOperations []string
	var monkeyTest []int
	var monkeyConditions [][]int
	var monkeyItems [][]int

	modulo := 1
	for i := 1; i < len(lines); i += 7 {
		items := strings.Split(strings.TrimSpace(lines[i][18:]), ", ")
		itemInts := make([]int, len(items))
		for j, item := range items {
			itemInts[j], _ = strconv.Atoi(item)
		}
		monkeyItems = append(monkeyItems, itemInts)

		monkeyOperations = append(monkeyOperations, strings.TrimSpace(lines[i+1][23:]))

		test, _ := strconv.Atoi(strings.TrimSpace(lines[i+2][21:]))
		monkeyTest = append(monkeyTest, test)
		modulo *= test

		trueCase, _ := strconv.Atoi(strings.TrimSpace(lines[i+3][29:]))
		falseCase, _ := strconv.Atoi(strings.TrimSpace(lines[i+4][30:]))
		monkeyConditions = append(monkeyConditions, []int{trueCase, falseCase})
	}

	part1result := calculate(1, monkeyItems, monkeyOperations, monkeyTest, monkeyConditions, modulo)
	fmt.Println("Part 1: What is the level of monkey business after 20 rounds of stuff-slinging simian shenanigans?")
	fmt.Println(part1result)

	part2result := calculate(2, monkeyItems, monkeyOperations, monkeyTest, monkeyConditions, modulo)
	fmt.Println()
	fmt.Println("Part 2: Starting again from the initial state in your puzzle input,")
	fmt.Println("what is the level of monkey business after 10000 rounds?")
	fmt.Println(part2result)
}
