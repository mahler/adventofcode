package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Disc struct {
	position int
	mod      int
	start    int
}

// extracts all numbers from a string using regex
func extractNumbers(s string) []int {
	re := regexp.MustCompile(`[0-9]+`)
	matches := re.FindAllString(s, -1)
	numbers := make([]int, 0, len(matches))

	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

// reads the puzzle input file and returns slice of Disc structs
func readDiscs(filename string) ([]Disc, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var discs []Disc
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := extractNumbers(line)
		if len(numbers) >= 4 {
			disc := Disc{
				position: numbers[0],
				mod:      numbers[1],
				start:    numbers[3],
			}
			discs = append(discs, disc)
		}
	}
	return discs, scanner.Err()
}

func readDiscs2(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var discs [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := extractNumbers(line)
		if len(numbers) > 0 {
			discs = append(discs, numbers)
		}
	}

	// Part 2: Add the additional disc
	discs = append(discs, []int{len(discs) + 1, 11, 0, 0})

	return discs, scanner.Err()
}

// finds the first time when all discs align
func findAlignment(discs []Disc) int {
	time := 0
	for {
		aligned := true
		for _, disc := range discs {
			if (time+disc.position+disc.start)%disc.mod != 0 {
				aligned = false
				break
			}
		}
		if aligned {
			return time
		}
		time++
	}
}

func findFirstTime(discs [][]int) int {
	time := 0
	for {
		valid := true
		for _, disc := range discs {
			position, mod, _, start := disc[0], disc[1], disc[2], disc[3]
			if (time+position+start)%mod != 0 {
				valid = false
				break
			}
		}
		if valid {
			return time
		}
		time++
	}
}

func main() {
	discs, err := readDiscs("puzzle.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	result := findAlignment(discs)
	fmt.Println("art 1: What is the first time you can press the button to get a capsule?")
	fmt.Println(result)

	discs2, _ := readDiscs2("puzzle.txt")
	result2 := findFirstTime(discs2)
	fmt.Println()
	fmt.Println("Part 2: With this new disc, and counting again starting from time=0")
	fmt.Println("with the configuration in your puzzle input, what is the first time you")
	fmt.Println("can press the button to get another capsule?")
	fmt.Println(result2)
}
