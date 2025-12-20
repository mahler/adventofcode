package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readAndProcess(processMove func(pos, dist int, dir byte) (int, int)) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pos := 50
	zeros := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Invalid distance value: %v", err)
		}

		if dir != 'L' && dir != 'R' {
			log.Fatal("Invalid direction character. Must be 'L' or 'R'.")
		}

		newPos, addZeros := processMove(pos, dist, dir)
		pos = newPos
		zeros += addZeros
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return zeros
}

func part1() {
	zeros := readAndProcess(func(pos, dist int, dir byte) (int, int) {
		if dir == 'L' {
			pos -= dist
		} else {
			pos += dist
		}
		pos = ((pos % 100) + 100) % 100

		if pos == 0 {
			return pos, 1
		}
		return pos, 0
	})

	fmt.Printf("Number of times position was zero: %d\n", zeros)
}

func part2() {
	zeros := readAndProcess(func(pos, dist int, dir byte) (int, int) {
		newPos := pos
		count := 0

		if dir == 'L' {
			newPos = ((pos-dist)%100 + 100) % 100
			if newPos > pos && pos != 0 {
				count++
			}
		} else {
			newPos = (pos + dist) % 100
			if newPos < pos && newPos != 0 {
				count++
			}
		}

		if newPos == 0 {
			count++
		}

		count += dist / 100

		return newPos, count
	})

	fmt.Printf("Number of times position was zero: %d\n", zeros)
}

func main() {
	fmt.Println("What's the actual password to open the door?")
	part1()
	fmt.Println()
	fmt.Println("what is the password to open the door?")
	part2()
}
