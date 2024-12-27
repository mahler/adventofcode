package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const (
	noop = iota
	addx
	empty
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	xVals := []int{}
	x := 1
	cmdBuffer := empty
	currentCmdIndex := -1
	for true {
		xVals = append(xVals, x)
		if cmdBuffer == addx {
			cmdBuffer = empty
			var add int
			command := fileLines[currentCmdIndex]
			fmt.Sscanf(command, "addx %d", &add)
			x += add
			continue
		}
		if cmdBuffer == empty {
			if currentCmdIndex == len(fileLines)-1 {
				break
			}
			currentCmdIndex++
			command := fileLines[currentCmdIndex]
			switch command[0] {
			case 'n':
				cmdBuffer = noop
			case 'a':
				cmdBuffer = addx
			}
		}

		if cmdBuffer == noop {
			cmdBuffer = empty
		}
	}

	part1(xVals)
	part2(xVals)
}

func part2(xVals []int) {
	var crt [][]string
	for x := 0; x < len(xVals)/40; x++ {
		crt = append(crt, []string{})
		for pixel, spriteCenter := range xVals[x*40 : (x+1)*40] {
			closeness := math.Abs(float64(pixel - spriteCenter))
			if closeness <= 1 {
				crt[x] = append(crt[x], "*")
			} else {
				crt[x] = append(crt[x], " 
				")
			}
		}
	}
	for _, row := range crt {
		fmt.Println(row)
	}
}

func part1(xVals []int) {
	total := 0
	for xIndex := 19; xIndex < len(xVals); xIndex += 40 {
		total += (xIndex + 1) * xVals[xIndex]
		fmt.Println(xVals[xIndex])
	}
	fmt.Println(total)
}
