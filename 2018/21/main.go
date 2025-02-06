package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	register     []int
	ip           int
	ipr          int
	program      []string
	counter      int
	instructions map[string]func(int, int, int)
}

func NewCPU(program []string) *CPU {
	cpu := &CPU{
		register: make([]int, 6),
		ip:       0,
		ipr:      0,
		program:  program,
		counter:  0,
	}

	cpu.instructions = map[string]func(int, int, int){
		"addr": cpu.addr,
		"addi": cpu.addi,
		"mulr": cpu.mulr,
		"muli": cpu.muli,
		"banr": cpu.banr,
		"bani": cpu.bani,
		"borr": cpu.borr,
		"bori": cpu.bori,
		"setr": cpu.setr,
		"seti": cpu.seti,
		"gtir": cpu.gtir,
		"gtri": cpu.gtri,
		"gtrr": cpu.gtrr,
		"eqir": cpu.eqir,
		"eqri": cpu.eqri,
		"eqrr": cpu.eqrr,
	}

	return cpu
}

func (c *CPU) bindIP(r int) {
	c.ipr = r
}

func (c *CPU) run() {
	c.register[c.ipr] = c.ip

	parts := strings.Fields(c.program[c.ip])
	opcode := parts[0]
	a, _ := strconv.Atoi(parts[1])
	b, _ := strconv.Atoi(parts[2])
	d, _ := strconv.Atoi(parts[3])

	c.instructions[opcode](a, b, d)

	c.ip = c.register[c.ipr]
	c.ip++
	c.counter++
}

func (c *CPU) addr(a, b, d int) { c.register[d] = c.register[a] + c.register[b] }
func (c *CPU) addi(a, b, d int) { c.register[d] = c.register[a] + b }
func (c *CPU) mulr(a, b, d int) { c.register[d] = c.register[a] * c.register[b] }
func (c *CPU) muli(a, b, d int) { c.register[d] = c.register[a] * b }
func (c *CPU) banr(a, b, d int) { c.register[d] = c.register[a] & c.register[b] }
func (c *CPU) bani(a, b, d int) { c.register[d] = c.register[a] & b }
func (c *CPU) borr(a, b, d int) { c.register[d] = c.register[a] | c.register[b] }
func (c *CPU) bori(a, b, d int) { c.register[d] = c.register[a] | b }
func (c *CPU) setr(a, b, d int) { c.register[d] = c.register[a] }
func (c *CPU) seti(a, b, d int) { c.register[d] = a }

func (c *CPU) gtir(a, b, d int) {
	if a > c.register[b] {
		c.register[d] = 1
	} else {
		c.register[d] = 0
	}
}

func (c *CPU) gtri(a, b, d int) {
	if c.register[a] > b {
		c.register[d] = 1
	} else {
		c.register[d] = 0
	}
}

func (c *CPU) gtrr(a, b, d int) {
	if c.register[a] > c.register[b] {
		c.register[d] = 1
	} else {
		c.register[d] = 0
	}
}

func (c *CPU) eqir(a, b, d int) {
	if a == c.register[b] {
		c.register[d] = 1
	} else {
		c.register[d] = 0
	}
}

func (c *CPU) eqri(a, b, d int) {
	if c.register[a] == b {
		c.register[d] = 1
	} else {
		c.register[d] = 0
	}
}

func (c *CPU) eqrr(a, b, d int) {
	if c.register[a] == c.register[b] {
		c.register[d] = 1
	} else {
		c.register[d] = 0
	}
}

func solve(puzzleInput []string) int {
	cpu := NewCPU(puzzleInput[1:])
	parts := strings.Fields(puzzleInput[0])
	ipReg, _ := strconv.Atoi(parts[1])
	cpu.bindIP(ipReg)

	for cpu.ip >= 0 && cpu.ip < len(cpu.program) {
		if cpu.ip == 28 {
			parts := strings.Fields(cpu.program[28])
			for _, x := range parts[1:3] {
				if x != "0" {
					val, _ := strconv.Atoi(x)
					return cpu.register[val]
				}
			}
		}
		cpu.run()
	}
	return 0
}

func solve2(puzzleInput []string) int {
	potential := make(map[int]bool)
	var last int

	// Parse constants from input
	l6, _ := strconv.Atoi(strings.Fields(puzzleInput[7])[2])
	l7, _ := strconv.Atoi(strings.Fields(puzzleInput[8])[1])
	l8, _ := strconv.Atoi(strings.Fields(puzzleInput[9])[2])
	l10, _ := strconv.Atoi(strings.Fields(puzzleInput[11])[2])
	l11, _ := strconv.Atoi(strings.Fields(puzzleInput[12])[2])
	l12, _ := strconv.Atoi(strings.Fields(puzzleInput[13])[2])
	l19, _ := strconv.Atoi(strings.Fields(puzzleInput[20])[2])

	a, b, c, d := 0, 0, 0, 0
	skipTo8 := false

	for {
		if !skipTo8 {
			c = d | l6
			d = l7
		}
		skipTo8 = false
		b = c & l8
		d += b
		d &= l10
		d *= l11
		d &= l12

		if 256 > c {
			if potential[d] {
				return last
			} else {
				potential[d] = true
				last = d
			}
		} else {
			b = 0
			for {
				a = b + 1
				a *= l19
				if a > c {
					c = b
					skipTo8 = true
					break
				} else {
					b++
					continue
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var puzzleInput []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}

	value := solve(puzzleInput)
	fmt.Printf("The lowest non-negative integer value that causes the program to halt with the fewest instructions is %d.\n", value)

	value2 := solve2(puzzleInput)
	fmt.Printf("The lowest non-negative integer value that causes the program to halt with the most instructions is %d.\n", value2)
}
