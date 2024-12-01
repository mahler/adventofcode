package main

import "fmt"

const N_ELVES = 3005290

type Elf struct {
	n     int
	left  *Elf
	right *Elf
}

var elfCounter int

func NewElf(right *Elf) *Elf {
	elfCounter++
	return &Elf{
		n:     elfCounter,
		right: right,
	}
}

func (e *Elf) removeSelf() *Elf {
	e.left.right = e.right
	e.right.left = e.left
	return e.left
}

func doStuff(partTwo bool) {
	// Reset elf counter
	elfCounter = 0

	// Create all of the elves and link them together
	firstElf := NewElf(nil)
	lastElf := firstElf
	for i := 0; i < N_ELVES-1; i++ {
		newElf := NewElf(lastElf)
		lastElf.left = newElf
		lastElf = newElf
	}
	lastElf.left = firstElf
	firstElf.right = lastElf
	totalElves := N_ELVES

	// Find first elf that is removed from the game
	currentElf := firstElf
	if partTwo {
		for range N_ELVES / 2 {
			currentElf = currentElf.left
		}
	} else {
		currentElf = currentElf.left
	}

	// Take one elf out of the game at a time
	for totalElves > 1 {
		if partTwo {
			currentElf = currentElf.removeSelf()
			if totalElves%2 == 1 {
				currentElf = currentElf.left
			}
		} else {
			currentElf = currentElf.removeSelf().left
		}
		totalElves--
	}

	// Show the remaining elf
	part := 1
	if partTwo {
		part = 2
	}
	fmt.Printf("Part %d: %d\n", part, currentElf.n)
}

func main() {
	doStuff(false)
	doStuff(true)
}
