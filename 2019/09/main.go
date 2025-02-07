package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntcodeComputer struct {
	memory map[int]int
	ip     int
	rb     int
}

func NewIntcodeComputer(program string) *IntcodeComputer {
	mem := make(map[int]int)
	for i, s := range strings.Split(program, ",") {
		num, _ := strconv.Atoi(s)
		mem[i] = num
	}
	return &IntcodeComputer{memory: mem, ip: 0, rb: 0}
}

func (c *IntcodeComputer) run(input []int) []int {
	var outputs []int
	inputQueue := make([]int, len(input))
	copy(inputQueue, input)

	for {
		op := c.memory[c.ip] % 100
		if op == 99 {
			break
		}

		size := []int{0, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
		args := make([]int, size-1)
		for i := 1; i < size; i++ {
			args[i-1] = c.memory[c.ip+i]
		}

		modes := make([]int, 3)
		for i := 0; i < 3; i++ {
			modes[i] = (c.memory[c.ip] / powInt(10, i+2)) % 10
		}

		reads := make([]int, len(args))
		writes := make([]int, len(args))
		for i, arg := range args {
			switch modes[i] {
			case 0: // position mode
				reads[i] = c.memory[arg]
				writes[i] = arg
			case 1: // immediate mode
				reads[i] = arg
				writes[i] = arg
			case 2: // relative mode
				reads[i] = c.memory[arg+c.rb]
				writes[i] = arg + c.rb
			}
		}

		c.ip += size

		switch op {
		case 1: // add
			c.memory[writes[2]] = reads[0] + reads[1]
		case 2: // multiply
			c.memory[writes[2]] = reads[0] * reads[1]
		case 3: // input
			c.memory[writes[0]] = inputQueue[0]
			inputQueue = inputQueue[1:]
		case 4: // output
			outputs = append(outputs, reads[0])
		case 5: // jump-if-true
			if reads[0] != 0 {
				c.ip = reads[1]
			}
		case 6: // jump-if-false
			if reads[0] == 0 {
				c.ip = reads[1]
			}
		case 7: // less than
			if reads[0] < reads[1] {
				c.memory[writes[2]] = 1
			} else {
				c.memory[writes[2]] = 0
			}
		case 8: // equals
			if reads[0] == reads[1] {
				c.memory[writes[2]] = 1
			} else {
				c.memory[writes[2]] = 0
			}
		case 9: // adjust relative base
			c.rb += reads[0]
		}
	}

	return outputs
}

func powInt(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	program := scanner.Text()

	computer1 := NewIntcodeComputer(program)
	fmt.Println("Part 1: What BOOST keycode does it produce?")
	fmt.Println(computer1.run([]int{1}))

	fmt.Println()
	fmt.Println("Part 2: What are the coordinates of the distress signal?")
	computer2 := NewIntcodeComputer(program)
	fmt.Println(computer2.run([]int{2}))
}
