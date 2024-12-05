package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("puzzle.txt")
	sections := strings.Split(string(data), "\n\n")
	p1, p2 := sections[0], sections[1]

	orders := make(map[int][]int)
	for _, order := range strings.Split(p1, "\n") {
		parts := strings.Split(order, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		orders[before] = append(orders[before], after)
	}

	updates := parseUpdates(p2)

	part1, part2 := 0, 0
	for _, pages := range updates {
		sortedPages := make([]int, len(pages))
		copy(sortedPages, pages)
		sort.Slice(sortedPages, func(i, j int) bool {
			return len(filterOrderedPages(orders, sortedPages[i], sortedPages)) >
				len(filterOrderedPages(orders, sortedPages[j], sortedPages))
		})

		if arraysEqual(pages, sortedPages) {
			part1 += pages[len(pages)/2]
		} else {
			part2 += sortedPages[len(sortedPages)/2]
		}
	}

	fmt.Println("Part 1: What do you get if you add up the middle page number from those correctly-ordered updates?")
	fmt.Println(part1)
	fmt.Println()
	fmt.Println("Part 2: What do you get if you add up the middle page numbers after correctly ordering just those updates?")
	fmt.Println(part2)
}

func parseUpdates(p2 string) [][]int {
	var updates [][]int
	for _, line := range strings.Split(p2, "\n") {
		if line != "" {
			nums := strings.Split(line, ",")
			update := make([]int, len(nums))
			for i, num := range nums {
				update[i], _ = strconv.Atoi(num)
			}
			updates = append(updates, update)
		}
	}
	return updates
}

func filterOrderedPages(orders map[int][]int, page int, pages []int) []int {
	var ordered []int
	for _, order := range orders[page] {
		if contains(pages, order) {
			ordered = append(ordered, order)
		}
	}
	return ordered
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func arraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
