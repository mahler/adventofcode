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

	fmt.Println("How many distinct molecules can be created after all the different ways")
	fmt.Println("you can do one replacement on the medicine molecule?")
	fmt.Println(len(molecules))

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2")

	p2molecule := string(fileLines[len(fileLines)-1])
	p2molecules := map[string]bool{}
	p2molecules[p2molecule] = true

	steps := 0
	for {
		steps++
		next := map[string]bool{}
		for p2molecule := range p2molecules {
			for t := range trans {
				if replaced(p2molecule, trans[t][1], trans[t][0], &next) {
					break
				}
			}
		}
		p2molecules = next
		if p2molecules["e"] {
			break
		}
	}

	fmt.Println("How long (how many steps) will it take to make the medicine?")
	fmt.Println(steps)
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

func replaced(str, old, new string, next *map[string]bool) bool {
	replaced := false
	p := len(str)
	for {
		p = strings.LastIndex(str[:p], old)
		if p < 0 {
			break
		}
		r := str[0:p] + new + str[p+len(old):]
		(*next)[r] = true
		replaced = true
	}
	return replaced
}
