package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	masterAunt := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}

	rxSue := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

	aunts := make(map[int]map[string]int, 500)
	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	for auntNumber, fileLine := range fileLines {
		thisAunt := map[string]int{}
		fields := rxSue.FindStringSubmatch(fileLine)
		numberOfField2, _ := strconv.Atoi(fields[3])
		thisAunt[fields[2]] = numberOfField2

		numberOfField4, _ := strconv.Atoi(fields[5])
		thisAunt[fields[4]] = numberOfField4

		numberOfField6, _ := strconv.Atoi(fields[7])
		thisAunt[fields[6]] = numberOfField6
		aunts[auntNumber] = thisAunt
	}

	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 14, part 1: Aunt Sue")
	for numberAunt, thisAunt := range aunts {
		check := true
		for k, v := range thisAunt {
			if masterValue, ok := masterAunt[k]; ok {
				if masterValue != v {
					check = false
					break
				}
			} else {
				check = false
				break
			}

		}
		if check {
			fmt.Println("The gift is from:")
			// Add one to numberAunt to offset 0-based slice.
			fmt.Println("Ant Sue", numberAunt+1)
		}
	}
	// --------------------------------
	// Part 2

}
