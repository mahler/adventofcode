package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	trans := [][2]string{}
	molecules := map[string]bool{}

	for _, line := range fileLines {
		if line == "" {
			break
		}
		re := regexp.MustCompile("^([a-zA-Z]+) => ([a-zA-Z]+)$")
		parts := re.FindStringSubmatch(line)
		trans = append(trans, [2]string{parts[1], parts[2]})
	}

	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 19, Part 1: Medicine for Rudolph")

	// Molecule is the last line of the file.
	molecule := string(fileLines[len(fileLines)-1])

	for t := range trans {
		replace(molecule, trans[t][0], trans[t][1], &molecules)
	}

	fmt.Println("How many distinct molecules can be created after all the different ways you can do one replacement on the medicine molecule?")
	fmt.Println(len(molecules))
	// ------

}

func replace(str, old, new string, next *map[string]bool) {
	p := len(str)
	for {
		p = strings.LastIndex(str[:p], old)
		if p < 0 {
			break
		}
		r := str[0:p] + new + str[p+len(old):]
		(*next)[r] = true
	}
}
