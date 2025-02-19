package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Program represents an Intcode computer instance
type Program struct {
	P       map[int]int
	ip      int
	pid     int
	relBase int
	halted  bool
	input   func() int
}

// NewProgram creates a new Intcode computer with the given program file and input function
func NewProgram(pid int, programFile string, input func() int) (*Program, error) {
	p := &Program{
		P:       make(map[int]int),
		pid:     pid,
		input:   input,
		halted:  false,
		ip:      0,
		relBase: 0,
	}

	data, err := os.ReadFile(programFile)
	if err != nil {
		return nil, err
	}

	for i, x := range strings.Split(strings.TrimSpace(string(data)), ",") {
		val, err := strconv.Atoi(x)
		if err != nil {
			return nil, err
		}
		p.P[i] = val
	}

	return p, nil
}

// idx returns the memory address based on the parameter mode
func (p *Program) idx(i int, modes []int) int {
	val := p.P[p.ip+1+i]
	if i >= len(modes) || modes[i] == 0 {
		// Position mode
		return val
	} else if modes[i] == 2 {
		// Relative mode
		return val + p.relBase
	}
	panic(fmt.Sprintf("Unknown mode: %d", modes[i]))
}

// val returns the parameter value based on the parameter mode
func (p *Program) val(i int, modes []int) int {
	val := p.P[p.ip+1+i]
	if i >= len(modes) || modes[i] == 0 {
		// Position mode
		return p.P[val]
	} else if modes[i] == 1 {
		// Immediate mode
		return val
	} else if modes[i] == 2 {
		// Relative mode
		return p.P[val+p.relBase]
	}
	panic(fmt.Sprintf("Unknown mode: %d", modes[i]))
}

// RunAll executes the program until it halts and returns all outputs
func (p *Program) RunAll() []int {
	var ans []int
	for {
		val := p.Run(0)
		if val == nil {
			return ans
		}
		ans = append(ans, *val)
	}
}

// Run executes the program until it produces an output or halts
// If timeout > 0, it will run at most timeout instructions
// Returns nil if the program halts or reaches the timeout
func (p *Program) Run(timeout int) *int {
	t := 0
	for (timeout == 0) || (t < timeout) {
		cmd := strconv.Itoa(p.P[p.ip])
		var opcode int
		if len(cmd) == 1 {
			opcode = p.P[p.ip]
		} else {
			opcode, _ = strconv.Atoi(cmd[len(cmd)-2:])
		}

		// Parse parameter modes
		var modes []int
		for i := len(cmd) - 3; i >= 0; i-- {
			mode, _ := strconv.Atoi(string(cmd[i]))
			modes = append(modes, mode)
		}

		switch opcode {
		case 1: // Add
			p.P[p.idx(2, modes)] = p.val(0, modes) + p.val(1, modes)
			p.ip += 4
		case 2: // Multiply
			p.P[p.idx(2, modes)] = p.val(0, modes) * p.val(1, modes)
			p.ip += 4
		case 3: // Input
			inp := p.input()
			p.P[p.idx(0, modes)] = inp
			p.ip += 2
		case 4: // Output
			ans := p.val(0, modes)
			p.ip += 2
			return &ans
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
				p.P[p.idx(2, modes)] = 1
			} else {
				p.P[p.idx(2, modes)] = 0
			}
			p.ip += 4
		case 8: // Equals
			if p.val(0, modes) == p.val(1, modes) {
				p.P[p.idx(2, modes)] = 1
			} else {
				p.P[p.idx(2, modes)] = 0
			}
			p.ip += 4
		case 9: // Adjust relative base
			p.relBase += p.val(0, modes)
			p.ip += 2
		case 99: // Halt
			p.halted = true
			return nil
		default:
			panic(fmt.Sprintf("Unknown opcode: %d", opcode))
		}
		t++
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <program_file>")
		os.Exit(1)
	}

	const N = 50
	queues := make([][]int, N)

	getInput := func(i int) func() int {
		return func() int {
			if len(queues[i]) > 0 {
				val := queues[i][0]
				queues[i] = queues[i][1:]
				return val
			}
			return -1
		}
	}

	programs := make([]*Program, N)
	for i := 0; i < N; i++ {
		queues[i] = []int{i} // Initialize with network address
		var err error
		programs[i], err = NewProgram(i, os.Args[1], getInput(i))
		if err != nil {
			fmt.Printf("Error loading program: %v\n", err)
			os.Exit(1)
		}
	}

	var nat [2]int
	var lastNat *int
	change := 0
	t := 0

	duplicateY := 0
	firstYtoAddr := true
	firstYvalue := 0

	for {
		if t > change+500 {
			if nat[0] != 0 || nat[1] != 0 { // NAT initialized
				queues[0] = append(queues[0], nat[0], nat[1])
				//fmt.Println("NAT sends:", nat)
				if firstYtoAddr {
					firstYtoAddr = false
					firstYvalue = nat[1]
				}

				if lastNat != nil && nat[1] == *lastNat {
					//fmt.Println("Duplicate Y value delivered by NAT:", *lastNat)
					duplicateY = *lastNat
					break
					//					os.Exit(0)
				}
				lastNat = new(int)
				*lastNat = nat[1]
				change = t
			}
		}

		for i := 0; i < N; i++ {
			addrPtr := programs[i].Run(10)
			if addrPtr == nil {
				continue
			}
			change = t

			addr := *addrPtr
			xPtr := programs[i].Run(0)
			yPtr := programs[i].Run(0)

			if xPtr == nil || yPtr == nil {
				continue // This shouldn't happen, but just in case
			}

			x, y := *xPtr, *yPtr

			if addr == 255 {
				//	fmt.Printf("To NAT: %d %d %d\n", addr, x, y)
				nat[0], nat[1] = x, y
			} else if addr < N {
				queues[addr] = append(queues[addr], x, y)
			} else {
				fmt.Printf("Invalid address: %d\n", addr)
			}
		}
		t++
	}
	fmt.Println("Part 1:  What is the Y value of the first packet sent to address 255?")
	fmt.Println(firstYvalue)

	fmt.Println()
	fmt.Println("Part 2: What is the first Y value delivered by the NAT to the computer at address 0 twice in a row?")
	fmt.Println(duplicateY)
}
