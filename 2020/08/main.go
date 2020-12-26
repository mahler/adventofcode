package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("game.console")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	records := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("DAY08, Part 1: Handheld Halting")

	mapInstructions := make(map[int]string)
	mapOperations := make(map[int]int)

	// Parse file database into two maps.
	for i, record := range records {
		if len(record) == 0 {
			break
		}
		strs := strings.Split(record, " ")
		instruction := strs[0]
		operation, _ := strconv.Atoi(strs[1])

		mapInstructions[i] = instruction
		mapOperations[i] = operation
	}

	// initialize accumulator.
	accumulator := 0
	// initialize stackPosition.
	stackPos := 0

	mapBreaker := make(map[int]int)

	loops := 0
	for {
		//fmt.Printf("%v: %v (%v) Acc: %v - pos: %v\n", loops, mapInstructions[stackPos], mapOperations[stackPos], accumulator, stackPos)

		if _, ok := mapBreaker[stackPos]; !ok {
			mapBreaker[stackPos] = 1
			switch mapInstructions[stackPos] {
			case "nop":
				stackPos++
			case "acc":
				accumulator += mapOperations[stackPos]
				stackPos++
			case "jmp":
				stackPos += mapOperations[stackPos]

			default:
				fmt.Println("PANIC! - Unknown instruction")
			}
		} else {
			break
		}

		// Loop circuit breaker
		loops++
		if loops > 2000 {
			break
		}

	}
	fmt.Println("Accumulator:", accumulator)

	fmt.Println()
	fmt.Println("DAY08, Part 2: Debugging")

	// Potentiallyy as many rounds as rows in dataset
	foundResult := 0
	for round := 0; round < len(mapInstructions); round++ {
		copyInstructions := make(map[int]string)

		for key, value := range mapInstructions {
			copyInstructions[key] = value
		}

		// Try to switch up instruction
		switch copyInstructions[round] {
		case "nop":
			copyInstructions[round] = "jmp"
		case "jmp":
			copyInstructions[round] = "nop"
		default:
			continue
		}

		// Initialize test run
		loops := 0
		accumulator := 0
		stackPos := 0
		mapBreaker := make(map[int]int)

		for {
			//	fmt.Printf("%v: %v (%v) Acc: %v - pos: %v\n", loops, copyInstructions[stackPos], mapOperations[stackPos], accumulator, stackPos)
			if _, ok := mapBreaker[stackPos]; !ok {
				mapBreaker[stackPos] = 1
				switch copyInstructions[stackPos] {
				case "nop":
					stackPos++
				case "acc":
					accumulator += mapOperations[stackPos]
					stackPos++
				case "jmp":
					stackPos += mapOperations[stackPos]

				default:
					fmt.Println("PANIC! - Unknown instruction")
				}
			} else {
				break
			}

			if stackPos == len(copyInstructions)-1 {
				foundResult = accumulator
				break
			}
			// Loop circuit breaker
			loops++
			if loops > 2000 {
				break
			}
		}
		if foundResult > 0 {
			break
		}
	}
	fmt.Println("Accumulator value:", foundResult)

}
