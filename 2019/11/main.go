package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func executeRobot(mem map[int]int, startColor int) {
	in, out := make(chan int, 1), make(chan int)
	go runProgram(mem, in, out)

	hull := make(map[image.Point]int)
	pos, dir := image.Point{}, 0

	in <- startColor
	for hull[pos] = range out {
		dir = (dir + 2*<-out + 1) % 4
		pos = pos.Add([]image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}[dir])
		in <- hull[pos]
	}

	if startColor == 1 {
		fmt.Println()
		fmt.Println("Part 2: After starting the robot on a single white panel instead, what registration identifier does it paint on your hull?")
		printHull(hull)
	} else {
		fmt.Println("Part 1: How many panels does it paint at least once?")
		fmt.Println(len(hull))
	}
}

func printHull(hull map[image.Point]int) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			fmt.Print([]string{"  ", "██"}[hull[image.Point{x, y}]])
		}
		fmt.Println()
	}
}

func runProgram(mem map[int]int, in <-chan int, out chan<- int) {
	ip, rb := 0, 0
	getParam := func(i int, mode byte) int {
		switch mode {
		case '1':
			return ip + i
		case '2':
			return rb + mem[ip+i]
		default:
			return mem[ip+i]
		}
	}

	for {
		ins := fmt.Sprintf("%05d", mem[ip])
		op, _ := strconv.Atoi(ins[3:])
		modes := ins[:3]

		switch op {
		case 1:
			mem[getParam(3, modes[0])] = mem[getParam(1, modes[2])] + mem[getParam(2, modes[1])]
		case 2:
			mem[getParam(3, modes[0])] = mem[getParam(1, modes[2])] * mem[getParam(2, modes[1])]
		case 3:
			mem[getParam(1, modes[2])] = <-in
		case 4:
			out <- mem[getParam(1, modes[2])]
		case 5:
			if mem[getParam(1, modes[2])] != 0 {
				ip = mem[getParam(2, modes[1])]
				continue
			}
		case 6:
			if mem[getParam(1, modes[2])] == 0 {
				ip = mem[getParam(2, modes[1])]
				continue
			}
		case 7:
			if mem[getParam(1, modes[2])] < mem[getParam(2, modes[1])] {
				mem[getParam(3, modes[0])] = 1
			} else {
				mem[getParam(3, modes[0])] = 0
			}
		case 8:
			if mem[getParam(1, modes[2])] == mem[getParam(2, modes[1])] {
				mem[getParam(3, modes[0])] = 1
			} else {
				mem[getParam(3, modes[0])] = 0
			}
		case 9:
			rb += mem[getParam(1, modes[2])]
		case 99:
			close(out)
			return
		}

		ip += []int{1, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
	}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), ",")

	mem := make(map[int]int)
	for i, s := range split {
		mem[i], _ = strconv.Atoi(s)
	}

	executeRobot(mem, 0)
	executeRobot(mem, 1)
}
