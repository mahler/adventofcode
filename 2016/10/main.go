package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bots := make(map[int][]int)
	output := make(map[int][]int)
	var specialbot int

	file, _ := os.Open("puzzle.txt")
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	// Initial distribution of chips
	for _, line := range data {
		if strings.HasPrefix(line, "value") {
			parts := strings.Fields(line)
			value, _ := strconv.Atoi(parts[1])
			botNum, _ := strconv.Atoi(parts[5])
			bots[botNum] = append(bots[botNum], value)
		}
	}

	changed := true
	for changed {
		changed = false
		for _, line := range data {
			if strings.HasPrefix(line, "bot") {
				parts := strings.Fields(line)
				botNum, _ := strconv.Atoi(parts[1])

				if len(bots[botNum]) >= 2 {
					changed = true
					min, max := minMax(bots[botNum])

					// Low value distribution
					if parts[5] == "bot" {
						lowNum, _ := strconv.Atoi(parts[6])
						bots[lowNum] = append(bots[lowNum], min)
					} else {
						lowNum, _ := strconv.Atoi(parts[6])
						output[lowNum] = append(output[lowNum], min)
					}

					// High value distribution
					if parts[10] == "bot" {
						highNum, _ := strconv.Atoi(parts[11])
						bots[highNum] = append(bots[highNum], max)
					} else {
						highNum, _ := strconv.Atoi(parts[11])
						output[highNum] = append(output[highNum], max)
					}

					// Clear current bot
					bots[botNum] = []int{}
				}

				// Check for special bot
				for bot, chips := range bots {
					if contains(chips, 17) && contains(chips, 61) {
						specialbot = bot
					}
				}
			}
		}
	}

	fmt.Println("Part 1: what is the number of the bot that is responsible for comparing value-61 microchips with value-17 microchips?")
	fmt.Println(specialbot)

	fmt.Println()
	fmt.Println("Part 2: What do you get if you multiply together the values of one chip in each of outputs 0, 1, and 2?")
	fmt.Println(output[0][0] * output[1][0] * output[2][0])
}

func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
