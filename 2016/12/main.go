package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 12: Leonardo's Monorail")
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	instructions := strings.Split(string(fileContent), "\n")
	registers := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	}

	fmt.Println("After executing the assembunny code in your puzzle input, what value is left in register a?")
	regA := runProgram(instructions, registers)
	fmt.Println(regA)

	registers = map[string]int{
		"a": 0,
		"b": 0,
		"c": 1,
		"d": 0,
	}
	fmt.Println()
	fmt.Println("Part 2/")
	fmt.Println("If you instead initialize register c to be 1, what value is now left in register a?")
	regA = runProgram(instructions, registers)
	fmt.Println(regA)
}

func runProgram(instructions []string, registers map[string]int) int {
	var i int
	for i < len(instructions) {
		instruction := instructions[i]
		params := strings.Fields(instruction)

		switch params[0] {
		case "cpy":
			v, err := strconv.Atoi(params[1])
			if err != nil {
				registers[params[2]] = registers[params[1]]
			} else {
				registers[params[2]] = v
			}
		case "inc":
			if _, ok := registers[params[1]]; !ok {
				registers[params[1]]++
			} else {
				registers[params[1]]++
			}
		case "dec":
			if _, ok := registers[params[1]]; !ok {
				registers[params[1]]--
			} else {
				registers[params[1]]--
			}
		case "jnz":
			v, _ := strconv.Atoi(params[1])
			if (unicode.IsLetter(rune(params[1][0])) && registers[params[1]] != 0) ||
				(unicode.IsNumber(rune(params[1][0])) && v != 0) {
				skip, _ := strconv.Atoi(params[2])
				i += skip
				continue
			}
		}
		i++
	}
	return registers["a"]
}
