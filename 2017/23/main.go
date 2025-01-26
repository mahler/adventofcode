package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(part int, instructions []string) int {
	registers := make(map[string]int)
	registers["a"] = part - 1

	interpret := func(val string) int {
		if strings.ContainsAny(val, "abcdefgh") {
			return registers[val]
		}
		num, _ := strconv.Atoi(val)
		return num
	}

	i := 0
	for i < 11 {
		parts := strings.Fields(instructions[i])
		op, reg, val := parts[0], parts[1], parts[2]

		switch op {
		case "set":
			registers[reg] = interpret(val)
		case "sub":
			registers[reg] -= interpret(val)
		case "mul":
			registers[reg] *= interpret(val)
		case "jnz":
			if interpret(reg) != 0 {
				i += interpret(val)
				continue
			}
		}
		i++
	}

	if part == 1 {
		return (registers["b"] - registers["e"]) * (registers["b"] - registers["d"])
	}

	nonprimes := 0
	for b := registers["b"]; b <= registers["c"]; b += 17 {
		if hasCompositeFactors(b) {
			nonprimes++
		}
	}
	return nonprimes
}

func hasCompositeFactors(b int) bool {
	for d := 2; d <= int(math.Sqrt(float64(b))); d++ {
		if b%d == 0 {
			return true
		}
	}
	return false
}

func main() {
	data, _ := os.ReadFile("input.txt")
	instructions := strings.Split(string(data), "\n")
	instructions = instructions[:len(instructions)-1] // remove last empty line

	fmt.Println(solve(1, instructions))
	fmt.Println(solve(2, instructions))
}
