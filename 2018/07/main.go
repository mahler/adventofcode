package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	numWorkers = 5
	baseTime   = 60
)

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var order []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		order = append(order, scanner.Text())
	}

	// Part 1
	required := make([][26]int, 26)
	for _, line := range order {
		step := int(line[36]) - 'A'
		prereq := int(line[5]) - 'A'
		// Find first empty slot
		for i := 0; i < 26; i++ {
			if required[step][i] == 0 {
				required[step][i] = prereq + 1 // +1 to distinguish from empty slots
				break
			}
		}
	}

	part1 := ""
	for len(part1) < 26 {
	outer:
		for i := 0; i < 26; i++ {
			// Check if this letter has any requirements
			for _, req := range required[i] {
				if req != 0 {
					continue outer
				}
			}
			// Found next letter
			part1 += string('A' + i)
			// Mark as completed
			for j := 0; j < 26; j++ {
				for k := 0; k < 26; k++ {
					if required[j][k] == i+1 {
						required[j][k] = 0
					}
				}
			}
			required[i] = [26]int{9999} // Mark as done
			break
		}
	}
	fmt.Println("Part 1: In what order should the steps in your instructions be completed?")
	fmt.Println(part1)

	// Part 2
	required = make([][26]int, 26)
	for _, line := range order {
		step := int(line[36]) - 'A'
		prereq := int(line[5]) - 'A'
		for i := 0; i < 26; i++ {
			if required[step][i] == 0 {
				required[step][i] = prereq + 1
				break
			}
		}
	}

	workers := make([]int, numWorkers)
	time := make([]int, 26)
	for i := range time {
		time[i] = baseTime + i + 1
	}
	done := make([]bool, 26)
	part2 := 0

	for {
		// Check if all tasks are done
		allDone := true
		for _, d := range done {
			if !d {
				allDone = false
				break
			}
		}
		if allDone {
			break
		}

		// Assign available workers
		for i := range workers {
			if workers[i] == 0 {
			workerLoop:
				for j := 0; j < 26; j++ {
					// Check if this task is available
					for _, req := range required[j] {
						if req != 0 {
							continue workerLoop
						}
					}
					if !done[j] && time[j] > 0 {
						workers[i] = j + 1 // +1 to distinguish from unassigned
						required[j] = [26]int{9999}
						break
					}
				}
			}
		}

		// Process work
		for i := range workers {
			if workers[i] != 0 {
				task := workers[i] - 1
				time[task]--
				if time[task] == 0 {
					done[task] = true
					// Clear requirements
					for j := 0; j < 26; j++ {
						for k := 0; k < 26; k++ {
							if required[j][k] == task+1 {
								required[j][k] = 0
							}
						}
					}
					workers[i] = 0
				}
			}
		}
		part2++
	}
	fmt.Println()
	fmt.Println("Part 2: With 5 workers and the 60+ second step durations described above,")
	fmt.Println(" how long will it take to complete all of the steps?")
	fmt.Println(part2)
}
