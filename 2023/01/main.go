package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	total := 0
	re := regexp.MustCompile(`\D`)

	for _, line := range lines {
		// Remove non-digit characters
		digitLine := re.ReplaceAllString(line, "")

		// Get first and last digit
		if len(digitLine) > 0 {
			firstLast := digitLine[0:1] + digitLine[len(digitLine)-1:]
			num, _ := strconv.Atoi(firstLast)
			total += num
		}
	}

	fmt.Println("Part 1: What is the sum of all of the calibration values?")
	fmt.Println(total)

	// Part 2
	values := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var pairs []int
	for _, line := range lines {
		var digits []string

		// Iterate through each character in the line
		for i := 0; i < len(line); i++ {
			// Check if current character is a digit
			if line[i] >= '0' && line[i] <= '9' {
				digits = append(digits, string(line[i]))
			} else {
				// Check for word representations of digits
				for k, v := range values {
					if strings.HasPrefix(line[i:], k) {
						digits = append(digits, v)
						break
					}
				}
			}
		}

		// Combine first and last digit
		if len(digits) > 0 {
			numStr := digits[0] + digits[len(digits)-1]
			num, _ := strconv.Atoi(numStr)
			pairs = append(pairs, num)
		}
	}

	// Sum all pairs
	total = 0
	for _, p := range pairs {
		total += p
	}

	fmt.Println()
	fmt.Println("Part 2: What is the sum of all of the calibration values?")
	fmt.Println(total)
}
