package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}

	containers := []int{}

	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	for _, fileLine := range fileLines {
		containerSize, _ := strconv.Atoi(fileLine)
		containers = append(containers, containerSize)

	}
	//fmt.Println(containers)
	totalContainers := len(containers)

	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 17, part 1: No Such Thing as Too Much")
	validCombination := 0
	used := make([]int, totalContainers)
	for i := 0; i < (1 << totalContainers); i++ {
		var current, filled int
		for j := 0; j < totalContainers; j++ {
			if (i & (1 << j)) != 0 {
				current += containers[j]
				filled++
			}
		}

		if current == 150 {
			validCombination++
			used[filled]++
		}
	}
	fmt.Println("how many different combinations of containers can exactly fit all 150 liters of eggnog?")
	fmt.Println(validCombination)

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2/")
	fmt.Println("How many different ways* can you fill that number of containers")
	fmt.Println("and still hold exactly 150 litres?")

	validMin := 0
	for _, v := range used {
		if v > 0 {
			validMin = v
			break
		}
	}

	fmt.Println(validMin)

}
