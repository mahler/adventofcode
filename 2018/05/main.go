package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func main() {
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	// Setup
	source := string(fileContent)
	polarities := []string{}
	for r := 'a'; r < 'z'; r++ {
		R := unicode.ToUpper(r)
		stringR := fmt.Sprintf("%c%c", r, R)
		polarities = append(polarities, stringR)
		// We also need the reverse combo - aA and Aa for all letters
		stringR = fmt.Sprintf("%c%c", R, r)
		polarities = append(polarities, stringR)
	}

	// Part 1 reduction
	fmt.Println()
	fmt.Println("2018 - Day 05, Part 1")
	fmt.Println("Alchemical Reduction")

	//source = "dabAcCaCBAcCcaDA"

	for {
		preLenght := len(source)
		source = reduce(source)
		if preLenght == len(source) {
			// No more reduction possible.
			break
		}
	}
	fmt.Println()
	fmt.Println("How many units remain after fully reacting the polymer you scanned? ")
	fmt.Println(len(source))
}

func reduce(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] && strings.EqualFold(string(s[i]), string(s[i+1])) {
			return s[0:i] + s[i+2:]
		}
	}
	return s
}
