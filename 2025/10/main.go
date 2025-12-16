package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type State struct {
	values []int
}

func (s State) String() string {
	return fmt.Sprint(s.values)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p1, p2 := 0, 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		// Parse diagram
		dia := parseDiagram(parts[0])

		// Parse buttons
		buttons := make([][]int, 0)
		for i := 1; i < len(parts)-1; i++ {
			buttons = append(buttons, parseList(parts[i]))
		}

		// Parse jolts
		jolts := parseList(parts[len(parts)-1])

		// Generate all possible button combinations
		ops := make(map[string][]int)
		patterns := make(map[string][][]int)

		numButtons := len(buttons)
		numCombos := 1 << numButtons // 2^numButtons

		for combo := 0; combo < numCombos; combo++ {
			pressed := make([]int, numButtons)
			jolt := make([]int, len(jolts))

			// Determine which buttons are pressed
			for i := 0; i < numButtons; i++ {
				if combo&(1<<i) != 0 {
					pressed[i] = 1
				}
			}

			// Calculate jolt values
			for i, p := range pressed {
				if p == 1 {
					for _, j := range buttons[i] {
						jolt[j]++
					}
				}
			}

			// Calculate light pattern (mod 2)
			lights := make([]int, len(jolt))
			for i, v := range jolt {
				lights[i] = v % 2
			}

			lightsKey := intSliceKey(lights)
			pressedKey := intSliceKey(pressed)

			ops[pressedKey] = jolt
			patterns[lightsKey] = append(patterns[lightsKey], pressed)
		}

		// Part 1: Find minimum button presses for target pattern
		diaKey := intSliceKey(dia)
		minPresses := math.MaxInt
		for _, pressed := range patterns[diaKey] {
			sum := 0
			for _, p := range pressed {
				sum += p
			}
			if sum < minPresses {
				minPresses = sum
			}
		}
		p1 += minPresses

		// Part 2: Recursive search with memoization
		cache := make(map[string]int)
		p2 += presses(jolts, patterns, ops, cache)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func parseDiagram(s string) []int {
	s = s[1 : len(s)-1] // Remove brackets
	result := make([]int, len(s))
	for i, c := range s {
		if c == '#' {
			result[i] = 1
		}
	}
	return result
}

func parseList(s string) []int {
	s = s[1 : len(s)-1] // Remove brackets
	if s == "" {
		return []int{}
	}
	parts := strings.Split(s, ",")
	result := make([]int, len(parts))
	for i, p := range parts {
		result[i], _ = strconv.Atoi(p)
	}
	return result
}

func intSliceKey(slice []int) string {
	strs := make([]string, len(slice))
	for i, v := range slice {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, ",")
}

func presses(target []int, patterns map[string][][]int, ops map[string][]int, cache map[string]int) int {
	// Check if all zeros
	allZero := true
	for _, v := range target {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	// Check if any negative
	for _, v := range target {
		if v < 0 {
			return math.MaxInt / 2 // Use half max to avoid overflow
		}
	}

	// Check cache
	key := intSliceKey(target)
	if val, ok := cache[key]; ok {
		return val
	}

	// Calculate lights pattern
	lights := make([]int, len(target))
	for i, v := range target {
		lights[i] = v % 2
	}
	lightsKey := intSliceKey(lights)

	total := math.MaxInt / 2
	for _, pressed := range patterns[lightsKey] {
		pressedKey := intSliceKey(pressed)
		diff := ops[pressedKey]

		// Calculate new target
		newTarget := make([]int, len(target))
		for i := range target {
			newTarget[i] = (target[i] - diff[i]) / 2
		}

		// Sum pressed buttons
		sum := 0
		for _, p := range pressed {
			sum += p
		}

		result := sum + 2*presses(newTarget, patterns, ops, cache)
		if result < total {
			total = result
		}
	}

	cache[key] = total
	return total
}
