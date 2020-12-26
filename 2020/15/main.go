package main

// Parts sourced from https://github.com/kindermoumoute/adventofcode/blob/master/2020/day15/main.go

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	puzzleInput := "11,0,1,10,5,19"

	list := make(map[int]int)
	for i, numbers := range strings.Split(puzzleInput, ",") {
		number, _ := strconv.Atoi(numbers)
		list[i] = number
	}

	usageIndices := make(map[int][]int)
	twentyTwentyWord := 0
	recently := 0
	// loops for part 2 needs to be set to 30.000.000
	for i := 0; i < 30000000; i++ {
		v, exist := list[i]
		if !exist {
			lastIndex, exist := usageIndices[recently]
			if exist && len(lastIndex) > 1 {
				v = lastIndex[len(lastIndex)-1] - lastIndex[len(lastIndex)-2]
			} else {
				v = 0
			}
		}

		recently = v
		usageIndices[v] = append(usageIndices[v], i)
		// usageIndices are index 0, thus 2019
		if i == 2019 {
			twentyTwentyWord = v
		}
	}

	fmt.Println()
	fmt.Println("Part 01/")
	fmt.Println("The 2020th number spoken:", twentyTwentyWord)
	fmt.Println()
	fmt.Println("Part 02/")
	fmt.Println("What will be the 30000000th number spoken:", recently)
}
