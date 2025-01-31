package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Regs [6]int
type OpCode struct {
	code int
	a    int
	b    int
	c    int
}

func eval(op *OpCode, regs *Regs) {
	code := op.code
	switch code {
	case 0x0:
		regs[op.c] = regs[op.a] + regs[op.b]
	case 0x1:
		regs[op.c] = regs[op.a] + op.b
	case 0x2:
		regs[op.c] = regs[op.a] * regs[op.b]
	case 0x3:
		regs[op.c] = regs[op.a] * op.b
	case 0x4:
		regs[op.c] = regs[op.a] & regs[op.b]
	case 0x5:
		regs[op.c] = regs[op.a] & op.b
	case 0x6:
		regs[op.c] = regs[op.a] | regs[op.b]
	case 0x7:
		regs[op.c] = regs[op.a] | op.b
	case 0x8:
		regs[op.c] = regs[op.a]
	case 0x9:
		regs[op.c] = op.a
	case 0xA:
		if op.a > regs[op.b] {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case 0xB:
		if regs[op.a] > op.b {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case 0xC:
		if regs[op.a] > regs[op.b] {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case 0xD:
		if op.a == regs[op.b] {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case 0xE:
		if regs[op.a] == op.b {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	case 0xF:
		if regs[op.a] == regs[op.b] {
			regs[op.c] = 1
		} else {
			regs[op.c] = 0
		}
	default:
		panic("invalid opcode")
	}
}

func parseInstr(s string) int {
	switch s {
	case "addr":
		return 0x0
	case "addi":
		return 0x1
	case "mulr":
		return 0x2
	case "muli":
		return 0x3
	case "banr":
		return 0x4
	case "bani":
		return 0x5
	case "borr":
		return 0x6
	case "bori":
		return 0x7
	case "setr":
		return 0x8
	case "seti":
		return 0x9
	case "gtir":
		return 0xA
	case "gtri":
		return 0xB
	case "gtrr":
		return 0xC
	case "eqir":
		return 0xD
	case "eqri":
		return 0xE
	case "eqrr":
		return 0xF
	default:
		panic("invalid instruction")
	}
}

func findNumberInLine(line string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)
	if len(matches) < 2 {
		return 0, fmt.Errorf("not enough numbers found in line")
	}
	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	return num, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instrs []OpCode

	// Read first line for IP register
	scanner.Scan()
	ipParts := strings.Fields(scanner.Text())
	ipr, err := strconv.Atoi(ipParts[1])
	if err != nil {
		panic(err)
	}

	// Read instructions
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		code := parseInstr(fields[0])
		a, _ := strconv.Atoi(fields[1])
		b, _ := strconv.Atoi(fields[2])
		c, _ := strconv.Atoi(fields[3])
		instrs = append(instrs, OpCode{code, a, b, c})
	}

	// Part 1
	regs := Regs{}
	ip := 0
	for ip < len(instrs) {
		eval(&instrs[ip], &regs)
		regs[ipr]++
		ip = regs[ipr]
	}
	fmt.Println("Part 1: What value is left in register 0 when the background process halts?")
	fmt.Println(regs[0])

	// Part 2
	file, _ = os.Open("input.txt")
	defer file.Close()
	scanner = bufio.NewScanner(file)
	var inputLines []string
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading input: %v\n", err)
		os.Exit(1)
	}

	// Extract numbers from lines 22 and 24
	if len(inputLines) < 25 {
		fmt.Fprintf(os.Stderr, "not enough input lines\n")
		os.Exit(1)
	}
	a, err := findNumberInLine(inputLines[22])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing line 22: %v\n", err)
		os.Exit(1)
	}
	b, err := findNumberInLine(inputLines[24])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing line 24: %v\n", err)
		os.Exit(1)
	}

	numberToFactorize := float64(10551236 + a*22 + b)

	// Create map for factors
	factors := make(map[float64]int)

	// Find prime factors
	possiblePrimeDivisor := float64(2)
	for possiblePrimeDivisor*possiblePrimeDivisor <= numberToFactorize {
		for math.Mod(numberToFactorize, possiblePrimeDivisor) == 0 {
			numberToFactorize /= possiblePrimeDivisor
			factors[possiblePrimeDivisor]++
		}
		possiblePrimeDivisor++
	}
	if numberToFactorize > 1 {
		factors[numberToFactorize]++
	}

	// Calculate sum of divisors
	sumOfDivisors := float64(1)
	for primeFactor, power := range factors {
		numerator := math.Pow(primeFactor, float64(power+1)) - 1
		denominator := primeFactor - 1
		sumOfDivisors *= numerator / denominator
	}

	fmt.Println()
	fmt.Println("Part 2: What value is left in register 0 when this new background process halts?")
	fmt.Printf("%.0f\n", sumOfDivisors)
}
