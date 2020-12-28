package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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
	fmt.Println("DAY04, Part 1: Repose Record")
	fmt.Println("Total puzzle input: ", len(fileContent))

	// --------------------------------------------
	shiftRegex := regexp.MustCompile(`\[\d+-\d+-\d+ \d+:\d+\] Guard #(\d+) begins shift`)
	asleepRegex := regexp.MustCompile(`\[\d+-\d+-\d+ 00:(\d+)\] falls asleep`)
	awakeRegex := regexp.MustCompile(`\[\d+-\d+-\d+ 00:(\d+)\] wakes up`)

}
