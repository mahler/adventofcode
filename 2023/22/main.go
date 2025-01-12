package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Brick struct {
	x, y, z [2]int
	below   map[*Brick]bool
	above   map[*Brick]bool
}

func NewBrick(line string) *Brick {
	parts := strings.Split(strings.TrimSpace(line), "~")
	a := parseCoords(parts[0])
	b := parseCoords(parts[1])

	return &Brick{
		x:     [2]int{min(a[0], b[0]), max(a[0], b[0])},
		y:     [2]int{min(a[1], b[1]), max(a[1], b[1])},
		z:     [2]int{min(a[2], b[2]), max(a[2], b[2])},
		below: make(map[*Brick]bool),
		above: make(map[*Brick]bool),
	}
}

func parseCoords(s string) [3]int {
	parts := strings.Split(s, ",")
	result := [3]int{}
	for i, p := range parts {
		val, _ := strconv.Atoi(p)
		result[i] = val
	}
	return result
}

func (b *Brick) isBelow(other *Brick) bool {
	return b.x[0] <= other.x[1] && other.x[0] <= b.x[1] &&
		b.y[0] <= other.y[1] && other.y[0] <= b.y[1] &&
		b.z[1] <= other.z[0]
}

func (b *Brick) drop(newz int) {
	length := b.z[1] - b.z[0]
	b.z = [2]int{newz, newz + length}
}

func (b *Brick) collapse() map[*Brick]bool {
	removed := make(map[*Brick]bool)
	removed[b] = true
	for other := range b.above {
		other.collapseHelper(removed)
	}
	delete(removed, b)
	return removed
}

func (b *Brick) collapseHelper(removed map[*Brick]bool) {
	// Check if all supporting bricks are in removed set
	allBelowRemoved := true
	for brick := range b.below {
		if !removed[brick] {
			allBelowRemoved = false
			break
		}
	}

	if allBelowRemoved {
		removed[b] = true
		for other := range b.above {
			other.collapseHelper(removed)
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var bricks []*Brick
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bricks = append(bricks, NewBrick(scanner.Text()))
	}

	// Sort bricks by z-value
	slices.SortFunc(bricks, func(a, b *Brick) int {
		return a.z[0] - b.z[0]
	})

	// Create slices to track bricks by z-value
	maxZ := bricks[len(bricks)-1].z[0]
	byZval := make([][]*Brick, maxZ)
	for i := range byZval {
		byZval[i] = make([]*Brick, 0)
	}

	// Add base floor
	base := NewBrick("0,0,0~1000,1000,0")
	settled := []*Brick{base}

	// Settle bricks
	for _, brick := range bricks {
		// Find highest z-value below current brick
		topZ := 0
		for _, other := range settled {
			if other.isBelow(brick) && other.z[1] > topZ {
				topZ = other.z[1]
			}
		}

		brick.drop(topZ + 1)
		settled = append(settled, brick)
		byZval[topZ+1] = append(byZval[topZ+1], brick)
	}

	// Find supporting relationships
	for _, brick := range bricks {
		for _, other := range byZval[brick.z[1]+1] {
			if brick.isBelow(other) {
				other.below[brick] = true
				brick.above[other] = true
			}
		}
	}

	// Count removable bricks and chain reactions
	removeCount := 0
	chainCount := 0

	for _, brick := range bricks {
		if len(brick.above) == 0 {
			removeCount++
			continue
		}

		canRemove := true
		for other := range brick.above {
			if len(other.below) <= 1 {
				canRemove = false
				break
			}
		}

		if canRemove {
			removeCount++
		} else {
			chainCount += len(brick.collapse())
		}
	}

	fmt.Println("Part 1: How many bricks could be safely chosen as the one to get disintegrated?")
	fmt.Println(removeCount)

	fmt.Println()
	fmt.Println("Part 2: What is the sum of the number of other bricks that would fall?")
	fmt.Println(chainCount)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
