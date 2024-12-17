package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ClawMachine struct {
	ButtonA [2]int
	ButtonB [2]int
	Prize   [2]int
}

func parseInput(scanner *bufio.Scanner) []ClawMachine {
	var machines []ClawMachine
	var buttonA, buttonB [2]int
	numberRegex := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		line := scanner.Text()
		numberStrs := numberRegex.FindAllString(line, -1)

		if len(numberStrs) > 0 {
			numbers := make([]int, len(numberStrs))
			for i, numStr := range numberStrs {
				numbers[i], _ = strconv.Atoi(numStr)
			}

			switch {
			case len(numbers) == 2 && (line[0] == 'B' && line[7] == 'A'):
				buttonA = [2]int{numbers[0], numbers[1]}
			case len(numbers) == 2 && (line[0] == 'B' && line[7] == 'B'):
				buttonB = [2]int{numbers[0], numbers[1]}
			case len(numbers) == 2 && line[0] == 'P':
				machines = append(machines, ClawMachine{
					ButtonA: buttonA,
					ButtonB: buttonB,
					Prize:   [2]int{numbers[0], numbers[1]},
				})
			}
		}
	}

	return machines
}

func solveMachine(machine ClawMachine) *int {
	aX, aY := machine.ButtonA[0], machine.ButtonA[1]
	bX, bY := machine.ButtonB[0], machine.ButtonB[1]
	prizeX, prizeY := machine.Prize[0], machine.Prize[1]

	aMultiplier := aX*bY - aY*bX
	aRhs := prizeX*bY - prizeY*bX

	if aMultiplier != 0 {
		if aRhs%aMultiplier == 0 && aRhs/aMultiplier >= 0 {
			a := aRhs / aMultiplier
			bRhs := prizeX - a*aX
			if bRhs%bX == 0 && bRhs/bX >= 0 {
				b := bRhs / bX
				result := 3*a + b
				return &result
			}
		}
	} else if aRhs == 0 {
		// Here be dragons!
		panic("Unhandled edge case")
	}

	return nil
}

func updateMachine(machine ClawMachine) ClawMachine {
	return ClawMachine{
		ButtonA: machine.ButtonA,
		ButtonB: machine.ButtonB,
		Prize:   [2]int{machine.Prize[0] + 10000000000000, machine.Prize[1] + 10000000000000},
	}
}

func main() {
	data, _ := os.Open("input.txt")
	defer data.Close()
	scanner := bufio.NewScanner(data)
	clawMachines := parseInput(scanner)

	// Part 1
	var solutions1 []int
	for _, machine := range clawMachines {
		if solution := solveMachine(machine); solution != nil {
			solutions1 = append(solutions1, *solution)
		}
	}

	part1Sum := 0
	for _, solution := range solutions1 {
		part1Sum += solution
	}
	fmt.Println("What is the fewest tokens you would have to spend to win all possible prizes?")
	fmt.Println(part1Sum)

	// Part 2
	updatedMachines := make([]ClawMachine, len(clawMachines))
	for i, machine := range clawMachines {
		updatedMachines[i] = updateMachine(machine)
	}

	var solutions2 []int
	for _, machine := range updatedMachines {
		if solution := solveMachine(machine); solution != nil {
			solutions2 = append(solutions2, *solution)
		}
	}

	part2Sum := 0
	for _, solution := range solutions2 {
		part2Sum += solution
	}
	fmt.Println()
	fmt.Println("What is the fewest tokens you would have to spend to win all possible prizes?")
	fmt.Println(part2Sum)
}
