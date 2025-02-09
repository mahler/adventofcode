package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	mem := map[int]int{}

	for i, s := range split {
		mem[i], _ = strconv.Atoi(s)
	}

	in, out := make(chan int, 1), make(chan int)
	go run(mem, in, out)
	count := 0

	for range out {
		<-out
		if <-out == 2 {
			count++
		}
	}

	fmt.Println("Part 1: How many block tiles are on the screen when the game exits?")
	fmt.Println(count)

	// part 2

	// Reset mem...
	mem = map[int]int{}
	for i, s := range split {
		mem[i], _ = strconv.Atoi(s)
	}

	mem[0] = 2
	in, out = make(chan int, 1), make(chan int)
	go run(mem, in, out)
	paddle := 0
	score := 0

	for x := range out {
		y := <-out
		id := <-out

		if x == -1 && y == 0 {
			score = id
		} else if id == 3 {
			paddle = x
		} else if id == 4 {
			if paddle < x {
				in <- 1
			} else if paddle > x {
				in <- -1
			} else {
				in <- 0
			}
		}
	}
	fmt.Println()
	fmt.Println("Part 2: What is your score after the last block is broken?")
	fmt.Println(score)

}

func run(mem map[int]int, in <-chan int, out chan<- int) {
	ip, rb := 0, 0

	for {
		ins := fmt.Sprintf("%05d", mem[ip])
		op, _ := strconv.Atoi(ins[3:])
		par := func(i int) int {
			switch ins[3-i] {
			case '1':
				return ip + i
			case '2':
				return rb + mem[ip+i]
			default:
				return mem[ip+i]
			}
		}

		switch op {
		case 1:
			mem[par(3)] = mem[par(1)] + mem[par(2)]
		case 2:
			mem[par(3)] = mem[par(1)] * mem[par(2)]
		case 3:
			mem[par(1)] = <-in
		case 4:
			out <- mem[par(1)]
		case 5:
			if mem[par(1)] != 0 {
				ip = mem[par(2)]
				continue
			}
		case 6:
			if mem[par(1)] == 0 {
				ip = mem[par(2)]
				continue
			}
		case 7:
			if mem[par(1)] < mem[par(2)] {
				mem[par(3)] = 1
			} else {
				mem[par(3)] = 0
			}
		case 8:
			if mem[par(1)] == mem[par(2)] {
				mem[par(3)] = 1
			} else {
				mem[par(3)] = 0
			}
		case 9:
			rb += mem[par(1)]
		case 99:
			close(out)
			return
		}

		ip += []int{1, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
	}
}
