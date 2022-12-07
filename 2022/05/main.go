package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Containers []string
}

func (s *Stack) Move(amount int, to *Stack, asUnit bool) {
	containersLength := len(s.Containers)
	containersToMove := s.Containers[containersLength-amount:]
	s.Containers = s.Containers[0 : containersLength-amount]
	if asUnit {
		to.Containers = append(to.Containers, containersToMove...)
	} else {

		for i := len(containersToMove) - 1; i >= 0; i-- {
			to.Containers = append(to.Containers, containersToMove[i])
		}
	}
}

func main() {
	fmt.Println()
	fmt.Println("Day 5: Supply Stacks")
	fmt.Println("After the rearrangement procedure completes, what crate ends up on top of each stack?")
	fmt.Println(helper("puzzle.txt", true))

	fmt.Println()
	fmt.Println("Part2:")
	fmt.Println("After the rearrangement procedure completes, what crate ends up on top of each stack?")
	fmt.Println(helper("puzzle.txt", false))
}

func helper(filename string, isFirstProb bool) string {
	stacks := map[int]*Stack{}

	file, _ := os.Open("puzzle.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	answer := ""
	maxKey := 1
	for lineIndex, val := range lines {
		// Checking if this is the line that containers the stack numbers
		_, err := strconv.Atoi(string(val[1]))

		// Not a stack if there's an error
		if err != nil {
			continue
		}

		// this only works for stacks less than 10 - bad coding really
		for i, v := range val {
			num, err := strconv.Atoi(string(v))

			if err != nil {
				continue
			}

			// set the max key - looping through map gives inconsistant orders.
			// This way we know stacks start at 1 and end at X
			maxKey = num

			s := &Stack{
				Containers: make([]string, 0),
			}

			// loop backwards and populate the stack
			for ii := lineIndex - 1; ii >= 0; ii-- {
				str := string(lines[ii][i])

				if str != "" && str != " " {
					s.Containers = append(s.Containers, str)
				}
			}

			stacks[num] = s
		}

		// getting the instruction data
		for i := lineIndex + 2; i < len(lines); i++ {
			line := lines[i]
			line = strings.Replace(line, "move", "", 1)
			line = strings.Replace(line, "from", ",", 1)
			line = strings.Replace(line, "to", ",", 1)
			line = strings.Replace(line, " ", "", -1)

			arr := strings.Split(line, ",")

			num1, _ := strconv.Atoi(arr[0])
			num2, _ := strconv.Atoi(arr[1])
			num3, _ := strconv.Atoi(arr[2])

			// we only want to "moveAsUnit" or the second problem.
			asUnit := !isFirstProb
			stacks[num2].Move(num1, stacks[num3], asUnit)
		}
		break
	}
	// loop through the stacks IN ORDER and get the top values.
	for i := 1; i <= maxKey; i++ {
		answer = fmt.Sprintf("%s%s", answer, stacks[i].Containers[len(stacks[i].Containers)-1])
	}
	return answer
}
