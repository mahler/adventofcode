package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	puzzleInput := string(content)

	// Solve and print results
	segments := strings.Split(puzzleInput, "\n\n")
	seedRe := regexp.MustCompile(`\d+`)
	seeds := seedRe.FindAllString(segments[0], -1)

	minLocation := math.MaxInt64
	for _, seedStr := range seeds {
		x, _ := strconv.Atoi(seedStr)

		for _, seg := range segments[1:] {
			conversionRe := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
			conversions := conversionRe.FindAllStringSubmatch(seg, -1)

			converted := false
			for _, conversion := range conversions {
				destination, _ := strconv.Atoi(conversion[1])
				start, _ := strconv.Atoi(conversion[2])
				delta, _ := strconv.Atoi(conversion[3])

				if x >= start && x < start+delta {
					x += destination - start
					converted = true
					break
				}
			}

			if !converted {
				// If no conversion found, x remains the same
				continue
			}
		}

		if x < minLocation {
			minLocation = x
		}
	}

	fmt.Println("Part 1: What is the lowest location number that corresponds to any of the initial seed numbers?")
	fmt.Println(minLocation)

	//segments := strings.Split(puzzleInput, "\n\n")
	p2seedRe := regexp.MustCompile(`(\d+) (\d+)`)
	seedMatches := p2seedRe.FindAllStringSubmatch(segments[0], -1)

	var intervals [][]int
	for _, match := range seedMatches {
		x1, _ := strconv.Atoi(match[1])
		dx, _ := strconv.Atoi(match[2])
		x2 := x1 + dx
		intervals = append(intervals, []int{x1, x2, 1})
	}

	p2minLocation := math.MaxInt64

	for len(intervals) > 0 {
		// Pop the last interval
		x1, x2, level := intervals[len(intervals)-1][0], intervals[len(intervals)-1][1], intervals[len(intervals)-1][2]
		intervals = intervals[:len(intervals)-1]

		if level == 8 {
			if x1 < p2minLocation {
				p2minLocation = x1
			}
			continue
		}

		conversionRe := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
		conversions := conversionRe.FindAllStringSubmatch(segments[level], -1)

		converted := false
		for _, conversion := range conversions {
			z, _ := strconv.Atoi(conversion[1])
			y1, _ := strconv.Atoi(conversion[2])
			dy, _ := strconv.Atoi(conversion[3])
			y2 := y1 + dy
			diff := z - y1

			// No overlap
			if x2 <= y1 || y2 <= x1 {
				continue
			}

			// Split intervals if necessary
			if x1 < y1 {
				intervals = append(intervals, []int{x1, y1, level})
				x1 = y1
			}

			if y2 < x2 {
				intervals = append(intervals, []int{y2, x2, level})
				x2 = y2
			}

			// Add converted interval to next level
			intervals = append(intervals, []int{x1 + diff, x2 + diff, level + 1})
			converted = true
			break
		}

		// If no conversion found, pass to next level
		if !converted {
			intervals = append(intervals, []int{x1, x2, level + 1})
		}
	}

	fmt.Println()
	fmt.Println("Part 2: What is the lowest location number that corresponds to any of the initial seed numbers?")
	fmt.Println(p2minLocation)
}
