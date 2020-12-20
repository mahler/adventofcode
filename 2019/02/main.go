package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}

	var intOpCode []int
	operations := strings.Split(string(data), ",")
	for _, value := range operations {
		opVal, _ := strconv.Atoi(value)
		intOpCode = append(intOpCode, opVal)
	}

	//  before running the program, replace position 1 with the value 12 and replace position 2 with the value 2.
	intOpCode[1] = 12
	intOpCode[2] = 2

	result := RunProgram(intOpCode)
	fmt.Println("result:", result)
}

func RunProgram(intOpCode []int) int {
	fmt.Println(intOpCode)
	currentPosition := 0
	lastSavedValue := 0
	for {
		opCode := intOpCode[currentPosition]
		if opCode == 1 {
			// ADDition
			lastSavedValue = intOpCode[intOpCode[currentPosition+1]] + intOpCode[intOpCode[currentPosition+2]]
			intOpCode[intOpCode[currentPosition+3]] = lastSavedValue

		} else if opCode == 2 {
			// MULTIPLICATION
			lastSavedValue = intOpCode[intOpCode[currentPosition+1]] * intOpCode[intOpCode[currentPosition+2]]
			intOpCode[intOpCode[currentPosition+3]] = lastSavedValue

		} else if opCode == 99 {
			// END
			break
		}
		fmt.Println(intOpCode)
		currentPosition += 4
	}
	return lastSavedValue
}
