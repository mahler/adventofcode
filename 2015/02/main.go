package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type xmasPackage struct {
	length int
	width  int
	height int
}

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println()
	fmt.Println("Day 02, part 1: I Was Told There Would Be No Math")

	packageList := []xmasPackage{}

	for _, fileRow := range fileLines {
		field := strings.Split(fileRow, "x")
		var myPack xmasPackage

		myPack.height, _ = strconv.Atoi(field[0])
		myPack.width, _ = strconv.Atoi(field[1])
		myPack.length, _ = strconv.Atoi(field[2])

		packageList = append(packageList, myPack)
	}
	fmt.Println(len(packageList), "on the wrapping list")
	totalWrappingPaper := 0
	for _, pack := range packageList {
		totalWrappingPaper += pack.size()
		// add slack
		totalWrappingPaper += pack.slack()
	}

	fmt.Println("How many total square feet of wrapping paper should they order?")
	fmt.Println(totalWrappingPaper)
	// ------------ PART 2 ------------------------

	fmt.Println()
	fmt.Println("Part 2: Ribbon Calc")
	ribbon := 0
	for _, pack := range packageList {
		ribbon += pack.ribbonLenght()
		ribbon += pack.ribbonBow()
	}
	fmt.Println("Ribbon to order:", ribbon)

}

func (gift xmasPackage) size() int {
	return (2 * gift.length * gift.width) + (2 * gift.width * gift.height) + (2 * gift.height * gift.length)
}

func (gift xmasPackage) slack() int {
	var size []int
	size = append(size, gift.height)
	size = append(size, gift.width)
	size = append(size, gift.length)

	sort.Ints(size)
	return size[0] * size[1]
}

func (gift xmasPackage) ribbonBow() int {
	return gift.height * gift.width * gift.length
}

func (gift xmasPackage) ribbonLenght() int {
	var ribbon []int
	ribbon = append(ribbon, gift.height)
	ribbon = append(ribbon, gift.width)
	ribbon = append(ribbon, gift.length)

	sort.Ints(ribbon)
	return ribbon[0] + ribbon[0] + ribbon[1] + ribbon[1]
}
