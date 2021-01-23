package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println()
	fmt.Println("Day 05, part 1: Doesn't He Have Intern-Elves For This?")

	niceStrings := 0
	for _, textString := range fileLines {
		if hasVovels(textString, 3) && !doesContainStrings(textString) && hasDoubleletter(textString) {
			niceStrings++
		}

	}
	fmt.Println("Number of nice strings:", niceStrings)
	// ------------ PART 2 ------------------------

	fmt.Println()
	fmt.Println("Part 2/")

	niceStrings = 0
	for _, textString := range fileLines {
		if doubleLetterPairExist(textString) && letterWithLetterBetween(textString) {
			niceStrings++
		}
	}
	fmt.Println("Number of nice strings (p2 rules):", niceStrings)

}

// It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
func hasDoubleletter(text string) bool {

	for c := 0; c < len(text)-1; c++ {
		if text[c] == text[c+1] {
			return true
		}
	}
	return false

}

// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
func doesContainStrings(text string) bool {
	if strings.Contains(text, "ab") {
		return true
	} else if strings.Contains(text, "cd") {
		return true
	} else if strings.Contains(text, "pq") {
		return true
	} else if strings.Contains(text, "xy") {
		return true
	}
	return false
}

// It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
func hasVovels(text string, number int) bool {
	numVovels := 0

	numVovels += strings.Count(text, "a")
	numVovels += strings.Count(text, "e")
	numVovels += strings.Count(text, "i")
	numVovels += strings.Count(text, "o")
	numVovels += strings.Count(text, "u")
	if numVovels >= number {
		return true
	}
	return false
}

func doubleLetterPairExist(text string) bool {
	for i := 0; i < len(text)-1; i++ {
		if strings.Contains(text[i+2:], text[i:i+2]) {
			return true
		}
	}
	return false

}

func letterWithLetterBetween(text string) bool {
	for c := 0; c < len(text)-2; c++ {
		if text[c] == text[c+2] {
			return true
		}
	}
	return false

}
