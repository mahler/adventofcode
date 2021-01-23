package main

import "fmt"

func main() {
	puzzle := 36000000

	presents := make([]int, puzzle+1)
	for elf := 1; elf <= puzzle; elf++ {
		for i := elf; i <= puzzle; i += elf {
			presents[i] += elf * 10
		}
	}

	fmt.Println()
	fmt.Println("2019")
	fmt.Println("Day 20, Part 1: Infinite Elves and Infinite Houses")
	houseNumber := 0
	for hNumber := 1; hNumber <= puzzle; hNumber++ {
		if presents[hNumber] >= puzzle {
			houseNumber = hNumber
			break
		}
	}
	fmt.Println("What is the lowest house number of the house to get at least")
	fmt.Println("as many presents as the number in your puzzle input?")
	fmt.Println(houseNumber)

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2")
	p2presents := make([]int, puzzle+1)
	for e := 1; e <= puzzle; e++ {
		delivered := 0
		for i := e; i <= puzzle && delivered < 50; i += e {
			p2presents[i] += e * 11
			delivered++
		}
	}
	for i := 1; i <= puzzle; i++ {
		if p2presents[i] >= puzzle {
			houseNumber = i
			break
		}
	}
	fmt.Println("what is the new *lowest house number* of the house to")
	fmt.Println("get at least as many presents as the number in your puzzle input?")
	fmt.Println(houseNumber)
}
