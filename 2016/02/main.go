package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 2: Bathroom Security")
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileRows := strings.Split(string(fileContent), "\n")

	// Keypad
	keypad := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	// Start with position "5" on the keypad (offset 0)
	xPos := 1
	yPos := 1

	pincode := ""

	for _, instructions := range fileRows {

		for i := 0; i < len(instructions); i++ {
			switch instructions[i : i+1] {
			case "U":
				yPos--
			case "D":
				yPos++
			case "L":
				xPos--
			case "R":
				xPos++
			}

			// Handle border overruns
			if xPos < 0 {
				xPos = 0
			} else if xPos > 2 {
				xPos = 2
			}

			if yPos < 0 {
				yPos = 0
			} else if yPos > 2 {
				yPos = 2
			}
		}
		pincode += keypad[yPos][xPos]
	}
	fmt.Println("What is the bathroom code?")
	fmt.Println(pincode)
	// Part 2 ---------------------------------------------
	fmt.Println()
	fmt.Println("Part 2")

	// Reset reused vars for part 2
	pincode = ""
	xPos = 0
	yPos = 2

	// setup keypad
	p2keypad := [5][5]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}

	for _, instructions := range fileRows {
		for i := 0; i < len(instructions); i++ {
			lastX := xPos
			lastY := yPos
			switch instructions[i : i+1] {
			case "U":
				yPos--
			case "D":
				yPos++
			case "L":
				xPos--
			case "R":
				xPos++
			}

			// Handle border overruns
			if xPos < 0 {
				xPos = 0
			} else if xPos > 4 {
				xPos = 4
			}

			if yPos < 0 {
				yPos = 0
			} else if yPos > 4 {
				yPos = 4
			}

			// If outside keypad keys, step back
			if p2keypad[yPos][xPos] == "" {
				xPos = lastX
				yPos = lastY
			}
		}
		pincode += p2keypad[yPos][xPos]
	}

	fmt.Println("what is the correct bathroom code with new keypad?")
	fmt.Println(pincode)

}
