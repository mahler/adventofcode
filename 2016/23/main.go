package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getVal(regs map[string]int, x string) int {
	if isNum(x) {
		val, _ := strconv.Atoi(x)
		return val
	}
	return regs[x]
}

func toggle(line string) string {
	parts := strings.Split(line, " ")
	instr := parts[0]
	switch instr {
	case "inc":
		return "dec " + strings.Join(parts[1:], " ")
	case "dec", "tgl":
		return "inc " + strings.Join(parts[1:], " ")
	case "jnz":
		return "cpy " + strings.Join(parts[1:], " ")
	case "cpy":
		return "jnz " + strings.Join(parts[1:], " ")
	default:
		panic("Invalid instruction")
	}
}

func run(lines []string, part2 bool) int {
	pc := 0
	regs := map[string]int{
		"a": 7,
		"b": 0,
		"c": 0,
		"d": 0,
	}

	if part2 {
		regs["a"] = 12
	}

	for pc < len(lines) {
		line := lines[pc]
		parts := strings.SplitN(line, " ", 2)
		instr, args := parts[0], parts[1]

		switch instr {
		case "cpy":
			argParts := strings.Split(args, " ")
			b, c := argParts[0], argParts[1]
			if _, ok := regs[c]; ok {
				regs[c] = getVal(regs, b)
			} else {
				fmt.Println("invalid")
			}
			pc++

		case "inc":
			if _, ok := regs[args]; ok {
				// Peephole optimize inc/dec/jnz loops
				if pc+3 < len(lines) && pc-1 >= 0 &&
					strings.HasPrefix(lines[pc-1], "cpy ") &&
					strings.HasPrefix(lines[pc+1], "dec") &&
					strings.HasPrefix(lines[pc+2], "jnz") &&
					strings.HasPrefix(lines[pc+3], "dec") &&
					strings.HasPrefix(lines[pc+4], "jnz") {

					incOp := args
					cpyParts := strings.Split(lines[pc-1], " ")[1:]
					cpySrc, cpyDest := cpyParts[0], cpyParts[1]
					dec1Op := strings.Split(lines[pc+1], " ")[1]
					jnz1Parts := strings.Split(lines[pc+2], " ")[1:]
					jnz1Cond, jnz1Off := jnz1Parts[0], jnz1Parts[1]
					dec2Op := strings.Split(lines[pc+3], " ")[1]
					jnz2Parts := strings.Split(lines[pc+4], " ")[1:]
					jnz2Cond, jnz2Off := jnz2Parts[0], jnz2Parts[1]

					if cpyDest == dec1Op && dec1Op == jnz1Cond &&
						dec2Op == jnz2Cond &&
						jnz1Off == "-2" && jnz2Off == "-5" {
						regs[incOp] += getVal(regs, cpySrc) * getVal(regs, dec2Op)
						regs[dec1Op] = 0
						regs[dec2Op] = 0
						pc += 5
						continue
					}
				}
				regs[args]++
			}
			pc++

		case "dec":
			if _, ok := regs[args]; ok {
				regs[args]--
			}
			pc++

		case "jnz":
			argParts := strings.Split(args, " ")
			b, c := argParts[0], argParts[1]
			if getVal(regs, b) != 0 {
				pc += getVal(regs, c)
			} else {
				pc++
			}

		case "tgl":
			x := getVal(regs, args)
			idx := pc + x
			if idx >= 0 && idx < len(lines) {
				lines[idx] = toggle(lines[idx])
			}
			pc++

		default:
			panic("Invalid instruction")
		}
	}

	return regs["a"]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	fmt.Println("part 1: What value should be sent to the safe?")
	part1result := run(append([]string(nil), lines...), false)
	fmt.Println(part1result)

	fmt.Println()
	fmt.Println("part 2: Anyway, what value should actually be sent to the safe?")
	part2result := run(append([]string(nil), lines...), true)
	fmt.Println(part2result)
}
