package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 4: Security Through Obscurity")
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	regexpData := regexp.MustCompile(`(.*)-(\d+)\[(\w+)\]`)

	sectorSum := 0

	p2Data := make(map[string]int)

	fileRows := strings.Split(string(fileContent), "\n")
	for _, fileRow := range fileRows {

		fields := regexpData.FindStringSubmatch(fileRow)

		// fmt.Println(fields)
		encryptedName := fields[1]
		sectorID, _ := strconv.Atoi(fields[2])
		checkSum := fields[3]

		eName := strings.ReplaceAll(encryptedName, "-", "")

		roomChecksum := calculateChecksum(eName)

		if roomChecksum == checkSum {
			sectorSum += sectorID
		}
		// -------------------------------------------------
		fmt.Println()
		fmt.Println("Part 2: northpole object storage")

	}

	fmt.Println("sum of the sector IDs of the real rooms?")
	fmt.Println(sectorSum)
}

// Originally sourced from https://github.com/Atvaark/AdventOfCode2016/blob/master/day04/main.go
func calculateChecksum(roomname string) string {
	charCountMap := make(map[rune]int)
	for _, char := range roomname {
		charCountMap[char] = charCountMap[char] + 1
	}

	// swap char and count pairs
	countCharsMap := make(map[int][]rune)
	for char, count := range charCountMap {
		chars := countCharsMap[count]
		chars = append(chars, char)
		countCharsMap[count] = chars
	}

	// sort counts descending
	var counts []int
	for count := range countCharsMap {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	checksumIndex := 0
	var checksum [5]rune
	for _, count := range counts {
		// sort chars alphabetically
		var chars runes
		chars = countCharsMap[count]
		sort.Sort(chars)

		// build result
		for _, char := range chars {
			checksum[checksumIndex] = char
			checksumIndex = checksumIndex + 1
			if checksumIndex > 4 {
				return string(checksum[:])
			}
		}
	}

	return "invalid_checksum"
}

// sort interface for a rune slice
type runes []rune

func (s runes) Len() int           { return len(s) }
func (s runes) Less(i, j int) bool { return s[i] < s[j] }
func (s runes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
