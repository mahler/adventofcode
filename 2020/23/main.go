package main

// Working of solution by https://github.com/kindermoumoute/adventofcode/blob/master/2020/day23/main.go

import (
	// https://golang.org/pkg/container/ring/
	"container/ring"
	"fmt"
)

func main() {
	input := "871369452"

	fmt.Println()
	fmt.Println("Day 23, Part 1")

	ringOfCups := ring.New(len(input))
	index := make(map[int]*ring.Ring)
	for _, cup := range input {
		ringOfCups.Value = int(cup - '0')
		index[ringOfCups.Value.(int)] = ringOfCups
		ringOfCups = ringOfCups.Next()
	}

	// part 1
	current := ringOfCups

	fmt.Println("Current:", current)
	fmt.Println("R: ", ringOfCups)

	initLen := 9
	for i := 0; i < 100; i++ {
		ringOfCups = moveCups(ringOfCups, index, initLen)
	}

	part1 := ""
	index[1].Next().Do(func(i interface{}) {
		if i == 1 {
			return
		}
		part1 += fmt.Sprintf("%d", i)
	})

	fmt.Println("What are the labels on the cups after cup 1:", part1)

	// part 2

	fmt.Println()
	fmt.Println("Part 2")
	ringOfCups = ring.New(len(input))
	index = make(map[int]*ring.Ring)
	for _, cup := range input {
		ringOfCups.Value = int(cup - '0')
		index[ringOfCups.Value.(int)] = ringOfCups
		ringOfCups = ringOfCups.Next()
	}
	current = ringOfCups

	fmt.Println("current:", current)
	fmt.Println("r: ", ringOfCups)

	additionalRings := ring.New(1000000 - 9)
	for i := 10; i <= 1000000; i++ {
		additionalRings.Value = i
		index[i] = additionalRings
		additionalRings = additionalRings.Next()
	}
	current.Prev().Link(additionalRings)
	for i := 0; i < 10000000; i++ {
		current = moveCups(current, index, 1000000)
	}
	part2 := index[1].Next().Value.(int) * index[1].Next().Next().Value.(int)

	fmt.Println("What do you get if you multiply their labels together:", part2)
}

func moveCups(current *ring.Ring, index map[int]*ring.Ring, ringLen int) *ring.Ring {
	excluded := map[interface{}]struct{}{current.Value: {}}
	tmp := current.Unlink(3)
	tmp.Do(func(i interface{}) {
		excluded[i] = struct{}{}
	})

	nextValue := current.Value.(int)
	for _, exist := excluded[nextValue]; exist; _, exist = excluded[nextValue] {
		if nextValue--; nextValue < 1 {
			nextValue = ringLen
		}
	}
	index[nextValue].Link(tmp)

	return current.Next()
}
