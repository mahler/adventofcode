package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	numIncreases := 0
	numDecreases := 0
	number := 0
	tmpNumber := 0

	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println()

	for _, fileRow := range fileLines {
		//	fmt.Println(fileRow)
		if number == 0 {
			number, _ = strconv.Atoi(fileRow)
		} else {
			tmpNumber, _ = strconv.Atoi(fileRow)
			if tmpNumber < number {
				numDecreases++
			} else if tmpNumber > number {
				numIncreases++
			} else {
				fmt.Println("Equal")
			}
			number = tmpNumber
		}

	}

	fmt.Println()
	fmt.Println("Day 1: Sonar Sweep")
	fmt.Println("Increases:", numIncreases)
	// fmt.Println("Decreases:", numDecreases)

}
