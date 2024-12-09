package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawData, _ := os.ReadFile("input.txt")
	stringsData := strings.TrimSpace(string(rawData))
	lines := strings.Split(stringsData, "\n")

	pos := [2]int{0, 0}
	sum1, sum2, sumDir := 0, 0, 0

	posCurr := [2]int{0, 0}
	sum1Curr, sum2Curr, sumDirCurr := 0, 0, 0

	directions := map[string][2]int{
		"L": {-1, 0},
		"R": {1, 0},
		"U": {0, -1},
		"D": {0, 1},
	}

	for _, l := range lines {
		parts := strings.Fields(l)
		d, n, col := parts[0], parts[1], parts[2]

		direction := directions[d]
		length, _ := strconv.Atoi(n)
		nextPos := [2]int{pos[0] + direction[0]*length, pos[1] + direction[1]*length}
		sum1 += pos[0] * nextPos[1]
		sum2 += pos[1] * nextPos[0]
		sumDir += length
		pos = nextPos

		cleaned := strings.Trim(col, "(#)")
		corrLen, _ := strconv.ParseInt(cleaned[:5], 16, 64)
		directionMap := map[string][2]int{
			"0": {1, 0},
			"1": {0, 1},
			"2": {-1, 0},
			"3": {0, -1},
		}
		directionCurr := directionMap[cleaned[5:]]
		nextPosCurr := [2]int{
			posCurr[0] + directionCurr[0]*int(corrLen),
			posCurr[1] + directionCurr[1]*int(corrLen),
		}
		sum1Curr += posCurr[0] * nextPosCurr[1]
		sum2Curr += posCurr[1] * nextPosCurr[0]
		sumDirCurr += int(corrLen)
		posCurr = nextPosCurr
	}

	area := math.Abs(float64(sum1-sum2)) / 2
	fmt.Println("Part 1: The Elves are concerned the lagoon won't be large enough;")
	fmt.Println("if they follow their dig plan, how many cubic meters of lava could it hold?")
	fmt.Println(int(area + float64(sumDir)/2 + 1))

	areaCurr := math.Abs(float64(sum1Curr-sum2Curr)) / 2
	fmt.Println()
	fmt.Println("Part 2: Convert the hexadecimal color codes into the correct instructions;")
	fmt.Println("if the Elves follow this new dig plan, how many cubic meters of lava could the lagoon hold?")
	fmt.Println(int(areaCurr + float64(sumDirCurr)/2 + 1))
}
