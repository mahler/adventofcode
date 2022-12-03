package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	//	partOne(fileLines)
	points := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	win := map[int]int{1: 3, 2: 1, 3: 2}
	score := 0

	for _, line := range fileLines {
		split := strings.Split(line, " ")
		them, you := points[split[0]], points[split[1]]
		score += you
		if win[you] == them {
			score += 6
		} else if them == you {
			score += 3
		}
	}

	fmt.Println()
	fmt.Println("Day 2: Rock Paper Scissors")
	fmt.Println("What would your total score be if everything goes exactly according to your strategy guide?")
	fmt.Println(score)

	// Part Two
	p2points := map[string]int{"A": 1, "B": 2, "C": 3, "X": 0, "Y": 3, "Z": 6}
	p2win := map[int]int{1: 2, 2: 3, 3: 1}
	p2loose := map[int]int{1: 3, 2: 1, 3: 2}
	score = 0

	for _, line := range fileLines {
		split := strings.Split(line, " ")
		them, you := p2points[split[0]], p2points[split[1]]
		score += you

		switch you {
		case 0:
			score += p2loose[them]
		case 3:
			score += them
		case 6:
			score += p2win[them]
		}
	}
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("Following the Elf's instructions for the second column,")
	fmt.Println("what would your total score be if everything goes exactly according to your strategy guide?")
	fmt.Println(score)
}
