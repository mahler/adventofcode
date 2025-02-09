package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run(program []int, move, pos, pc int, field [][]string) {
	numOfOperands := []int{0, 3, 3, 1, 1, 2, 2, 3, 3, 1}
	i, base := pc, 0
	directions := map[int][2]int{1: {-1, 0}, 2: {1, 0}, 3: {0, -1}, 4: {0, 1}}
	forbidden := map[int]int{1: 2, 2: 1, 3: 4, 4: 3}
	tiles := []string{"▒", ".", "W"}

	for program[i] != 99 {
		modes := []int{program[i] / 100 % 10, program[i] / 1000 % 10, program[i] / 10000 % 10}
		instruction := program[i] % 100
		baseTmp := make([]int, numOfOperands[instruction])
		for x := 0; x < numOfOperands[instruction]; x++ {
			if modes[x] == 2 {
				baseTmp[x] = base
			} else {
				baseTmp[x] = 0
			}
		}
		operands := make([]int, numOfOperands[instruction])
		for x := 0; x < numOfOperands[instruction]; x++ {
			if modes[x] == 1 {
				operands[x] = program[i+x+1]
			} else {
				operands[x] = program[baseTmp[x]+program[i+x+1]]
			}
		}
		switch instruction {
		case 1:
			program[baseTmp[2]+program[i+3]] = operands[0] + operands[1]
		case 2:
			program[baseTmp[2]+program[i+3]] = operands[0] * operands[1]
		case 3:
			program[baseTmp[0]+program[i+1]] = move
		case 4:
			field[pos/100+directions[move][0]][pos%100+directions[move][1]] = tiles[operands[0]]
			if operands[0] > 0 {
				for x := 1; x <= 4; x++ {
					if x != forbidden[move] {
						newPos := (pos/100+directions[move][0])*100 + pos%100 + directions[move][1]
						run(append([]int(nil), program...), x, newPos, i+numOfOperands[instruction]+1, field)
					}
				}
			}
			return
		case 5:
			if operands[0] != 0 {
				i = operands[1] - 3
			}
		case 6:
			if operands[0] == 0 {
				i = operands[1] - 3
			}
		case 7:
			if operands[0] < operands[1] {
				program[baseTmp[2]+program[i+3]] = 1
			} else {
				program[baseTmp[2]+program[i+3]] = 0
			}
		case 8:
			if operands[0] == operands[1] {
				program[baseTmp[2]+program[i+3]] = 1
			} else {
				program[baseTmp[2]+program[i+3]] = 0
			}
		case 9:
			base += operands[0]
		}
		i += numOfOperands[instruction] + 1
	}
}

func findPath(x, y, length int, field [][]string) (int, int, int) {
	if field[x][y] == "▒" || field[x][y] == "W" {
		if field[x][y] == "▒" {
			return 0, 0, 0
		}
		return x, y, length
	}
	field[x][y] = "▒"
	maxX, maxY, maxLength := 0, 0, 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range directions {
		tmpX, tmpY, tmpLength := findPath(x+d[0], y+d[1], length+1, field)
		if abs(tmpX)+abs(tmpY) > abs(maxX)+abs(maxY) {
			maxX, maxY, maxLength = tmpX, tmpY, tmpLength
		}
	}
	return maxX, maxY, maxLength
}

func distance(x1, y1, x2, y2, length int, field [][]string) int {
	if field[x1][y1] == "▒" {
		return 0
	}
	if x1 == x2 && y1 == y2 {
		return length
	}
	field[x1][y1] = "▒"
	maxDistance := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range directions {
		tmpDistance := distance(x1+d[0], y1+d[1], x2, y2, length+1, field)
		if tmpDistance > maxDistance {
			maxDistance = tmpDistance
		}
	}
	return maxDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strProgram := strings.Split(strings.TrimSpace(string(data)), ",")
	program := make([]int, len(strProgram)+10000)
	for i, s := range strProgram {
		program[i], _ = strconv.Atoi(s)
	}
	field := make([][]string, 41)
	for i := range field {
		field[i] = make([]string, 41)
		for j := range field[i] {
			field[i][j] = " "
		}
	}
	start := 21*100 + 21
	for x := 1; x <= 4; x++ {
		run(append([]int(nil), program...), x, start, 0, field)
	}
	field[21][21] = "S"
	foundX, foundY, foundLength := findPath(21, 21, 0, deepCopy(field))

	//	fmt.Println(foundX, foundY, foundLength)
	fmt.Println("Part 1: What is the fewest number of movement commands required to move")
	fmt.Println("the repair droid from its starting position to the location of the oxygen system?")
	fmt.Println(foundLength)

	// Part 2
	maxDistance := 0
	for x := range field {
		for y := range field[x] {
			if field[x][y] == "." {
				tmpDistance := distance(foundX, foundY, x, y, 0, deepCopy(field))
				if tmpDistance > maxDistance {
					maxDistance = tmpDistance
				}
			}
		}
	}
	fmt.Println()
	fmt.Println("Part 2: How many minutes will it take to fill with oxygen?")
	fmt.Println(maxDistance)
}

func deepCopy(field [][]string) [][]string {
	copiedField := make([][]string, len(field))
	for i := range field {
		copiedField[i] = make([]string, len(field[i]))
		copy(copiedField[i], field[i])
	}
	return copiedField
}
