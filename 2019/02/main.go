package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}

	var intOpCode []int
	operations := strings.Split(string(data), ",")
	for _, value := range operations {
		opVal, _ := strconv.Atoi(value)
		intOpCode = append(intOpCode, opVal)
	}

	//  before running the program, replace position 1 with the value 12 and replace position 2 with the value 2.
	intOpCode[0] = 1
	p1opCode := make([]int, len(intOpCode))
	copy(p1opCode, intOpCode)

	p1opCode[1] = 12
	p1opCode[2] = 2

	fmt.Println()
	fmt.Println("Day 02: 1202 Program Alarm")

	result, _ := runProgram(p1opCode)
	fmt.Println("result:", result)

	fmt.Println()
	fmt.Println("Part 2 - looking for 19690720")

	foundMatch := false
	p2result := 0
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			// make copy of program to test on
			tmpCode := make([]int, len(intOpCode))
			copy(tmpCode, intOpCode)
			// setup test by modifying program data
			tmpCode[1] = noun
			tmpCode[2] = verb
			// Run test
			result, _ := runProgram(tmpCode)
			if result == 19690720 {
				p2result = noun*100 + verb
				foundMatch = true
				break
			}
		}
		if foundMatch {
			break
		}
	}
	fmt.Println("Result for Part2:", p2result)
}

func runProgram(program []int) (result int, fault bool) {
	ip := 0
	for {
		switch program[ip] {
		case 1:
			program[program[ip+3]] = program[program[ip+1]] + program[program[ip+2]]
			ip += 4
		case 2:
			program[program[ip+3]] = program[program[ip+1]] * program[program[ip+2]]
			ip += 4
		case 99:
			return program[0], false
		default:
			return 0, true
		}
	}
}
