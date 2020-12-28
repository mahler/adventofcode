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
		if hasNumberLetters(boxid, 2) {
			twoLetters++
		}
		if hasNumberLetters(boxid, 3) {
			threeLetters++
		}

	}

	fmt.Println("Boxes with two letters:", twoLetters)
	fmt.Println("Boxes with three letters:", threeLetters)
	fmt.Println("Checksum:", twoLetters*threeLetters)
	// that have an ID containing exactly two of any letter
	// and then separately counting those with exactly three of any letter.
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
