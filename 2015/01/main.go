package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	dataStr := string(data)

	fmt.Println()
	fmt.Println("2019 - Day 01 Part 1")
	floor := 0
	for pos := 0; pos < len(dataStr); pos++ {
		dir := dataStr[pos : pos+1]
		switch dir {
		case "(":
			floor++
		case ")":
			floor--

		}
	}
	fmt.Println("To what floor do the instructions take Santa?", floor)
	fmt.Println()
	fmt.Println("2019 - Day 01 Part 2")
	// reset floor for part 2
	floor = 0
	firstTimeBasement := 0
	for pos := 0; pos < len(dataStr); pos++ {
		dir := dataStr[pos : pos+1]
		switch dir {
		case "(":
			floor++
		case ")":
			floor--

		}
		fmt.Println(pos, " floor", floor)
		if floor == -1 {
			firstTimeBasement = pos
			break
		}
	}
	// String has 0-base, so add one to get the correct character number.
	firstTimeBasement++
	fmt.Println("What is the position of the character that causes Santa to first enter the basement?", firstTimeBasement)

}
