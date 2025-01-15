package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	registers := make(map[string]int)
	maxThroughout := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		register := parts[0]
		instr := parts[1]
		amount, _ := strconv.Atoi(parts[2])
		ifRegister := parts[4]
		cond := parts[5]
		ifAmount, _ := strconv.Atoi(parts[6])

		ifVal := registers[ifRegister]
		if fulfillsCondition(ifVal, cond, ifAmount) {
			newVal := performOp(registers, register, instr, amount)
			if newVal > maxThroughout {
				maxThroughout = newVal
			}
		}
	}

	currentMax := 0
	for _, val := range registers {
		if val > currentMax {
			currentMax = val
		}
	}

	fmt.Println("Part 1: What is the largest value in any register after completing the instructions in your puzzle input?")
	fmt.Println(currentMax)
	fmt.Println()
	fmt.Println("Part 2: The highest value held in any register during this process?")
	fmt.Println(maxThroughout)
}

func fulfillsCondition(ifVal int, cond string, ifAmount int) bool {
	switch cond {
	case "<":
		return ifVal < ifAmount
	case ">":
		return ifVal > ifAmount
	case ">=":
		return ifVal >= ifAmount
	case "<=":
		return ifVal <= ifAmount
	case "==":
		return ifVal == ifAmount
	case "!=":
		return ifVal != ifAmount
	default:
		panic("unreachable")
	}
}

func performOp(registers map[string]int, reg string, instr string, amount int) int {
	if instr == "dec" {
		amount = -amount
	}
	registers[reg] += amount
	return registers[reg]
}
