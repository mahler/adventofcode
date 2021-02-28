package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := os.ReadFIle("puzzle.tst")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("DAY12, Part 1: Subterranean Sustainability")

	initialState := fileContent[0][15:]
	fmt.Println(initialState)
	rulesInput := fileContent[2:]

	rules := make(map[string]string)
	for _, rule := range rulesInput {
		key := rule[0:5]
		value := rule[9:]
		fmt.Println(key, "XXX", value)

		rules[key] = value
	}

	for gen := 0; gen < 20; gen++ {
		for 

	}

	fmt.Println("After 20 generations, what is the sum of the numbers of all pots which contain a plant?")
}
