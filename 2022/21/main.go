package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	op     func(int, int) int
	left   string
	right  string
	number int
	isNum  bool
}

type Operations map[string]Operation

func strToOp(s string) func(int, int) int {
	switch s {
	case "+":
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	case "/":
		return func(a, b int) int { return a / b }
	default:
		return nil
	}
}

func parseLine(line string) (string, Operation) {
	parts := strings.Fields(strings.ReplaceAll(line, ":", ""))
	target := parts[0]

	if num, err := strconv.Atoi(parts[1]); err == nil {
		return target, Operation{number: num, isNum: true}
	}

	return target, Operation{
		op:    strToOp(parts[2]),
		left:  parts[1],
		right: parts[3],
		isNum: false,
	}
}

func compute(ops Operations, target string) int {
	rule := ops[target]
	if rule.isNum {
		return rule.number
	}
	return rule.op(compute(ops, rule.left), compute(ops, rule.right))
}

type computeFunc func(int) int

func isAssociative(op func(int, int) int) bool {
	return op(1, 2) == op(2, 1)
}

func inverseOp(op func(int, int) int) func(int, int) int {
	switch {
	case op(4, 2) == 6:
		return func(a, b int) int { return a - b }
	case op(6, 2) == 4:
		return func(a, b int) int { return a + b }
	case op(6, 2) == 12:
		return func(a, b int) int { return a / b }
	case op(12, 2) == 6:
		return func(a, b int) int { return a * b }
	default:
		return nil
	}
}

func compute2(ops Operations, target string) interface{} {
	rule := ops[target]
	switch {
	case target == "humn":
		return computeFunc(func(x int) int { return x })
	case rule.isNum:
		return rule.number
	default:
		a := compute2(ops, rule.left)
		b := compute2(ops, rule.right)

		aFunc, aIsFunc := a.(computeFunc)
		bFunc, bIsFunc := b.(computeFunc)

		switch {
		case !aIsFunc && !bIsFunc:
			return rule.op(a.(int), b.(int))
		case target == "root":
			if aIsFunc {
				return aFunc(b.(int))
			}
			return bFunc(a.(int))
		case aIsFunc:
			return computeFunc(func(x int) int {
				return aFunc(inverseOp(rule.op)(x, b.(int)))
			})
		case isAssociative(rule.op):
			return computeFunc(func(x int) int {
				return bFunc(inverseOp(rule.op)(x, a.(int)))
			})
		default:
			return computeFunc(func(x int) int {
				return bFunc(rule.op(a.(int), x))
			})
		}
	}
}

func solvePart1(ops Operations) int {
	return compute(ops, "root")
}

func solvePart2(ops Operations) int {
	return compute2(ops, "root").(int)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ops := make(Operations)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		target, op := parseLine(line)
		ops[target] = op
	}

	fmt.Println("Part 1: What number will the monkey named root yell?")
	fmt.Println(solvePart1(ops))

	fmt.Println()
	fmt.Println("Part 2: What number do you yell to pass root's equality test?")
	fmt.Println(solvePart2(ops))
}
