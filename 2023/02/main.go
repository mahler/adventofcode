package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1 := 0
	part2 := 0

	for i, l := range strings.Split(string(input), "\n") {
		if l == "" {
			continue
		}

		l = l[8:]
		maxRed, maxGreen, maxBlue := 0, 0, 0
		gamePossible := true

		for _, s := range strings.Split(l, ";") {
			for _, h := range strings.Split(s, ",") {
				h = strings.TrimSpace(h)
				parts := strings.Split(h, " ")
				amount, _ := strconv.Atoi(parts[0])
				color := parts[1]

				if color == "red" {
					if amount > 12 {
						gamePossible = false
					}
					if amount > maxRed {
						maxRed = amount
					}
				}
				if color == "green" {
					if amount > 13 {
						gamePossible = false
					}
					if amount > maxGreen {
						maxGreen = amount
					}
				}
				if color == "blue" {
					if amount > 14 {
						gamePossible = false
					}
					if amount > maxBlue {
						maxBlue = amount
					}
				}
			}
		}

		if gamePossible {
			part1 += i + 1
		}

		part2 += maxRed * maxGreen * maxBlue
	}
	fmt.Println("Day 2: Cube Conundrum")
	fmt.Println("Part 1: What is the sum of the IDs of those games?")
	fmt.Printf("%d\n", part1)
	fmt.Println()
	fmt.Println("Part 2: What is the sum of the power of these sets?")
	fmt.Printf("%d\n", part2)
}
