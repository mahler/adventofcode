package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CircularList implements a double-ended queue with circular behavior
type CircularList struct {
	items [][2]int
}

func NewCircularList(items [][2]int) *CircularList {
	return &CircularList{items: items}
}

func (c *CircularList) Len() int {
	return len(c.items)
}

func (c *CircularList) PopLeft() [2]int {
	if len(c.items) == 0 {
		panic("empty list")
	}
	item := c.items[0]
	c.items = c.items[1:]
	return item
}

func (c *CircularList) Append(item [2]int) {
	c.items = append(c.items, item)
}

func solve(part int, lines []string) int {
	// Parse input
	X := make([][2]int, len(lines))
	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		if part == 2 {
			num *= 811589153
		}
		X[i] = [2]int{i, num}
	}

	// Create circular list
	clist := NewCircularList(X)
	iterations := 1
	if part == 2 {
		iterations = 10
	}

	// Process mixing
	for t := 0; t < iterations; t++ {
		for i := 0; i < len(X); i++ {
			// Find position of current number
			pos := 0
			for j := 0; j < clist.Len(); j++ {
				if clist.items[j][0] == i {
					pos = j
					break
				}
			}

			// Rotate until target is at front
			temp := make([][2]int, pos)
			copy(temp, clist.items[:pos])
			clist.items = append(clist.items[pos:], temp...)

			// Pop the value and calculate new position
			val := clist.PopLeft()
			toPop := val[1]
			if toPop < 0 {
				toPop = (-toPop) % clist.Len()
				toPop = clist.Len() - toPop
			} else {
				toPop = toPop % clist.Len()
			}

			// Insert at new position
			for k := 0; k < toPop; k++ {
				clist.items = append(clist.items[1:], clist.items[0])
			}
			clist.Append(val)
		}
	}

	// Find zero and calculate result
	zeroPos := 0
	for j := 0; j < clist.Len(); j++ {
		if clist.items[j][1] == 0 {
			zeroPos = j
			break
		}
	}

	result := 0
	l := clist.Len()
	result += clist.items[(zeroPos+1000)%l][1]
	result += clist.items[(zeroPos+2000)%l][1]
	result += clist.items[(zeroPos+3000)%l][1]
	return result
}

func main() {
	// Read input file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	fmt.Println(solve(1, lines))
	fmt.Println(solve(2, lines))
}
