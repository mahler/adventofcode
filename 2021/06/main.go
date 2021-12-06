package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	// Setup
	source := string(fileContent)
	strLaternfish := strings.Split(source, ",")

	lanternfish := []int{}
	for i := 0; i < len(strLaternfish); i++ {
		iFish, _ := strconv.Atoi(strLaternfish[i])
		lanternfish = append(lanternfish, iFish)
	}

	//fmt.Print(laternfish)

	rounds := 80

	for round := 0; round < rounds; round++ {
		fmt.Println("Round", round+1)
		newFish := []int{}

		for numberFish := 0; numberFish < len(lanternfish); numberFish++ {
			if lanternfish[numberFish] > 0 {
				lanternfish[numberFish]--
			} else {
				// new fish
				newFish = append(newFish, 8)

				// reset timer for current fish
				lanternfish[numberFish] = 6
			}
		}
		lanternfish = append(lanternfish, newFish...)
	}

	fmt.Println("Rounds run:", rounds)
	fmt.Println("Number of fish:", len(lanternfish))
}
