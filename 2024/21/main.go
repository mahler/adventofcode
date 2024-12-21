package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// Point represents a 2D coordinate
type Point struct {
	x, y int32
}

// Cache for memoization
var memoCache map[string]int

//go:embed input.txt
var input string

// numpad converts a key to its coordinates
func numpad(key rune) Point {
	switch key {
	case '7':
		return Point{0, 0}
	case '8':
		return Point{0, 1}
	case '9':
		return Point{0, 2}
	case '4':
		return Point{1, 0}
	case '5':
		return Point{1, 1}
	case '6':
		return Point{1, 2}
	case '1':
		return Point{2, 0}
	case '2':
		return Point{2, 1}
	case '3':
		return Point{2, 2}
	case '0':
		return Point{3, 1}
	case 'A':
		return Point{3, 2}
	default:
		panic("Invalid key")
	}
}

// converts an arrow key to its coordinates
func arrowpad(key rune) Point {
	switch key {
	case '^':
		return Point{0, 1}
	case 'A':
		return Point{0, 2}
	case '<':
		return Point{1, 0}
	case 'v':
		return Point{1, 1}
	case '>':
		return Point{1, 2}
	default:
		panic("Invalid arrow key")
	}
}

// returns the absolute value of an int32
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// generates a unique key for memoization
func getCacheKey(i, j int32, steps int, hFirst bool) string {
	return fmt.Sprintf("%d,%d,%d,%v", i, j, steps, hFirst)
}

// implements the arrow navigation logic with memoization
func doArrows(i, j int32, steps int, hFirst bool) int {
	key := getCacheKey(i, j, steps, hFirst)
	if val, ok := memoCache[key]; ok {
		return val
	}

	ii, jj := abs(i), abs(j)
	chunk := make([]rune, 0)

	// Build the chunk slice
	for k := int32(0); k < ii; k++ {
		if i > 0 {
			chunk = append(chunk, '^')
		} else {
			chunk = append(chunk, 'v')
		}
	}
	for k := int32(0); k < jj; k++ {
		if j > 0 {
			chunk = append(chunk, '<')
		} else {
			chunk = append(chunk, '>')
		}
	}

	if hFirst {
		// Reverse chunk
		for i, j := 0, len(chunk)-1; i < j; i, j = i+1, j-1 {
			chunk[i], chunk[j] = chunk[j], chunk[i]
		}
	}

	chunk = append(chunk, 'A')

	if steps == 0 {
		memoCache[key] = len(chunk)
		return len(chunk)
	}

	loc := arrowpad('A')
	sum := 0

	for _, c := range chunk {
		n := arrowpad(c)
		p := loc
		loc = n
		d := Point{p.x - n.x, p.y - n.y}

		var result int
		if d.x == 0 || d.y == 0 {
			// straight line, search only once, order is irrelevant
			result = doArrows(d.x, d.y, steps-1, false)
		} else if n == (Point{1, 0}) && p.x == 0 {
			// must search down first
			result = doArrows(d.x, d.y, steps-1, false)
		} else if p == (Point{1, 0}) && n.x == 0 {
			// must search horiz first
			result = doArrows(d.x, d.y, steps-1, true)
		} else {
			// can search in either order
			result = min(
				doArrows(d.x, d.y, steps-1, false),
				doArrows(d.x, d.y, steps-1, true),
			)
		}
		sum += result
	}

	memoCache[key] = sum
	return sum
}

// processes a single sequence
func enterSequence(sequence string, steps int) int {
	loc := numpad('A')
	memoCache = make(map[string]int) // Reset cache

	// Parse first 3 chars as number
	multiplier, _ := strconv.Atoi(sequence[:3])

	sum := 0
	for _, c := range sequence {
		n := numpad(c)
		p := loc
		d := Point{loc.x - n.x, loc.y - n.y}
		loc = n

		var result int
		if p.x == 3 && n.y == 0 {
			// must move up first
			result = doArrows(d.x, d.y, steps, false)
		} else if p.y == 0 && n.x == 3 {
			// must move right first
			result = doArrows(d.x, d.y, steps, true)
		} else {
			// move in either direction
			result = min(
				doArrows(d.x, d.y, steps, true),
				doArrows(d.x, d.y, steps, false),
			)
		}
		sum += result
	}

	return multiplier * sum
}

// returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			sum += enterSequence(line, 2)
		}
	}

	fmt.Println("Part 1: What is the sum of the complexities of the five codes on your list?")
	fmt.Println(sum)

	sum = 0
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			sum += enterSequence(line, 25)
		}
	}

	fmt.Println("Part 2: What is the sum of the complexities of the five codes on your list?")
	fmt.Println(sum)
}
