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
	fmt.Println(inverseCaptcha(dataStr))

	//-------------------------------------
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println(part2Captcha(dataStr))
}

func inverseCaptcha(line string) int {
	sum := 0
	offset := 1
	length := len(line)
	for i := 0; i < length; i++ {
		if line[i] == line[(i+offset)%length] {
			tmp, _ := strconv.Atoi(line[i : i+1])
			sum += tmp
		}
	}
	return sum
}

func part2Captcha(line string) int {
	sum := 0
	a := line[:len(line)/2]
	b := line[len(line)/2:]
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			tmp, _ := strconv.Atoi(string(a[i]))
			sum += tmp * 2
		}
	}
	return sum
}
