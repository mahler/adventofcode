package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type group struct {
	c  []int
	qe int64
}

type groupList []group

func main() {
	// Read instructions
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	var weights []int

	for _, s := range fileLines {
		if s != "" {
			w, _ := strconv.Atoi(s)
			weights = append(weights, w)
		}
	}

	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 24, Part 1: It Hangs in the Balance")
	fmt.Println("What is the quantum entanglement of the first 3 group of packages in the ideal configuration?")
	qe3 := findLowest(weights, 3)
	fmt.Println(qe3)

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("Now, what is the quantum entanglement of the first 4 group of ackages in the ideal configuration?")
	qe4 := findLowest(weights, 4)
	fmt.Println(qe4)
}

func combinationsQE(n int, r int, weights []int, target int) <-chan group {
	c := make(chan group)
	go func() {
		defer close(c)

		indices := make([]int, r)
		for i := range indices {
			indices[i] = i
		}

		cw := 0
		qe := int64(1)
		for _, w := range indices {
			cw += weights[w]
			qe *= int64(weights[w])
		}
		if cw == target {
			out := make([]int, len(indices))
			copy(out, indices)
			c <- group{out, qe}
		}

		for n > 0 {
			i := r - 1
			for ; i >= 0; i-- {
				if indices[i] != i+n-r {
					break
				}
			}
			if i < 0 {
				break
			}

			indices[i]++
			for j := i + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}

			cw = 0
			qe = 1
			for _, w := range indices {
				cw += weights[w]
				qe *= int64(weights[w])
			}
			if cw == target {
				out := make([]int, len(indices))
				copy(out, indices)
				c <- group{out, qe}
			}
		}
	}()

	return c
}

func findLowest(weights []int, groups int) int64 {
	totalWeight := 0
	for _, v := range weights {
		totalWeight += v
	}

	targetWeight := totalWeight / groups

	for l := 1; l < len(weights)-(groups-2); l++ {
		for c := range combinationsQE(len(weights), l, weights, targetWeight) {
			return c.qe
		}
	}

	return 0
}
