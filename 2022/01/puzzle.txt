package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elves := make([][]int, 0)
	elf := make([]int, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			elves = append(elves, elf)
			elf = make([]int, 0)
		} else {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			elf = append(elf, i)
		}
	}
	// most := mostCalories(elves)
	fmt.Println(topThree(elves))
}

func mostCalories(elves [][]int) int {
	highest := 0
	for _, elf := range elves {
		sum := 0
		for _, cal := range elf {
			sum += cal
		}
		if sum > highest {
			highest = sum
		}
	}
	return highest
}

func topThree(elves [][]int) int {
	tops := [3]int{0, 0, 0}
	for _, elf := range elves {
		sum := 0
		for _, cal := range elf {
			sum += cal
		}
		x := sum
		for i, n := range tops {
			if x > tops[i] {
				tops[i] = x
				x = n
				continue
			}
		}
	}
	sum := 0
	for _, n := range tops {
		sum += n
	}
	return sum
}