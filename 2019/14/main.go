package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Reaction struct {
	outputAmount int
	inputs       map[string]int
}

func parseInput(filename string) map[string]Reaction {
	file, _ := os.Open(filename)
	defer file.Close()

	reactions := make(map[string]Reaction)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " => ")
		inputs := strings.Split(parts[0], ", ")
		outputParts := strings.Split(parts[1], " ")
		outputAmount, _ := strconv.Atoi(outputParts[0])
		outputChemical := outputParts[1]

		inputMap := make(map[string]int)
		for _, input := range inputs {
			inputParts := strings.Split(input, " ")
			amount, _ := strconv.Atoi(inputParts[0])
			chemical := inputParts[1]
			inputMap[chemical] = amount
		}

		reactions[outputChemical] = Reaction{outputAmount, inputMap}
	}
	return reactions
}

func calculateOre(fuelAmount int, reactions map[string]Reaction) int {
	requirements := make(map[string]int)
	requirements["FUEL"] = fuelAmount
	surplus := make(map[string]int)
	oreTotal := 0

	for len(requirements) > 0 {
		for chemical, amountNeeded := range requirements {
			if chemical == "ORE" {
				oreTotal += amountNeeded
				delete(requirements, chemical)
				continue
			}

			reaction := reactions[chemical]
			if surplus[chemical] > 0 {
				if surplus[chemical] >= amountNeeded {
					surplus[chemical] -= amountNeeded
					delete(requirements, chemical)
					continue
				} else {
					amountNeeded -= surplus[chemical]
					surplus[chemical] = 0
				}
			}

			batches := int(math.Ceil(float64(amountNeeded) / float64(reaction.outputAmount)))
			surplus[chemical] += batches*reaction.outputAmount - amountNeeded
			delete(requirements, chemical)

			for inputChemical, inputAmount := range reaction.inputs {
				requirements[inputChemical] += batches * inputAmount
			}
		}
	}

	return oreTotal
}

func maxFuel(oreLimit int, reactions map[string]Reaction) int {
	low, high := 1, oreLimit
	for low < high {
		mid := (low + high + 1) / 2
		if calculateOre(mid, reactions) > oreLimit {
			high = mid - 1
		} else {
			low = mid
		}
	}
	return low
}

func main() {
	reactions := parseInput("input.txt")
	oneFuel := calculateOre(1, reactions)
	fmt.Println("Part 1: What is the minimum amount of ORE required to produce exactly 1 FUEL?")
	fmt.Println(oneFuel)

	maxFuelProduced := maxFuel(1000000000000, reactions)
	fmt.Println()
	fmt.Println("Part 2: Given 1 trillion ORE, what is the maximum amount of FUEL you can produce?")
	fmt.Println(maxFuelProduced)
}
