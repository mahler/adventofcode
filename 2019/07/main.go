package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run(program []int, inputs []int, pc int) (int, int) {
	numOfOperands := []int{0, 3, 3, 1, 1, 2, 2, 3, 3}
	i := pc
	inputIdx := 0
	for program[i] != 99 {
		instrStr := fmt.Sprintf("%05d", program[i])
		modes := []int{
			int(instrStr[2] - '0'),
			int(instrStr[1] - '0'),
			int(instrStr[0] - '0'),
		}
		instruction, _ := strconv.Atoi(instrStr[3:])
		opCount := numOfOperands[instruction]
		operands := make([]int, opCount)

		for x := 0; x < opCount; x++ {
			if modes[x] == 1 {
				operands[x] = program[i+x+1]
			} else {
				operands[x] = program[program[i+x+1]]
			}
		}

		switch instruction {
		case 1:
			program[program[i+3]] = operands[0] + operands[1]
		case 2:
			program[program[i+3]] = operands[0] * operands[1]
		case 3:
			program[program[i+1]] = inputs[inputIdx]
			inputIdx++
		case 4:
			return operands[0], i + 2
		case 5:
			if operands[0] != 0 {
				i = operands[1] - 3
			}
		case 6:
			if operands[0] == 0 {
				i = operands[1] - 3
			}
		case 7:
			if operands[0] < operands[1] {
				program[program[i+3]] = 1
			} else {
				program[program[i+3]] = 0
			}
		case 8:
			if operands[0] == operands[1] {
				program[program[i+3]] = 1
			} else {
				program[program[i+3]] = 0
			}
		}
		i += opCount + 1
	}
	return -1, -1
}

func setSequence(line []int, settings []int, feedback bool) int {
	prev := 0
	amplifiers := make([][]int, 5)
	pcs := make([]int, 5)

	for i := 0; i < 5; i++ {
		amplifiers[i] = append([]int(nil), line...)
		output, pc := run(amplifiers[i], []int{settings[i], prev}, 0)
		prev = output
		pcs[i] = pc
	}

	if !feedback {
		return prev
	}
	i := 0
	for {
		output, pc := run(amplifiers[i], []int{prev}, pcs[i])
		if output == -1 {
			break
		}
		prev = output
		pcs[i] = pc
		i = (i + 1) % 5
	}
	return prev
}

func permute(arr []int, index int, res *[][]int) {
	if index >= len(arr) {
		copyArr := append([]int(nil), arr...)
		*res = append(*res, copyArr)
		return
	}
	for i := index; i < len(arr); i++ {
		arr[index], arr[i] = arr[i], arr[index]
		permute(arr, index+1, res)
		arr[index], arr[i] = arr[i], arr[index]
	}
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lineStr := strings.Split(strings.TrimSpace(string(data)), ",")
	line := make([]int, len(lineStr))
	for i, v := range lineStr {
		line[i], _ = strconv.Atoi(v)
	}

	var res1, res2 [][]int
	permute([]int{0, 1, 2, 3, 4}, 0, &res1)
	permute([]int{5, 6, 7, 8, 9}, 0, &res2)

	max1, max2 := 0, 0
	for _, per := range res1 {
		if val := setSequence(line, per, false); val > max1 {
			max1 = val
		}
	}

	fmt.Println("Part 1: What is the highest signal that can be sent to the thrusters?")
	fmt.Println(max1)

	for _, per := range res2 {
		if val := setSequence(line, per, true); val > max2 {
			max2 = val
		}
	}

	fmt.Println()
	fmt.Println("Part 2: What is the highest signal that can be sent to the thrusters?")
	fmt.Println(max2)
}
