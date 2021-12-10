package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("Syntax Scoring")
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	// Setup
	fileRows := string(fileContent)
	fileLines := strings.Split(strings.TrimSpace(fileRows), "\n")

	starters := map[rune]rune{'(': ')', '<': '>', '[': ']', '{': '}'}
	rev := map[rune]rune{')': '(', '>': '<', ']': '[', '}': '{'}
	points := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}

	syntaxErrorScore := 0
	for _, line := range fileLines {
		var q []rune
		for _, c := range line {
			if _, ok := starters[c]; ok {
				q = append(q, c)
			} else {
				if q[len(q)-1] != rev[c] {
					syntaxErrorScore += points[c]
					break
				}
				q = q[0 : len(q)-1]
			}
		}
	}
	fmt.Println()
	fmt.Println("2021-10: Part 1/")
	fmt.Println("What is the total syntax error score for those errors?", syntaxErrorScore)

	// Part 2
	var completeScores []int
	points2 := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	for _, line := range fileLines {
		var bad bool
		var q []rune
		for _, c := range line {
			if _, ok := starters[c]; ok {
				q = append(q, c)
			} else {
				if q[len(q)-1] != rev[c] {
					bad = true
					break
				}
				q = q[0 : len(q)-1]
			}
		}
		if !bad {
			var score2 int
			for i := len(q) - 1; i >= 0; i-- {
				score2 = score2*5 + points2[q[i]]
			}
			completeScores = append(completeScores, score2)
		}
	}
	sort.Ints(completeScores)

	fmt.Println()
	fmt.Println("Part 2/")
	fmt.Println("What is the middle score?", completeScores[len(completeScores)/2])

}
