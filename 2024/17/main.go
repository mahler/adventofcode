package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func run(prog []int, a int) []int {
	ip, b, c := 0, 0, 0
	out := []int{}

	for ip >= 0 && ip < len(prog) {
		lit := prog[ip+1]
		combo := []int{0, 1, 2, 3, a, b, c, 99999}[lit]

		switch prog[ip] {
		case 0: // adv
			a = a >> combo
		case 1: // bxl
			b = b ^ lit
		case 2: // bst
			b = combo % 8
		case 3: // jnz
			if a != 0 {
				ip = lit - 2
			}
		case 4: // bxc
			b = b ^ c
		case 5: // out
			out = append(out, combo%8)
		case 6: // bdv
			b = a >> combo
		case 7: // cdv
			c = a >> combo
		}

		ip += 2
	}

	return out
}

func reverseIntSlice(slice []int) []int {
	reversed := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		reversed[i] = slice[len(slice)-1-i]
	}
	return reversed
}

func findA(prog []int, a int, depth int) int {
	target := reverseIntSlice(prog)

	if depth == len(target) {
		return a
	}

	for i := 0; i < 8; i++ {
		output := run(prog, a*8+i)
		if len(output) > 0 && output[0] == target[depth] {
			result := findA(prog, a*8+i, depth+1)
			if result != 0 {
				return result
			}
		}
	}

	return 0
}

func main() {
	// Read input file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Use regex to extract numbers
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(string(content), -1)

	// Convert matches to integers
	nums := make([]int, len(matches))
	for i, match := range matches {
		nums[i], _ = strconv.Atoi(match)
	}

	// Extract initial values and program
	a := nums[0]
	prog := nums[3:]

	// Part 1
	part1Output := run(prog, a)
	part1Str := make([]string, len(part1Output))
	for i, n := range part1Output {
		part1Str[i] = strconv.Itoa(n)
	}
	fmt.Println("Part 1: Once it halts, what do you get if you use commas to join the values it output into a single string?")
	fmt.Println(strings.Join(part1Str, ","))

	// Part 2
	part2Result := findA(prog, 0, 0)
	fmt.Println()
	fmt.Println("Part 2: What is the lowest positive initial value for register A that causes the program to output a copy of itself?")
	fmt.Println(part2Result)
}
