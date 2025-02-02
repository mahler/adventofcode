package main

import (
	"fmt"
	"sort"
)

const (
	caveSystemDepth = 3066
	tx              = 13
	ty              = 726
)

// Cache for erosion values
var erosionCache = make(map[[2]int]int)

// Region types
const (
	rock   = 0
	water  = 1
	narrow = 2
)

// Region symbols
var regionSymbols = map[int]string{
	rock:   ".",
	water:  "=",
	narrow: "|",
}

type State struct {
	x, y int
	item string
}

type PathNode struct {
	state     State
	action    string
	totalCost int
}

type Path []PathNode

// Erosion calculates the erosion level with memoization
func erosion(x, y int) int {
	pos := [2]int{x, y}
	if val, exists := erosionCache[pos]; exists {
		return val
	}

	result := (gindex(x, y) + caveSystemDepth) % 20183
	erosionCache[pos] = result
	return result
}

// Geological index calculation
func gindex(x, y int) int {
	if (x == 0 && y == 0) || (x == tx && y == ty) {
		return 0
	}
	if y == 0 {
		return x * 16807
	}
	if x == 0 {
		return y * 48271
	}
	return erosion(x-1, y) * erosion(x, y-1)
}

// Check if item is valid for region
func itemValidForRegion(item string, region int) bool {
	switch region {
	case rock:
		return item == "Torch" || item == "Climb"
	case water:
		return item == "Nothing" || item == "Climb"
	case narrow:
		return item == "Torch" || item == "Nothing"
	}
	return false
}

// Get successors for the current state
func getSuccessors(state State) map[State]string {
	successors := make(map[State]string)
	region := erosion(state.x, state.y) % 3

	// Movement successors
	moves := [][3]interface{}{
		{state.x + 1, state.y, "Move R"},
		{state.x - 1, state.y, "Move L"},
		{state.x, state.y + 1, "Move D"},
		{state.x, state.y - 1, "Move U"},
	}

	for _, move := range moves {
		x, y := move[0].(int), move[1].(int)
		if x >= 0 && y >= 0 && x <= tx+49 && y <= ty+49 {
			if itemValidForRegion(state.item, erosion(x, y)%3) {
				successors[State{x, y, state.item}] = move[2].(string)
			}
		}
	}

	// Equipment changes
	switch region {
	case rock:
		if state.item == "Torch" {
			successors[State{state.x, state.y, "Climb"}] = "Torch -> Climb"
		} else if state.item == "Climb" {
			successors[State{state.x, state.y, "Torch"}] = "Climb -> Torch"
		}
	case water:
		if state.item == "Nothing" {
			successors[State{state.x, state.y, "Climb"}] = "Nothing -> Climb"
		} else if state.item == "Climb" {
			successors[State{state.x, state.y, "Nothing"}] = "Climb -> Nothing"
		}
	case narrow:
		if state.item == "Torch" {
			successors[State{state.x, state.y, "Nothing"}] = "Torch -> Nothing"
		} else if state.item == "Nothing" {
			successors[State{state.x, state.y, "Torch"}] = "Nothing -> Torch"
		}
	}

	return successors
}

// Calculate cost of an action
func actionCost(action string) int {
	if len(action) >= 4 && action[:4] == "Move" {
		return 1
	}
	return 7
}

// Get final state from path
func finalState(path Path) State {
	if len(path) == 0 {
		return State{}
	}
	return path[len(path)-1].state
}

// Add path to frontier with cost consideration
func addToFrontier(frontier *[]Path, newPath Path) {
	newState := finalState(newPath)
	newCost := pathCost(newPath)

	// Find if there's an existing path to the same state
	for i, path := range *frontier {
		if finalState(path) == newState {
			if pathCost(path) <= newCost {
				return // Existing path is better
			}
			// Remove more expensive path
			*frontier = append((*frontier)[:i], (*frontier)[i+1:]...)
			break
		}
	}

	// Add new path
	*frontier = append(*frontier, newPath)

	// Sort frontier by path cost
	sort.Slice(*frontier, func(i, j int) bool {
		return pathCost((*frontier)[i]) < pathCost((*frontier)[j])
	})
}

// Calculate total cost of a path
func pathCost(path Path) int {
	if len(path) == 0 {
		return 0
	}
	return path[len(path)-1].totalCost
}

func main() {
	riskSum := 0
	for y := 0; y <= ty; y++ {
		for x := 0; x <= tx; x++ {
			riskSum += erosion(x, y) % 3
		}
	}
	fmt.Println("Part 1: What is the total risk level for the smallest rectangle that includes 0,0 and the target's coordinates?")
	fmt.Println(riskSum)

	// Part 2
	p2path := Path{}
	start := State{0, 0, "Torch"}
	explored := make(map[State]bool)
	frontier := []Path{{PathNode{start, "", 0}}}

	for len(frontier) > 0 {
		path := frontier[0]
		frontier = frontier[1:]

		currentState := finalState(path)
		if currentState.x == tx && currentState.y == ty && currentState.item == "Torch" {
			p2path = path
			break
		}

		explored[currentState] = true
		currentCost := pathCost(path)

		for nextState, action := range getSuccessors(currentState) {
			if !explored[nextState] {
				totalCost := currentCost + actionCost(action)
				newPath := make(Path, len(path))
				copy(newPath, path)
				newPath = append(newPath, PathNode{nextState, action, totalCost})
				addToFrontier(&frontier, newPath)
			}
		}
	}
	fmt.Println()
	fmt.Println("Part 2: What is the fewest number of minutes you can take to reach the target?")
	fmt.Println(pathCost(p2path))
}
