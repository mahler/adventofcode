package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	OpLT          = "<"
	OpGT          = ">"
	StartWorkflow = "in"
	Accept        = "A"
	Reject        = "R"
	ValLowerBound = 1
	ValUpperBound = 4001
)

var Categories = []string{"x", "m", "a", "s"}

type Rule struct {
	lhsField string
	opChar   string
	rhs      int
	action   string
}

type Workflow struct {
	name  string
	rules []Rule
}

type Constraint struct {
	op  string
	val int
}

func (c Constraint) negate() Constraint {
	if c.op == OpLT {
		return Constraint{OpGT, c.val - 1}
	}
	return Constraint{OpLT, c.val + 1}
}

type CriteriaSet struct {
	vars map[string][2]int
}

func NewCriteriaSet() *CriteriaSet {
	cs := &CriteriaSet{
		vars: make(map[string][2]int),
	}
	for _, c := range Categories {
		cs.vars[c] = [2]int{ValLowerBound, ValUpperBound}
	}
	return cs
}

func (cs *CriteriaSet) copy() *CriteriaSet {
	newCS := &CriteriaSet{
		vars: make(map[string][2]int),
	}
	for k, v := range cs.vars {
		newCS.vars[k] = v
	}
	return newCS
}

func (cs *CriteriaSet) apply(variable string, constraint Constraint) {
	current := cs.vars[variable]
	low, high := current[0], current[1]

	if constraint.op == OpLT {
		cs.vars[variable] = [2]int{low, min(constraint.val, high)}
	} else {
		cs.vars[variable] = [2]int{max(low, constraint.val+1), high}
	}
}

func (cs *CriteriaSet) combinations() int64 {
	result := int64(1)
	for _, bounds := range cs.vars {
		diff := max(0, bounds[1]-bounds[0])
		result *= int64(diff)
	}
	return result
}

func readWorkflow(s string) Workflow {
	rulesBegin := strings.Index(s, "{")
	name := s[:rulesBegin]
	rulesStr := s[rulesBegin+1 : len(s)-1]
	rules := make([]Rule, 0)

	for _, r := range strings.Split(rulesStr, ",") {
		if strings.Contains(r, ":") {
			parts := strings.Split(r, ":")
			pred, action := parts[0], parts[1]
			lhs := string(pred[0])
			op := string(pred[1])
			rhs, _ := strconv.Atoi(pred[2:])
			rules = append(rules, Rule{lhs, op, rhs, action})
		} else {
			rules = append(rules, Rule{"", "", 0, r})
		}
	}

	return Workflow{name, rules}
}

func readRegisters(s string) map[string]int {
	s = strings.Trim(s, "{}")
	vals := make(map[string]int)
	for _, item := range strings.Split(s, ",") {
		parts := strings.Split(item, "=")
		v, _ := strconv.Atoi(parts[1])
		vals[parts[0]] = v
	}
	return vals
}

func evalRule(rule Rule, registers map[string]int) bool {
	if rule.lhsField == "" {
		return true
	}

	lhsVal := registers[rule.lhsField]
	if rule.opChar == OpLT {
		return lhsVal < rule.rhs
	}
	return lhsVal > rule.rhs
}

func evalRating(registers map[string]int, workflows map[string]Workflow) bool {
	w := workflows[StartWorkflow]
	for {
		var action string
		for _, rule := range w.rules {
			if evalRule(rule, registers) {
				action = rule.action
				break
			}
		}

		if action == Accept {
			return true
		} else if action == Reject {
			return false
		} else {
			w = workflows[action]
		}
	}
}

func countCombinations(workflows map[string]Workflow) int64 {
	var count func(w Workflow, state *CriteriaSet) int64
	count = func(w Workflow, state *CriteriaSet) int64 {
		var sum int64
		for _, rule := range w.rules {
			substate := state.copy()

			if rule.lhsField != "" {
				cons := Constraint{rule.opChar, rule.rhs}
				substate.apply(rule.lhsField, cons)
				state.apply(rule.lhsField, cons.negate())
			}

			if rule.action == Accept {
				sum += substate.combinations()
			} else if rule.action == Reject {
				continue
			} else {
				sum += count(workflows[rule.action], substate)
			}
		}
		return sum
	}

	return count(workflows[StartWorkflow], NewCriteriaSet())
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	readingWorkflows := true
	workflows := make(map[string]Workflow)
	var registers []map[string]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingWorkflows = false
			continue
		}

		if readingWorkflows {
			w := readWorkflow(line)
			workflows[w.name] = w
		} else {
			registers = append(registers, readRegisters(line))
		}
	}

	// Part 1
	var sum int
	for _, r := range registers {
		if evalRating(r, workflows) {
			for _, v := range r {
				sum += v
			}
		}
	}
	fmt.Println("Part 1: Sort through all of the parts you've been given; what do you get if you add together")
	fmt.Println("all of the rating numbers for all of the parts that ultimately get accepted?")
	fmt.Println(sum)

	// Part 2
	fmt.Println()
	fmt.Println("Part 2: How many distinct combinations of ratings will be accepted by the Elves' workflows?")
	fmt.Println(countCombinations(workflows))
}
