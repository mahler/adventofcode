package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Range represents a firewall layer with its depth and range
type Range struct {
	depth int
	rng   int
}

// ParseRange parses a single line into a Range struct
func ParseRange(line string) Range {
	re := regexp.MustCompile(`(\d+): (\d+)`)
	matches := re.FindStringSubmatch(line)
	depth, _ := strconv.Atoi(matches[1])
	rng, _ := strconv.Atoi(matches[2])
	return Range{depth: depth, rng: rng}
}

// RangePosition calculates scanner position at a given time (not needed in final solution)
func RangePosition(rng, time int) int {
	range1 := rng - 1
	q := time / range1
	r := time % range1
	if q%2 == 0 {
		return r
	}
	return range1 - r
}

// RangeCaught determines if a packet would be caught by a scanner
func RangeCaught(rng, time int) bool {
	return time%(2*(rng-1)) == 0
}

// RangesCaught returns the set of depths where packet is caught
func RangesCaught(ranges map[int]int, delay int) map[int]bool {
	caught := make(map[int]bool)
	for depth, rng := range ranges {
		if RangeCaught(rng, delay+depth) {
			caught[depth] = true
		}
	}
	return caught
}

func main() {
	// Read input file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(content)
	ranges := make(map[int]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		r := ParseRange(scanner.Text())
		ranges[r.depth] = r.rng
	}

	// Part 1
	caught := RangesCaught(ranges, 0)
	severity := 0
	for depth := range caught {
		severity += depth * ranges[depth]
	}

	fmt.Println("Part 1: Given the details of the firewall you've recorded,")
	fmt.Println("if you leave immediately, what is the severity of your whole trip?")
	fmt.Println(severity)

	// Part 2
	delay := 0
	for {
		if len(RangesCaught(ranges, delay)) == 0 {
			break
		}
		delay++
	}
	fmt.Println()
	fmt.Println("Part 2: What is the fewest number of picoseconds that you need to delay")
	fmt.Println("the packet to pass through the firewall without being caught?")
	fmt.Println(delay)
}
