package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//test1 := "1122"
	//test2 := "1111"
	//test3 := "1234"
	//test4 := "91212129"
	//
	//fmt.Println(inverseCaptcha(test1, 1))
	//fmt.Println(inverseCaptcha(test2, 1))
	//fmt.Println(inverseCaptcha(test3, 1))
	//fmt.Println(inverseCaptcha(test4, 1))
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	dataStr := string(data)

	fmt.Println()
	fmt.Println("2017 - Day 01 Part 1")
	fmt.Println(inverseCaptcha(dataStr, 1))

}

func inverseCaptcha(line string, offset int) int {
	sum := 0
	length := len(line)
	for i := 0; i < length; i++ {
		if line[i] == line[(i+offset)%length] {
			tmp, _ := strconv.Atoi(line[i : i+1])
			sum += tmp
		}
	}
	return sum
}
