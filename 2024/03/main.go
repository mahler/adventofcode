package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed puzzle.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	part1result := 0
	part2result := 0
	enabled := true
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	for _, line := range lines {
		for _, match := range r.FindAllString(line, -1) {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else {
				s := strings.Split(match[4:len(match)-1], ",")
				x, _ := strconv.Atoi(s[0])
				y, _ := strconv.Atoi(s[1])
				part1result += x * y
				if enabled {
					part2result += x * y
				}
			}
		}
	}

	fmt.Println("Part1 : What do you get if you add up all of the results of the multiplications?")
	fmt.Println(part1result)
	fmt.Println()
	fmt.Println("Part2 : what do you get if you add up all of the results of just the enabled multiplications?")
	fmt.Println(part2result)
}
