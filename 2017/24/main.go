package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Component [2]int

var (
	strongestBridge      []Component
	strongestBridgeScore int
	longestBridge        []Component
	longestBridgeScore   int
)

func pickNextPort(a, b Component) int {
	if b[0] == a[0] || b[0] == a[1] {
		return b[1]
	}
	return b[0]
}

func scoreBridge(bridge []Component) int {
	score := 0
	for _, comp := range bridge {
		score += comp[0] + comp[1]
	}
	return score
}

func check(comps []Component, bridge []Component) {
	var nextPort int

	switch len(bridge) {
	case 0:
		nextPort = 0
	case 1:
		nextPort = pickNextPort(Component{0, 0}, bridge[0])
	default:
		nextPort = pickNextPort(bridge[len(bridge)-2], bridge[len(bridge)-1])
	}

	foundABridge := false

	for i, comp := range comps {
		if comp[0] == nextPort || comp[1] == nextPort {
			foundABridge = true

			// Create next bridge by copying and appending
			nextBridge := make([]Component, len(bridge))
			copy(nextBridge, bridge)
			nextBridge = append(nextBridge, comp)

			// Create next components by removing current component
			nextComps := make([]Component, 0, len(comps)-1)
			nextComps = append(nextComps, comps[:i]...)
			nextComps = append(nextComps, comps[i+1:]...)

			check(nextComps, nextBridge)
		}
	}

	if !foundABridge {
		score := scoreBridge(bridge)

		if score > strongestBridgeScore {
			strongestBridge = make([]Component, len(bridge))
			copy(strongestBridge, bridge)
			strongestBridgeScore = score
		}

		if len(bridge) >= len(longestBridge) {
			if score > longestBridgeScore {
				longestBridge = make([]Component, len(bridge))
				copy(longestBridge, bridge)
				longestBridgeScore = score
			}
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var allComps []Component
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "/")
		if len(parts) != 2 {
			continue
		}

		a, err1 := strconv.Atoi(parts[0])
		b, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		allComps = append(allComps, Component{a, b})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	check(allComps, []Component{})

	fmt.Println("Part 1: What is the strength of the strongest bridge you can make with the components you have available?")
	fmt.Println(strongestBridgeScore)

	fmt.Println()
	fmt.Println("Part 2: What is the strength of the longest bridge you can make? ")
	fmt.Println(longestBridgeScore)
}
