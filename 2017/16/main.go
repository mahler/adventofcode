package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func spin(s string, n int) string {
	return s[len(s)-n:] + s[:len(s)-n]
}

func partner(s string, a, b byte) string {
	temp := strings.ReplaceAll(s, string(a), "x")
	temp = strings.ReplaceAll(temp, string(b), "y")
	temp = strings.ReplaceAll(temp, "x", string(b))
	temp = strings.ReplaceAll(temp, "y", string(a))
	return temp
}

func exchange(s string, a, b int) string {
	chars := []byte(s)
	chars[a], chars[b] = chars[b], chars[a]
	return string(chars)
}

func dance(s string, moves []string) string {
	result := s
	for _, m := range moves {
		switch m[0] {
		case 's':
			n, _ := strconv.Atoi(m[1:])
			result = spin(result, n)
		case 'x':
			positions := strings.Split(m[1:], "/")
			a, _ := strconv.Atoi(positions[0])
			b, _ := strconv.Atoi(positions[1])
			result = exchange(result, a, b)
		case 'p':
			result = partner(result, m[1], m[3])
		}
	}
	return result
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	moves := strings.Split(strings.TrimSpace(string(data)), ",")
	alpha := "abcdefghijklmnop"

	// Part A
	result := dance(alpha, moves)
	fmt.Println("Part 1: In what order are the programs standing after their dance?")
	fmt.Println(result)

	// Part B
	i := 1
	for result != alpha {
		result = dance(result, moves)
		i++
	}

	j := 1000000000 % i
	result = alpha
	for k := 0; k < j; k++ {
		result = dance(result, moves)
	}

	fmt.Println()
	fmt.Println("Part 2: In what order are the programs standing after their billion dances?")
	fmt.Println(result)
}
