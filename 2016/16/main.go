package main

import (
	"fmt"
)

func main() {
	input := "01111001100111011"

	data := make([]bool, len(input))
	for i, c := range input {
		if c == '1' {
			data[i] = true
		} else {
			data[i] = false
		}
	}
	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Part One: Dragon Checksum")
	fmt.Println("The first disk you have to fill has length 272.")
	fmt.Println("Using the initial state in your puzzle input, what is the correct checksum?")
	fmt.Println(generateDataAndCalculateChecksum(data, 272))

	fmt.Println()
	fmt.Println("Part 2/")
	fmt.Println("The second disk you have to fill has length 35651584. Again using the initial")
	fmt.Println("state in your puzzle input, what is the correct checksum for this disk?")
	fmt.Println(generateDataAndCalculateChecksum(data, 35651584))
}

func generateDataAndCalculateChecksum(base []bool, target int) string {
	copyB := make([]bool, len(base))

	// Make a copy of "a"; call this copy "b".
	copy(copyB, base)

	for len(copyB) < target {
		next := make([]bool, 2*len(copyB)+1)
		copy(next, copyB)
		next[len(copyB)] = false
		for i := 0; i < len(copyB); i++ {
			next[2*len(copyB)-i] = !copyB[i]
		}
		copyB = next
	}

	copyB = copyB[0:target]

	for len(copyB)%2 == 0 {
		next := make([]bool, len(copyB)/2)
		for i := 0; i < len(next); i++ {
			next[i] = (copyB[2*i+0] == copyB[2*i+1])
		}
		copyB = next
	}

	output := make([]byte, len(copyB))
	for i, d := range copyB {
		if d {
			output[i] = '1'
		} else {
			output[i] = '0'
		}
	}

	return string(output)
}
