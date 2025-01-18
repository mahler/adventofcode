package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// modulo implements a proper modulo operation that matches Fortran's MODULO
func modulo(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

// createChain initializes a chain of numbers from 0 to 255
func createChain() []int {
	chain := make([]int, 256)
	for i := range chain {
		chain[i] = i
	}
	return chain
}

// reverseSubchain reverses a portion of the chain starting from currentPos with given length
func reverseSubchain(chain []int, currentPos, length int) {
	if length > len(chain) {
		return
	}

	// Create temporary slice for reversed elements
	subchain := make([]int, length)
	for i := 0; i < length; i++ {
		pos := modulo(currentPos+length-1-i, len(chain))
		subchain[i] = chain[pos]
	}

	// Put reversed elements back into chain
	for i := 0; i < length; i++ {
		pos := modulo(currentPos+i, len(chain))
		chain[pos] = subchain[i]
	}
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Part 1
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	// Parse comma-separated numbers
	strNums := strings.Split(input, ",")
	instructions := make([]int, len(strNums))
	for i, s := range strNums {
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		instructions[i] = num
	}

	// Process Part 1
	chain1 := createChain()
	currentPos := 0
	skipSize := 0

	for _, length := range instructions {
		reverseSubchain(chain1, currentPos, length)
		currentPos = modulo(currentPos+length+skipSize, len(chain1))
		skipSize++
	}

	fmt.Println("Part1: What is the result of multiplying the first two numbers in the list?")
	fmt.Println(chain1[0] * chain1[1])

	// Part 2
	file.Seek(0, 0) // Rewind file
	scanner = bufio.NewScanner(file)
	scanner.Scan()
	input = scanner.Text()

	// Create ASCII instructions
	instructions = make([]int, len(input)+5)
	for i := 0; i < len(input); i++ {
		instructions[i] = int(input[i])
	}
	// Append standard suffix
	suffix := []int{17, 31, 73, 47, 23}
	copy(instructions[len(input):], suffix)

	// Process Part 2
	chain2 := createChain()
	currentPos = 0
	skipSize = 0

	// Perform 64 rounds
	for round := 0; round < 64; round++ {
		for _, length := range instructions {
			reverseSubchain(chain2, currentPos, length)
			currentPos = modulo(currentPos+length+skipSize, len(chain2))
			skipSize++
		}
	}

	// Calculate dense hash
	denseHash := make([]int, 16)
	for i := 0; i < 16; i++ {
		result := chain2[i*16]
		for j := 1; j < 16; j++ {
			result ^= chain2[i*16+j]
		}
		denseHash[i] = result
	}

	// Convert to hex string
	fmt.Println()
	fmt.Println("Part2: Treating your puzzle input as a string of ASCII characters,")
	fmt.Println("what is the Knot Hash of your puzzle input?")
	sPart2 := ""
	for _, val := range denseHash {
		sPart2 = sPart2 + fmt.Sprintf("%02x", val)
	}
	fmt.Println(sPart2)
}
