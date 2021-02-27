package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var cache map[int]int

func main() {
	data, err := os.ReadFIle("jolt.data")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	records := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Records in Dataset:", len(records))
	fmt.Println()
	fmt.Println("Day 09: PART 1 - Adapter Array")

	adaptorArray := make([]int, len(records))
	for i, record := range records {
		adaptorArray[i], _ = strconv.Atoi(record)
	}
	fmt.Println(adaptorArray)
	fmt.Println()
	fmt.Println("Sorted:")
	sort.Ints(adaptorArray)
	fmt.Println(adaptorArray)

	joltRating := 0

	joltDiffOne := 0
	joltDiffTwo := 0
	joltDiffThree := 0
	previousAdaptorJoltage := 0

	for _, adaptor := range adaptorArray {
		switch adaptor - previousAdaptorJoltage {
		case 1:
			joltDiffOne++
		case 2:
			joltDiffTwo++
		case 3:
			joltDiffThree++
		default:
			fmt.Println("Joltdiff wrong", adaptor-previousAdaptorJoltage)
			break
		}

		previousAdaptorJoltage = adaptor
		joltRating = adaptor
	}
	// The final adaptor can also accomondate a three jolt
	joltDiffThree++

	fmt.Println()
	fmt.Println("Jolt diff One:", joltDiffOne)
	fmt.Println("Jolt diff Two:", joltDiffTwo)
	fmt.Println("Jolt diff Three:", joltDiffThree)
	fmt.Println("Max Jolt rating:", joltRating)
	fmt.Println("Calculated rating:", (joltDiffOne * joltDiffThree))

	fmt.Println()
	fmt.Println("Day 10: PART 2 - Combinations")
	// reset counters
	joltDiffOne = 0
	joltDiffThree = 0
	combiCount := 1
	cache = make(map[int]int, len(adaptorArray))

	currentGroup := []int{0}
	for _, adaptor := range adaptorArray {
		diff := adaptor - previousAdaptorJoltage
		if diff == 1 {
			joltDiffOne++
		} else if diff == 3 {
			joltDiffOne++
			// for each group we count the combinations inside it
			combiCount *= countGroupCombinations(currentGroup)
			currentGroup = nil
		}
		currentGroup = append(currentGroup, adaptor)
		previousAdaptorJoltage = adaptor
	}

	fmt.Println("Combination Count:", combiCount)

}

func countGroupCombinations(group []int) int {
	if len(group) == 1 {
		return 1
	}

	target := group[len(group)-1]
	if v, ok := cache[target]; ok {
		return v
	}

	var combinations int
	for i := len(group) - 2; i >= 0 && target-group[i] <= 3; i-- {
		combinations += countGroupCombinations(group[:i+1])
	}
	cache[target] = combinations
	return combinations

}
