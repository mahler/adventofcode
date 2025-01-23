package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Value struct {
	isReg  bool
	regIdx int
	val    int64
}

func (v Value) eval(cpu *Cpu) int64 {
	if v.isReg {
		return cpu.regs[v.regIdx]
	}
	return v.val
}

type Op interface {
	execute(cpu *Cpu, other *Cpu) bool
}

type Snd struct {
	value Value
}

func (s Snd) execute(cpu *Cpu, other *Cpu) bool {
	v := s.value.eval(cpu)
	cpu.lastSent = v
	other.queue = append(other.queue, v)
	cpu.sent++
	cpu.pc++
	return true
}

type Set struct {
	reg   int
	value Value
}

func (s Set) execute(cpu *Cpu, other *Cpu) bool {
	cpu.regs[s.reg] = s.value.eval(cpu)
	cpu.pc++
	return true
}

type Add struct {
	reg   int
	value Value
}

func (a Add) execute(cpu *Cpu, other *Cpu) bool {
	cpu.regs[a.reg] += a.value.eval(cpu)
	cpu.pc++
	return true
}

type Mul struct {
	reg   int
	value Value
}

func (m Mul) execute(cpu *Cpu, other *Cpu) bool {
	cpu.regs[m.reg] *= m.value.eval(cpu)
	cpu.pc++
	return true
}

type Mod struct {
	reg   int
	value Value
}

func (m Mod) execute(cpu *Cpu, other *Cpu) bool {
	cpu.regs[m.reg] %= m.value.eval(cpu)
	cpu.pc++
	return true
}

type Rcv struct {
	reg int
}

func (r Rcv) execute(cpu *Cpu, other *Cpu) bool {
	if cpu.firstNonZeroRecv == 0 && cpu.regs[r.reg] != 0 {
		cpu.firstNonZeroRecv = cpu.lastSent
	}
	if len(cpu.queue) > 0 {
		cpu.regs[r.reg] = cpu.queue[0]
		cpu.queue = cpu.queue[1:]
		cpu.pc++
		return true
	}
	return false
}

type Jgz struct {
	test   Value
	offset Value
}

func (j Jgz) execute(cpu *Cpu, other *Cpu) bool {
	if j.test.eval(cpu) > 0 {
		cpu.pc += int(j.offset.eval(cpu))
	} else {
		cpu.pc++
	}
	return true
}

type Cpu struct {
	regs             [26]int64
	mem              []Op
	pc               int
	queue            []int64
	sent             int
	lastSent         int64
	firstNonZeroRecv int64
}

func newCpu(mem []Op, id int64) *Cpu {
	cpu := &Cpu{
		mem: mem,
		pc:  0,
	}
	cpu.regs[int('p'-'a')] = id
	return cpu
}

func (cpu *Cpu) runDual(other *Cpu, first bool) {
	cpu.run(other, first)
}

func (cpu *Cpu) run(other *Cpu, first bool) {
	for cpu.pc < len(cpu.mem) {
		advance := cpu.mem[cpu.pc].execute(cpu, other)
		if !advance {
			if !first {
				break
			}
			other.run(cpu, false)
			if !cpu.mem[cpu.pc].execute(cpu, other) {
				break
			}
		}
	}
}

func parseValue(s string) Value {
	if len(s) == 1 && s[0] >= 'a' && s[0] <= 'z' {
		return Value{isReg: true, regIdx: int(s[0] - 'a')}
	}
	val, _ := strconv.ParseInt(s, 10, 64)
	return Value{isReg: false, val: val}
}

func parseOp(s string) Op {
	parts := strings.Fields(s)
	switch parts[0] {
	case "snd":
		return Snd{value: parseValue(parts[1])}
	case "set":
		return Set{reg: int(parts[1][0] - 'a'), value: parseValue(parts[2])}
	case "add":
		return Add{reg: int(parts[1][0] - 'a'), value: parseValue(parts[2])}
	case "mul":
		return Mul{reg: int(parts[1][0] - 'a'), value: parseValue(parts[2])}
	case "mod":
		return Mod{reg: int(parts[1][0] - 'a'), value: parseValue(parts[2])}
	case "rcv":
		return Rcv{reg: int(parts[1][0] - 'a')}
	case "jgz":
		return Jgz{test: parseValue(parts[1]), offset: parseValue(parts[2])}
	default:
		panic("unknown instruction")
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var mem []Op
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mem = append(mem, parseOp(scanner.Text()))
	}

	cpu1 := newCpu(mem, 0)
	cpu2 := newCpu(mem, 1)
	cpu1.runDual(cpu2, true)

	fmt.Println("Part 1: What is the value of the recovered frequency (the value of the most recently played sound)")
	fmt.Println("the first time a rcv instruction is executed with a non-zero value?")
	fmt.Println(cpu1.firstNonZeroRecv)

	fmt.Println()
	fmt.Println("Part 2: Once both of your programs have terminated (regardless of what caused them to do so),")
	fmt.Println("how many times did program 1 send a value?")
	fmt.Println(cpu2.sent)
}
