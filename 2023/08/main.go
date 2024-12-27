package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	ll := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	inst := strings.Split(ll[0], "")
	conn := make(map[string][2]string)

	for _, l := range strings.Split(ll[1], "\n") {
		parts := strings.Split(l, " ")
		a := parts[0]
		b := strings.Split(strings.Split(l, "(")[1], ",")[0]
		c := strings.Split(parts[3], ")")[0]
		conn[a] = [2]string{b, c}
	}

	pos := "AAA"
	idx := 0
	for pos != "ZZZ" {
		d := inst[idx%len(inst)]
		if d == "L" {
			pos = conn[pos][0]
		} else {
			pos = conn[pos][1]
		}
		idx++
	}
	fmt.Println("part 1: How many steps are required to reach ZZZ?")
	fmt.Println(idx)

	ret := int64(1)
	for start := range conn {
		if strings.HasSuffix(start, "A") {
			ret = lcm(ret, solveSteps(start, inst, conn))
		}
	}
	fmt.Println()
	fmt.Println("Part 2: How many steps does it take before you're only on nodes that end with Z?")
	fmt.Println(ret)
}

func solveSteps(start string, inst []string, conn map[string][2]string) int64 {
	pos := start
	idx := 0
	for !strings.HasSuffix(pos, "Z") {
		d := inst[idx%len(inst)]
		if d == "L" {
			pos = conn[pos][0]
		} else {
			pos = conn[pos][1]
		}
		idx++
	}
	return int64(idx)
}

func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
