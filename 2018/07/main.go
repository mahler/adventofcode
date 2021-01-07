package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("DAY07, Part 1: The Sum of Its Parts")

	var work []string
	var current string

	instructions := make(map[string][]string)
	var complete []string
	//var order []string

	for _, fileLine := range fileContent {
		s := strings.Split(fileLine, " ")
		instructions[s[7]] = append(instructions[s[7]], s[1])
		if _, ok := instructions[s[1]]; !ok {
			instructions[s[1]] = make([]string, 0)
		}
	}

	for len(instructions) > 0 {
		for k, v := range instructions {
			for count, i := range v {
				if i == string(current) {
					instructions[k] = append(instructions[k][:count], instructions[k][count+1:]...)
				}
			}
			if len(instructions[k]) == 0 {
				work = append(work, k)
			}
		}
		sort.Strings(work)
		current = work[0]
		delete(instructions, current)
		complete = append(complete, work[0])
		work = nil
	}
	fmt.Println("In what order should the steps in your instructions be completed?")
	fmt.Println(strings.Join(complete, ""))
}
