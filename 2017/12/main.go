package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reachableFrom(start int, net [][]int) map[int]struct{} {
	reachable := make(map[int]struct{})
	reachable[start] = struct{}{}
	done := false

	for !done {
		frontier := []int{}
		for r := range reachable {
			for _, n := range net[r] {
				if _, exists := reachable[n]; !exists {
					frontier = append(frontier, n)
				}
			}
		}

		if len(frontier) == 0 {
			done = true
		} else {
			for _, n := range frontier {
				reachable[n] = struct{}{}
			}
		}
	}

	return reachable
}

func main() {
	net := [][]int{}
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " <-> ")
		if len(parts) != 2 {
			panic("invalid input format")
		}

		neighbors := []int{}
		for _, n := range strings.Split(parts[1], ", ") {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			neighbors = append(neighbors, num)
		}

		net = append(net, neighbors)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part 1: How many programs are in the group that contains program ID 0?")
	fmt.Println(len(reachableFrom(0, net)))

	// Part 2
	comps := make(map[int]struct{})
	for i := range net {
		comps[i] = struct{}{}
	}

	count := 0
	for len(comps) > 0 {
		var start int
		for c := range comps {
			start = c
			break
		}

		reachable := reachableFrom(start, net)
		for r := range reachable {
			delete(comps, r)
		}
		count++
	}
	fmt.Println()
	fmt.Println("Part 2: How many groups are there in total?")
	fmt.Println(count)
}
