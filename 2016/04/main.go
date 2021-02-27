package main

import (
	"fmt"
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
	fileContent, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	regexpData := regexp.MustCompile(`(.*)-(\d+)\[(\w+)\]`)

	sectorSum := 0

	p2data := make(map[string]int)

	fileRows := strings.Split(string(fileContent), "\n")
	for _, fileRow := range fileRows {

		fields := regexpData.FindStringSubmatch(fileRow)

		encryptedName := fields[1]
		sectorID, _ := strconv.Atoi(fields[2])
		checkSum := fields[3]

		p2data[encryptedName] = sectorID

		eName := strings.ReplaceAll(encryptedName, "-", "")

		roomChecksum := calculateChecksum(eName)

		if roomChecksum == checkSum {
			sectorSum += sectorID
		}

	}

	fmt.Println("sum of the sector IDs of the real rooms?")
	fmt.Println(sectorSum)

	// -------------------------------------------------
	fmt.Println()
	fmt.Println("Part 2: northpole object storage")

	for eName, sectorid := range p2data {
		decryptedName := decryptName(eName, sectorid)
		if decryptedName == "northpole object storage" {
			fmt.Println("The Northpole Object Storage is located in sector", sectorid)
			break
		}
	}
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

func decryptName(encryptedName string, seed int) string {
	minAlph := int('a')
	maxAlph := int('z')
	diffAlph := maxAlph - minAlph + 1

	decrypted := make([]rune, len(encryptedName))
	for i, char := range encryptedName {
		if char == '-' {
			decrypted[i] = ' '
			continue
		}

		value := int(char)
		decryptedValue := minAlph + ((value - minAlph + seed) % diffAlph)
		decrypted[i] = rune(decryptedValue)
	}

	return string(decrypted[:])
}
