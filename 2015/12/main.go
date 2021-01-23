package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 12, part 1: JSAbacusFramework.io")

	input, _ := ioutil.ReadFile("puzzle.txt")
	var data interface{}
	json.Unmarshal(input, &data)

	numberCount := count(data)
	fmt.Println("What is the sum of all numbers in the document?")
	fmt.Println(numberCount)

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2:")

	numberCount = recount(data)
	fmt.Println("What is the sum with red correction?")
	fmt.Println(numberCount)

}

//  Functions sourced from https://github.com/mevdschee/AdventOfCode2015/tree/master/day12
func count(val interface{}) float64 {
	sum := 0.0
	switch reflect.TypeOf(val).String() {
	case "float64":
		sum += val.(float64)
	case "[]interface {}":
		array := val.([]interface{})
		for i := 0; i < len(array); i++ {
			sum += count(array[i])
		}
	case "map[string]interface {}":
		object := val.(map[string]interface{})
		for _, v := range object {
			sum += count(v)
		}
	}
	return sum
}

func recount(val interface{}) float64 {
	sum := 0.0
	switch reflect.TypeOf(val).String() {
	case "float64":
		sum += val.(float64)
	case "[]interface {}":
		array := val.([]interface{})
		for i := 0; i < len(array); i++ {
			sum += recount(array[i])
		}
	case "map[string]interface {}":
		object := val.(map[string]interface{})
		skip := false
		for _, v := range object {
			if v == "red" {
				skip = true
				break
			}
		}
		if !skip {
			for _, v := range object {
				sum += recount(v)
			}
		}
	}
	return sum
}
