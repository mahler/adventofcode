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
		return
	}

	input := strings.TrimSpace(string(fileContent))

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 9: Explosives in Cyberspace")

	decompressResult := decompress(input, false)
	fmt.Println("What is the decompressed length of the file (your puzzle input)?")
	fmt.Println(decompressResult)

	decompressResult = decompress(input, true)
	fmt.Println("What is the decompressed length of the file using this improved format?")
	fmt.Println(decompressResult)

}

func decompress(input string, recursive bool) int {
	calcLength := 0
	rxMarker := regexp.MustCompile(`\((\d+)x(\d+)\)`)
	inputPos := rxMarker.FindStringIndex(input)
	for inputPos != nil {
		data := rxMarker.FindStringSubmatch(input[inputPos[0]:inputPos[1]])
		sourceLength, _ := strconv.Atoi(data[1])
		repeats, _ := strconv.Atoi(data[2])
		decompressedLength := sourceLength
		if recursive {
			decompressedLength = decompress(input[inputPos[1]:inputPos[1]+sourceLength], recursive)
		}
		calcLength += inputPos[0] + repeats*decompressedLength
		input = input[inputPos[1]+sourceLength:]
		inputPos = rxMarker.FindStringIndex(input)
	}
	calcLength += len(input)
	return calcLength
}
