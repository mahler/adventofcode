package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 6: Signals and Noise")
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	codeWord := ""
	fileRows := strings.Split(string(fileContent), "\n")
	for charNum := 0; charNum < len(fileRows[0]); charNum++ {
		counterMap := make(map[string]int)
		for _, fileRow := range fileRows {
			char := fileRow[charNum : charNum+1]

			if _, ok := counterMap[char]; !ok {
				counterMap[char] = 1
			} else {
				counterMap[char]++
			}

		}

		// Sort the map by value (letter occurence)
		rankedLetterCount := rankByCountDesc(counterMap)

		codeWord += rankedLetterCount[0].Key
	}
	fmt.Println("The secret message (desc):", codeWord)
	//----------------------------------------------
	fmt.Println()
	fmt.Println("Part 2")
	// Reset code word
	codeWord = ""
	for charNum := 0; charNum < len(fileRows[0]); charNum++ {
		counterMap := make(map[string]int)
		for _, fileRow := range fileRows {
			char := fileRow[charNum : charNum+1]

			if _, ok := counterMap[char]; !ok {
				counterMap[char] = 1
			} else {
				counterMap[char]++
			}

		}

		// Sort the map by value (letter occurence)
		rankedLetterCount := rankByCountAsc(counterMap)

		codeWord += rankedLetterCount[0].Key
	}
	fmt.Println("The secret message (asc):", codeWord)

}

func rankByCountDesc(counterMap map[string]int) PairList {
	pl := make(PairList, len(counterMap))
	i := 0
	for k, v := range counterMap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func rankByCountAsc(counterMap map[string]int) PairList {
	pl := make(PairList, len(counterMap))
	i := 0
	for k, v := range counterMap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
