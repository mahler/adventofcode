package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	// Read guest preferences into rule datastructure
	rxGuestPref := regexp.MustCompile(`^(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)`)

	// Used to store seating preferences
	guestList := make(map[string]map[string]int)
	// Used as input for permutations
	guestKeys := make(map[string]bool)
	// Used to translate the letters used for permutations to guestnames.
	letter2guest := make(map[string]string)

	for _, guestPref := range lines {
		// Parse line into vars
		fields := rxGuestPref.FindStringSubmatch(guestPref)
		guestname := fields[1]
		letter := guestname[0:1]

		letter2guest[letter] = guestname
		nextTo := fields[4]
		happyDir := fields[2]
		happiness, _ := strconv.Atoi(fields[3])

		if happyDir == "lose" {
			happiness = 0 - happiness
		}

		_, exists := guestList[guestname]
		if !exists {
			guestList[guestname] = make(map[string]int)
		}
		guestList[guestname][nextTo] = happiness
		guestKeys[letter] = true
	}
	// ------------ PART 1 ------------------------
	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 13 part 1: Knights of the Dinner Table")
	permGuests := []rune{}

	for key := range guestKeys {
		rKey := []rune(key)[0]
		permGuests = append(permGuests, rKey)
	}

	combiList := []string{}
	Perm([]rune(permGuests), func(a []rune) {
		cGuest := fmt.Sprintf("%c", a)

		combiList = append(combiList, cGuest)
	})

	maxHappiness := 0
	// Loop through guest permutations (combilist)
	for _, guestRuneCombi := range combiList {
		happiness := 0
		gRC := guestRuneCombi[1 : len(guestRuneCombi)-1]
		fields := strings.Fields(gRC)
		for gPos, gLetter := range fields {
			thisGuest := letter2guest[gLetter]
			nextGuest := ""
			// Find next guest: If last guest in the circle...
			if gPos+1 > len(fields)-1 {
				nextGuest = letter2guest[fields[0]]
			} else {
				nextGuest = letter2guest[fields[gPos+1]]
			}

			// Collect happiness from both sidess
			happiness += guestList[thisGuest][nextGuest]
			happiness += guestList[nextGuest][thisGuest]
		}

		if happiness > maxHappiness {
			maxHappiness = happiness
		}
	}

	fmt.Println("What is the total change in happiness for")
	fmt.Println("the optimal seating arrangement of the actual guest list?")
	fmt.Println(maxHappiness)

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2")
	// Insert "Me" in guest list
	guestList["Me"] = make(map[string]int)
	for _, guestName := range letter2guest {
		guestList["Me"][guestName] = 0
	}
	letter2guest["X"] = "Me"
	permGuests = append(permGuests, rune('X'))

	// Part 2 setup down, let's run the planning again...
	p2combiList := []string{}
	Perm([]rune(permGuests), func(a []rune) {
		cGuest := fmt.Sprintf("%c", a)

		p2combiList = append(p2combiList, cGuest)
	})

	p2maxHappiness := 0
	// Loop through guest permutations (combilist)
	for _, guestRuneCombi := range p2combiList {
		happiness := 0
		gRC := guestRuneCombi[1 : len(guestRuneCombi)-1]
		fields := strings.Fields(gRC)
		for gPos, gLetter := range fields {
			thisGuest := letter2guest[gLetter]
			nextGuest := ""
			if gPos+1 > len(fields)-1 {
				nextGuest = letter2guest[fields[0]]
			} else {
				nextGuest = letter2guest[fields[gPos+1]]
			}
			happiness += guestList[thisGuest][nextGuest]
			happiness += guestList[nextGuest][thisGuest]
		}

		if happiness > p2maxHappiness {
			p2maxHappiness = happiness
		}
	}

	fmt.Println("What is the total change in happiness for")
	fmt.Println("the optimal seating with me at table?")
	fmt.Println(p2maxHappiness)
}

// Perm generate all possible permutations of the rune slice input
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
// Sourced from https://yourbasic.org/golang/generate-permutation-slice-string/
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
