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
		return
	}

	fileContent := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("DAY02, Part 1: Inventory Management System")
	fmt.Println("Total puzzle input: ", len(fileContent))

	twoLetters := 0
	threeLetters := 0

	for _, boxid := range fileContent {
		// that have an ID containing exactly two of any letter
		if hasNumberLetters(boxid, 2) {
			twoLetters++
		}
		// and then separately counting those with exactly three of any letter.
		if hasNumberLetters(boxid, 3) {
			threeLetters++
		}

	}

	fmt.Println("Boxes with two letters:", twoLetters)
	fmt.Println("Boxes with three letters:", threeLetters)
	fmt.Println("Checksum:", twoLetters*threeLetters)
	// ------------------------------------------------------------------------
	fmt.Println()
	fmt.Println("Part 2")
	seen := make(map[string]bool)
	for _, boxid := range fileContent {
		// For every letter position in boxid
		for i := 0; i < len(boxid); i++ {
			// Create a map key with the truncated version
			truncated := boxid[:i] + "_" + boxid[i+1:]
			// If map key exists we found the match
			if seen[truncated] {
				common := strings.Replace(truncated, "_", "", 1)
				fmt.Println(common)
			}
			seen[truncated] = true
		}
	}

}

func hasNumberLetters(boxid string, letterOccurence int) bool {
	letters := strings.Split(boxid, "")
	dictionary := make(map[string]int)

	for _, letter := range letters {
		if _, ok := dictionary[letter]; ok {
			dictionary[letter]++
		} else {
			dictionary[letter] = 1
		}
	}

	for _, count := range dictionary {
		if count == letterOccurence {
			return true
		}
	}

	return false
}
