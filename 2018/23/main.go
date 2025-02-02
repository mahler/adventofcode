package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Nanobot struct {
	Pos [3]int64
	R   int64
}

func manhattanDist(a, b [3]int64) int64 {
	return abs(a[0]-b[0]) + abs(a[1]-b[1]) + abs(a[2]-b[2])
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	bots := []Nanobot{}
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var bot Nanobot
		posStart := strings.Index(line, "<") + 1
		posEnd := strings.Index(line, ">")
		posParts := strings.Split(line[posStart:posEnd], ",")
		for i := 0; i < 3; i++ {
			bot.Pos[i], _ = strconv.ParseInt(strings.TrimSpace(posParts[i]), 10, 64)
		}
		rStart := strings.LastIndex(line, "=") + 1
		bot.R, _ = strconv.ParseInt(strings.TrimSpace(line[rStart:]), 10, 64)
		bots = append(bots, bot)
	}

	// Part 1
	biggestIdx := 0
	for i, bot := range bots {
		if bot.R > bots[biggestIdx].R {
			biggestIdx = i
		}
	}

	part1 := 0
	biggestBot := bots[biggestIdx]
	for _, bot := range bots {
		if manhattanDist(bot.Pos, biggestBot.Pos) <= biggestBot.R {
			part1++
		}
	}
	fmt.Println("Part 1: Find the nanobot with the largest signal radius. How many nanobots are in range of its signals?")
	fmt.Println(part1)

	// Part 2
	outGroup := make([]bool, len(bots))
	outOfRange := make([]int, len(bots))

	for {
		outOfRange = make([]int, len(bots))
		for j := range bots {
			if outGroup[j] {
				continue
			}
			for i := range bots {
				if outGroup[i] {
					continue
				}
				if manhattanDist(bots[i].Pos, bots[j].Pos) > bots[i].R+bots[j].R {
					outOfRange[i]++
				}
			}
		}
		maxOut := 0
		for _, v := range outOfRange {
			if v > maxOut {
				maxOut = v
			}
		}
		if maxOut == 0 {
			break
		}
		for i := range outGroup {
			if outOfRange[i] == maxOut {
				outGroup[i] = true
			}
		}
	}

	dists := make([]int64, len(bots))
	for i, bot := range bots {
		if outGroup[i] {
			continue
		}
		dists[i] = manhattanDist(bot.Pos, [3]int64{}) - bot.R
	}

	maxDist := int64(math.MinInt64)
	for _, d := range dists {
		if d > maxDist {
			maxDist = d
		}
	}
	fmt.Println()
	fmt.Println("Part 2: What is the shortest manhattan distance between any of those points and 0,0,0?")
	fmt.Println(maxDist)
}
