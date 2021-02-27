package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 25: Clock Signal")
	fileContent, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	intResult := 0
	fileRows := strings.Split(string(fileContent), "\n")
	for i := 0; i < 100000; i++ {
		result := execute(fileRows, map[string]int{"a": i})
		if result == "01010101010101010101" {
			intResult = i
			break
		}
	}
	fmt.Println("What is the lowest positive integer...?")
	fmt.Println(intResult)
}

func execute(rows []string, registers map[string]int) string {
	outputs := 20
	result := []rune{}
	for pPointer := 0; pPointer < len(rows); pPointer++ {
		command := strings.Fields(rows[pPointer])
		if len(command) == 0 {
			continue
		}
		switch command[0] {
		case "cpy":
			source := command[1]
			destination := command[2]
			var val int
			if rune(source[0]) >= 'a' {
				val = registers[source]
			} else {
				val, _ = strconv.Atoi(source)
			}
			registers[destination] = val
			break

		case "inc":
			registers[command[1]] = registers[command[1]] + 1
			break
		case "dec":
			registers[command[1]] = registers[command[1]] - 1
			break
		case "jnz":
			var val int
			if rune(command[1][0]) >= 'a' {
				val = registers[command[1]]
			} else {
				val, _ = strconv.Atoi(command[1])
			}
			if val != 0 {
				tmp, _ := strconv.Atoi(command[2])
				pPointer = pPointer + tmp - 1
			}
			break
		case "out":
			source := command[1]
			var val int
			if rune(source[0]) >= 'a' {
				val = registers[source]
			} else {
				val, _ = strconv.Atoi(source)
			}
			result = append(result, rune(val+48))
			outputs--
			if outputs == 0 {
				return string(result)
			}
			break

		default:
			panic(fmt.Sprintf("command not implemented %s", command))
		}
	}
	return ""
}
