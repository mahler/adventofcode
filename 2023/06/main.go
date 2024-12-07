package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Helper function to extract numbers from a string using a regex
func extractNumbers(re *regexp.Regexp, s string) []int {
	matches := re.FindAllString(s, -1)
	numbers := make([]int, len(matches))

	for i, match := range matches {
		num, _ := strconv.Atoi(match)
		numbers[i] = num
	}

	return numbers
}

func main() {
	// Read input file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	puzzleInput := string(content)

	// Part 1
	lines := strings.Split(puzzleInput, "\n")

	// Parse times and distances
	timeRe := regexp.MustCompile(`\d+`)
	times := extractNumbers(timeRe, lines[0])
	distances := extractNumbers(timeRe, lines[1])

	total := 1
	for i, t := range times {
		d := distances[i]
		wins := 0

		for speed := 1; speed < t; speed++ {
			travelled := (t - speed) * speed
			if travelled > d {
				wins++
			} else if wins > 0 {
				// Stop if we've already seen wins and now we're not winning
				break
			}
		}

		total *= wins
	}

	fmt.Println("Part 1: What do you get if you multiply these numbers together?")
	fmt.Println(total)

	// Part 2

	// Parse time and distance, removing spaces
	timeStr := strings.Join(timeRe.FindAllString(lines[0], -1), "")
	distanceStr := strings.Join(timeRe.FindAllString(lines[1], -1), "")

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)

	// Quadratic formula to find minimum acceleration point
	exactAcceleration := (float64(time) - math.Sqrt(float64(time*time-4*distance))) / 2
	minAcceleration := int(exactAcceleration + 1)

	fmt.Println()
	fmt.Println("Part 2: How many ways can you beat the record in this one much longer race?")
	fmt.Println(time - 2*minAcceleration + 1)
}
