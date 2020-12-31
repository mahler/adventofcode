package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	fmt.Println("Guest list:", len(lines))

	// Read guest preferences into rule datastructure
	rxGuestPref := regexp.MustCompile(`^(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)`)

	guestList := make(map[string]map[string]int)
	guestKeys := make(map[string]bool)
	for _, guestPref := range lines {
		// Parse line into vars
		fields := rxGuestPref.FindStringSubmatch(guestPref)
		guestname := fields[1]
		letter := guestname[0:1]
		nextTo := fields[4]
		happyDir := fields[2]
		happiness, _ := strconv.Atoi(fields[3])

		if happyDir == "lose" {
			happiness = 0 - happiness
		}
		//fmt.Println("gname:", guestname, "nextTo:", nextTo, "letter:", letter, "happiness:", happiness)

		_, exists := guestList[guestname]
		if !exists {
			guestList[guestname] = make(map[string]int)
		}
		guestList[guestname][nextTo] = happiness
		guestKeys[letter] = true
	}
	// Part 1 ------------------------------------------

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

	fmt.Println("Permutations:", len(combiList))

	maxHappiness := 0
	// Loop through guest permutations (combilist)
	// Calculate the happiness
	// save the highest total hapiness as result for part one.
	// Bonus: Consider saving the permutation which gave the highest happiness.

	fmt.Println("What is the *total change in happiness for the optimal seating arrangement of the actual guest list?")
	fmt.Println(maxHappiness)
	// Part 2 -----------------------------

}

// Perm generate all possible permutations of the rune slice input
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
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
