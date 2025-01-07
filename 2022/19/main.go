package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Recipe [4]uint16
type Blueprint struct {
	recipes [4]Recipe
}

func parseBlueprint(line string) Blueprint {
	words := strings.Fields(line)
	getNum := func(i int) uint16 {
		n, _ := strconv.ParseUint(words[i], 10, 16)
		return uint16(n)
	}

	return Blueprint{recipes: [4]Recipe{
		{getNum(6), 0, 0, 0},
		{getNum(12), 0, 0, 0},
		{getNum(18), getNum(21), 0, 0},
		{getNum(27), 0, getNum(30), 0},
	}}
}

type State struct {
	ores   [4]uint16
	robots [4]uint16
	time   uint16
}

func max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

func recurseSimulation(blueprint *Blueprint, state State, maxTime uint16, maxRobots *[4]uint16, maxGeodes *uint16) {
	hasRecursed := false
	for i := 0; i < 4; i++ {
		if state.robots[i] == maxRobots[i] {
			continue
		}
		recipe := &blueprint.recipes[i]

		var waitTime uint16
		for oreType := 0; oreType < 3; oreType++ {
			if recipe[oreType] == 0 {
				continue
			}
			if recipe[oreType] <= state.ores[oreType] {
				continue
			}
			if state.robots[oreType] == 0 {
				waitTime = maxTime + 1
				break
			}
			t := (recipe[oreType] - state.ores[oreType] + state.robots[oreType] - 1) / state.robots[oreType]
			if t > waitTime {
				waitTime = t
			}
		}

		timeFinished := state.time + waitTime + 1
		if timeFinished >= maxTime {
			continue
		}

		var newOres, newRobots [4]uint16
		for o := 0; o < 4; o++ {
			newOres[o] = state.ores[o] + state.robots[o]*(waitTime+1) - recipe[o]
			newRobots[o] = state.robots[o]
			if o == i {
				newRobots[o]++
			}
		}

		remainingTime := maxTime - timeFinished
		if ((remainingTime-1)*remainingTime)/2+newOres[3]+remainingTime*newRobots[3] < *maxGeodes {
			continue
		}

		hasRecursed = true
		recurseSimulation(blueprint, State{
			ores:   newOres,
			robots: newRobots,
			time:   timeFinished,
		}, maxTime, maxRobots, maxGeodes)
	}

	if !hasRecursed {
		*maxGeodes = max(*maxGeodes, state.ores[3]+state.robots[3]*(maxTime-state.time))
	}
}

func simulateBlueprint(blueprint *Blueprint, maxTime uint16) uint16 {
	var maxRobots [4]uint16
	for i := 0; i < 4; i++ {
		maxRobots[i] = ^uint16(0)
	}
	for i := 0; i < 3; i++ {
		var m uint16
		for _, r := range blueprint.recipes {
			if r[i] > m {
				m = r[i]
			}
		}
		maxRobots[i] = m
	}

	var maxGeodes uint16
	recurseSimulation(blueprint, State{
		robots: [4]uint16{1, 0, 0, 0},
	}, maxTime, &maxRobots, &maxGeodes)
	return maxGeodes
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var blueprints []Blueprint
	for scanner.Scan() {
		blueprints = append(blueprints, parseBlueprint(scanner.Text()))
	}

	sum := 0
	for i, b := range blueprints {
		sum += int(simulateBlueprint(&b, 24)) * (i + 1)
	}

	fmt.Println("Part 1: What do you get if you add up the quality level of all of the blueprints in your list?")
	fmt.Println(sum)

	product := 1
	for i := 0; i < 3 && i < len(blueprints); i++ {
		product *= int(simulateBlueprint(&blueprints[i], 32))
	}
	fmt.Println()
	fmt.Println("Part 2: What do you get if you multiply these numbers together?")
	fmt.Println(product)
}
