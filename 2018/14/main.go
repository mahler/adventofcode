package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = 768071

func main() {
	scores := []byte{'3', '7'}
	elf1, elf2 := 0, 1 // Indexes of the two elves

	// Set a reasonable limit for the scoreboard length
	const maxScores = 50000000

	for len(scores) < maxScores {
		// Calculate the sum of the current recipes of the two elves
		sum := int(scores[elf1]-'0') + int(scores[elf2]-'0')

		// Append each digit of the sum to the scoreboard
		if sum >= 10 {
			scores = append(scores, byte('1'))
		}
		scores = append(scores, byte('0'+sum%10))

		elf1 = (elf1 + 1 + int(scores[elf1]-'0')) % len(scores)
		elf2 = (elf2 + 1 + int(scores[elf2]-'0')) % len(scores)
	}

	part1 := string(scores[input : input+10])
	fmt.Println("Part 1: What are the scores of the ten recipes immediately after the number of recipes in your puzzle input?")
	fmt.Println(part1)

	part2 := strings.Index(string(scores), strconv.Itoa(input))
	fmt.Println()
	fmt.Println("Part 2: How many recipes appear on the scoreboard to the left of the score sequence in your puzzle input?")
	fmt.Println(part2)
}
