package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	fileContent, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	fmt.Println("Instructions in dataset:", len(lines))
	//	fmt.Println("lines", len(lines))
	sum := 0
	for _, line := range lines {
		re := regexp.MustCompile("(\\\\\\\\)|(\\\\\\\")|(\\\\x[0-9a-f]{2})")
		matches := re.FindAllStringSubmatch(line, -1)
		sum += 2
		count := len(matches)
		for i := 0; i < count; i++ {
			sum += len(matches[i][0]) - 1
		}
	}
	fmt.Println()
	fmt.Println("Day 8: Matchsticks")
	fmt.Println("Total for the entire file:", sum)

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2")
	sum = 0
	for _, line := range lines {
		re1 := regexp.MustCompile("\\\\")
		new := re1.ReplaceAllLiteralString(line, "\\\\")

		re2 := regexp.MustCompile("\\\"")
		new = re2.ReplaceAllLiteralString(new, "\\\"")
		sum += 2 + len(new) - len(line)
	}
	fmt.Println(sum)

}
