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

	fishCount := make(map[int]int64)
	fishCount[0] = 0
	fishCount[1] = 0
	fishCount[2] = 0
	fishCount[3] = 0
	fishCount[4] = 0
	fishCount[5] = 0
	fishCount[6] = 0
	fishCount[7] = 0
	fishCount[8] = 0

	for i := 0; i < len(lanternfish); i++ {
		//	fmt.Println("found fish:", lanternfish[i])
		fishCount[lanternfish[i]]++
	}

	fmt.Println("Initial state:", fishCount)

	// Set rounds to 80 for part 1
	// set rounds to 256 for part 2
	rounds := 256

	for round := 0; round < rounds; round++ {
		tmpCount := make(map[int]int64)
		tmpCount[0] = fishCount[1]
		tmpCount[1] = fishCount[2]
		tmpCount[2] = fishCount[3]
		tmpCount[3] = fishCount[4]
		tmpCount[4] = fishCount[5]
		tmpCount[5] = fishCount[6]
		tmpCount[6] = fishCount[7]
		tmpCount[7] = fishCount[8]
		tmpCount[8] = fishCount[0]
		tmpCount[6] += fishCount[0]
		fishCount = tmpCount
		//fmt.Println("After Round", round+1, ":", fishCount)
	}

	// --- Count number of fish
	numberOfFish := int64(0)
	for _, n := range fishCount {
		numberOfFish += n
	}

	fmt.Println("Rounds run:", rounds)
	fmt.Println("Number of fish:", numberOfFish)
}
