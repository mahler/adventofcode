package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	memory    map[int]int
	ip        int
	pid       string
	relBase   int
	halted    bool
	inputFunc func() int
}

func NewProgram(pid string, programFile string, inputFunc func() int) (*Program, error) {
	p := &Program{
		memory:    make(map[int]int),
		pid:       pid,
		inputFunc: inputFunc,
	}

	file, err := os.Open(programFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := bufio.NewReader(file).ReadString('\n')
	if err != nil {
		return nil, err
	}

	for i, numStr := range strings.Split(strings.TrimSpace(content), ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		p.memory[i] = num
	}

	return p, nil
}

func (p *Program) idx(i int, modes []int) int {
	mode := 0
	if i < len(modes) {
		mode = modes[i]
	}
	val := p.memory[p.ip+1+i]

	switch mode {
	case 0:
		return val
	case 2:
		return val + p.relBase
	default:
		panic(fmt.Sprintf("invalid mode: %d", mode))
	}
}

func (p *Program) val(i int, modes []int) int {
	mode := 0
	if i < len(modes) {
		mode = modes[i]
	}
	val := p.memory[p.ip+1+i]

	switch mode {
	case 0:
		return p.memory[val]
	case 2:
		return p.memory[val+p.relBase]
	case 1:
		return val
	default:
		panic(fmt.Sprintf("invalid mode: %d", mode))
	}
}

func (p *Program) RunAll() []int {
	var result []int
	for {
		val := p.Run()
		if val == nil {
			return result
		}
		result = append(result, *val)
	}
}

func (p *Program) Run() *int {
	for {
		cmd := strconv.Itoa(p.memory[p.ip])
		var opcode int
		if len(cmd) >= 2 {
			opcode, _ = strconv.Atoi(cmd[len(cmd)-2:])
		} else {
			opcode, _ = strconv.Atoi(cmd)
		}

		// Parse modes
		modes := make([]int, 0)
		for i := len(cmd) - 3; i >= 0; i-- {
			mode, _ := strconv.Atoi(string(cmd[i]))
			modes = append(modes, mode)
		}

		switch opcode {
		case 1: // Add
			p.memory[p.idx(2, modes)] = p.val(0, modes) + p.val(1, modes)
			p.ip += 4
		case 2: // Multiply
			p.memory[p.idx(2, modes)] = p.val(0, modes) * p.val(1, modes)
			p.ip += 4
		case 3: // Input
			p.memory[p.idx(0, modes)] = p.inputFunc()
			p.ip += 2
		case 4: // Output
			val := p.val(0, modes)
			p.ip += 2
			return &val
		case 5: // Jump if true
			if p.val(0, modes) != 0 {
				p.ip = p.val(1, modes)
			} else {
				p.ip += 3
			}
		case 6: // Jump if false
			if p.val(0, modes) == 0 {
				p.ip = p.val(1, modes)
			} else {
				p.ip += 3
			}
		case 7: // Less than
			if p.val(0, modes) < p.val(1, modes) {
				p.memory[p.idx(2, modes)] = 1
			} else {
				p.memory[p.idx(2, modes)] = 0
			}
			p.ip += 4
		case 8: // Equals
			if p.val(0, modes) == p.val(1, modes) {
				p.memory[p.idx(2, modes)] = 1
			} else {
				p.memory[p.idx(2, modes)] = 0
			}
			p.ip += 4
		case 9: // Adjust relative base
			p.relBase += p.val(0, modes)
			p.ip += 2
		case 99: // Halt
			p.halted = true
			return nil
		default:
			panic(fmt.Sprintf("invalid opcode: %d", opcode))
		}
	}
}

// Queue implementation
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	if len(*q) == 0 {
		panic("empty queue")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

var inputQueue Queue

func getInput() int {
	return inputQueue.Pop()
}

func query(r, c int) int {
	p, err := NewProgram("0", "input.txt", getInput)
	if err != nil {
		panic(err)
	}
	inputQueue.Push(c)
	inputQueue.Push(r)
	val := p.Run()
	if val == nil {
		panic("unexpected nil output")
	}
	return *val
}

func main() {
	// Part 1
	part1 := 0
	for r := 0; r < 50; r++ {
		for c := 0; c < 50; c++ {
			if query(r, c) == 1 {
				part1++
			}
		}
	}
	fmt.Println("Part 1: How many points are affected by the tractor beam in the 50x50 area closest to the emitter?")
	fmt.Println(part1)

	// Part 2
	part2 := 0
	N := 10000
	lengths := make(map[int]struct{ start, end, length int })
	for r := 5; r < N; r++ {
		cstart := int(float64(r) * 1.01)
		cend := int(float64(r) * 1.26)

		// Find start of beam
		for query(r, cstart) == 0 {
			cstart++
		}

		// Find end of beam
		for query(r, cend) == 1 {
			cend++
		}
		cend--

		lengths[r] = struct{ start, end, length int }{cstart, cend, cend - cstart + 1}

		if prev, ok := lengths[r-99]; ok {
			if prev.end >= cstart+99 {
				sr, sc := r-99, cstart
				// Verify 100x100 square fits
				allOnes := true
				for dr := 0; dr < 100 && allOnes; dr++ {
					for dc := 0; dc < 100 && allOnes; dc++ {
						if query(sr+dr, sc+dc) != 1 {
							allOnes = false
						}
					}
				}
				if allOnes {
					part2 = (sc*10000 + sr)
					break
				}
			}
		}
	}
	fmt.Println()
	fmt.Println("Part 2: What value do you get if you take that point's X coordinate,")
	fmt.Println("multiply it by 10000, then add the point's Y coordinate?")
	fmt.Println(part2)
}
