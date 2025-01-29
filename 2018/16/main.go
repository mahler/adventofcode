package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type OpFunc func(regs []int, a, b int) int

var ops = map[string]OpFunc{
	"addr": func(regs []int, a, b int) int { return regs[a] + regs[b] },
	"addi": func(regs []int, a, b int) int { return regs[a] + b },
	"mulr": func(regs []int, a, b int) int { return regs[a] * regs[b] },
	"muli": func(regs []int, a, b int) int { return regs[a] * b },
	"banr": func(regs []int, a, b int) int { return regs[a] & regs[b] },
	"bani": func(regs []int, a, b int) int { return regs[a] & b },
	"borr": func(regs []int, a, b int) int { return regs[a] | regs[b] },
	"bori": func(regs []int, a, b int) int { return regs[a] | b },
	"setr": func(regs []int, a, b int) int { return regs[a] },
	"seti": func(regs []int, a, b int) int { return a },
	"gtir": func(regs []int, a, b int) int {
		if a > regs[b] {
			return 1
		} else {
			return 0
		}
	},
	"gtri": func(regs []int, a, b int) int {
		if regs[a] > b {
			return 1
		} else {
			return 0
		}
	},
	"gtrr": func(regs []int, a, b int) int {
		if regs[a] > regs[b] {
			return 1
		} else {
			return 0
		}
	},
	"eqir": func(regs []int, a, b int) int {
		if a == regs[b] {
			return 1
		} else {
			return 0
		}
	},
	"eqri": func(regs []int, a, b int) int {
		if regs[a] == b {
			return 1
		} else {
			return 0
		}
	},
	"eqrr": func(regs []int, a, b int) int {
		if regs[a] == regs[b] {
			return 1
		} else {
			return 0
		}
	},
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(input, "\n\n")
	samples := parts[:len(parts)-2]
	program := parts[len(parts)-1]

	indeterminate := 0
	possible := make(map[int]map[string]bool)

	// Initialize possible operations map
	for i := 0; i < 16; i++ {
		possible[i] = make(map[string]bool)
		for op := range ops {
			possible[i][op] = true
		}
	}

	r := regexp.MustCompile(`-?\d+`)
	for _, sample := range samples {
		lines := strings.Split(sample, "\n")
		before := extractNumbers(r.FindAllString(lines[0], -1))
		op := extractNumbers(r.FindAllString(lines[1], -1))
		after := extractNumbers(r.FindAllString(lines[2], -1))

		count := 0
		for opcode := range ops {
			result := ops[opcode](before, op[1], op[2])
			resultRegs := make([]int, len(before))
			copy(resultRegs, before)
			resultRegs[op[3]] = result

			if sliceEqual(resultRegs, after) {
				count++
			} else {
				if possible[op[0]][opcode] {
					delete(possible[op[0]], opcode)
				}
			}
		}

		if count >= 3 {
			indeterminate++
		}
	}

	fmt.Println("Part 1: Ignoring the opcode numbers, how many samples in your puzzle input behave like three or more opcodes?")
	fmt.Println(indeterminate)

	// Determine final mapping
	mapping := make(map[int]string)
	for len(mapping) < len(ops) {
		for number, opcodes := range possible {
			if len(opcodes) == 1 {
				var op string
				for k := range opcodes {
					op = k
					break
				}
				mapping[number] = op
				// Remove this opcode from all other possibilities
				for _, remaining := range possible {
					delete(remaining, op)
				}
			}
		}
	}

	// Process program
	registers := []int{0, 0, 0, 0}
	for _, line := range strings.Split(program, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 4 {
			continue
		}

		opcode, _ := strconv.Atoi(parts[0])
		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		c, _ := strconv.Atoi(parts[3])

		registers[c] = ops[mapping[opcode]](registers, a, b)
	}

	fmt.Println()
	fmt.Println("Part 2: What value is contained in register 0 after executing the test program?")
	fmt.Println(registers[0])
}

func readInput(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func extractNumbers(numbers []string) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i], _ = strconv.Atoi(n)
	}
	return result
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
