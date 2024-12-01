package main

import (
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var s scanner.Scanner
	s.Init(file)
	s.Mode = scanner.ScanInts
	s.Whitespace ^= 1 << '\n'

	// Flag that determines card section
	isWinningNumbersSection := true

	// Map that stores winning numbers of each card
	winningNumbers := make(map[int]bool)

	// Map that pairs card indexes with number of corresponding cards
	cardNumbers := make(map[int]int)

	// Number of matches for each card
	var matches int

	// currentCard sores the current card index minus 1
	currentCard := 0

	var resultPartOne, resultPartTwo int

	// Iterate through the file
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case scanner.Int:
			// Skip the card number in "Card <card number>:"
			if s.Peek() == ':' {
				continue
			}

			// Parse the number and add it to the winning numbers slice if it is
			// in the winning numbers section, increment number of matches otherwise
			if num, _ := strconv.Atoi(s.TokenText()); isWinningNumbersSection {
				winningNumbers[num] = true
			} else if winningNumbers[num] {
				matches++
			}
		case '|':
			isWinningNumbersSection = false
		case '\n':
			resultPartOne += (1 << matches) >> 1
			resultPartTwo += cardNumbers[currentCard] + 1

			for match := 1; match <= matches; match++ {
				cardNumbers[currentCard+match] += cardNumbers[currentCard] + 1
			}

			matches = 0
			isWinningNumbersSection = true
			clear(winningNumbers)
			delete(cardNumbers, currentCard)

			currentCard++
		}
	}

	fmt.Printf("Part one: %d\nPart two: %d\n", resultPartOne, resultPartTwo)
}
