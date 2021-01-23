package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read instructions
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	registers := map[string]int{"a": 0, "b": 0}
	pc := 0

	for pc < len(fileLines) {
		line := fileLines[pc]
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "hlf":
			registers[parts[1]] /= 2
		case "tpl":
			registers[parts[1]] *= 3
		case "inc":
			registers[parts[1]]++
		case "jmp":
			offset, _ := strconv.Atoi(parts[1])
			pc += offset - 1
		case "jie":
			register := strings.Trim(parts[1], ",")
			if registers[register]%2 == 0 {
				offset, _ := strconv.Atoi(parts[2])
				pc += offset - 1
			}
		case "jio":
			register := strings.Trim(parts[1], ",")
			if registers[register] == 1 {
				offset, _ := strconv.Atoi(parts[2])
				pc += offset - 1
			}
		}
		pc++
	}
	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 23, part 1:")
	fmt.Println("What is the value in register b when the program in your puzzle input is finished executing?")
	fmt.Println(registers["b"])

	registers = map[string]int{"a": 1, "b": 0}
	pc = 0

	for pc < len(fileLines) {
		line := fileLines[pc]
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "hlf":
			registers[parts[1]] /= 2
		case "tpl":
			registers[parts[1]] *= 3
		case "inc":
			registers[parts[1]]++
		case "jmp":
			offset, _ := strconv.Atoi(parts[1])
			pc += offset - 1
		case "jie":
			register := strings.Trim(parts[1], ",")
			if registers[register]%2 == 0 {
				offset, _ := strconv.Atoi(parts[2])
				pc += offset - 1
			}
		case "jio":
			register := strings.Trim(parts[1], ",")
			if registers[register] == 1 {
				offset, _ := strconv.Atoi(parts[2])
				pc += offset - 1
			}
		}
		pc++
	}

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2:")
	fmt.Println("what is the value in register b after the program is finished executing if register a starts as 1 instead?")
	fmt.Println(registers["b"])

}
