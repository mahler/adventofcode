package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Shift struct {
	asleep [60]bool
}

type Guard struct {
	shifts []*Shift
}

func main() {
	fileContent, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}
	fileData := strings.Split(string(fileContent), "\n")
	sort.Strings(fileData)

	shiftRx := regexp.MustCompile(`\[\d+-\d+-\d+ \d+:\d+\] Guard #(\d+) begins shift`)
	asleepRx := regexp.MustCompile(`\[\d+-\d+-\d+ 00:(\d+)\] falls asleep`)
	awakeRx := regexp.MustCompile(`\[\d+-\d+-\d+ 00:(\d+)\] wakes up`)

	// Setup data structure for Guard plans
	guards := make(map[int]*Guard)
	var lastShift *Shift
	for _, line := range fileData {
		if result := shiftRx.FindStringSubmatch(line); result != nil {
			id, _ := strconv.Atoi(result[1])
			guard := guards[id]
			if guard == nil {
				guard = new(Guard)
				guards[id] = guard
			}
			lastShift = new(Shift)
			guard.shifts = append(guard.shifts, lastShift)
		} else if result := asleepRx.FindStringSubmatch(line); result != nil {
			startI, _ := strconv.Atoi(result[1])
			for i := startI; i < 60; i++ {
				lastShift.asleep[i] = true
			}
		} else if result := awakeRx.FindStringSubmatch(line); result != nil {
			startI, _ := strconv.Atoi(result[1])
			for i := startI; i < 60; i++ {
				lastShift.asleep[i] = false
			}
		}
	}
	// --------------------- PART 1
	fmt.Println("--- Part One ---")
	bestMinutes := 0
	bestResult := 0
	for id, guard := range guards {
		minutes := 0
		for _, shift := range guard.shifts {
			for i := 0; i < 60; i++ {
				if shift.asleep[i] {
					minutes++
				}
			}
		}

		bestCount := 0
		bestTarget := 0
		for i := 0; i < 60; i++ {
			count := 0
			for _, shift := range guard.shifts {
				if shift.asleep[i] {
					count++
				}
			}

			if count > bestCount {
				bestCount = count
				bestTarget = i
			}
		}

		if minutes > bestMinutes {
			bestMinutes = minutes
			bestResult = id * bestTarget
		}
	}
	fmt.Println("Strategy 1: What is the ID of the guard you chose multiplied by the minute you chose?", bestResult)

	// -------- Part 2
	fmt.Println()
	fmt.Println("Part 2")
	// Reset counters for part 2
	bestCount := 0
	bestResult = 0
	for id, guard := range guards {
		for i := 0; i < 60; i++ {
			count := 0
			for _, shift := range guard.shifts {
				if shift.asleep[i] {
					count++
				}
			}

			if count > bestCount {
				bestCount = count
				bestResult = id * i
			}
		}
	}
	fmt.Println("Strategy 2: What is the ID of the guard you chose multiplied by the minute you chose?", bestResult)
}
