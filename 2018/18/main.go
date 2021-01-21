package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// . - Open Ground
// | - Tree
// # - Lumberyard

// NPmap defined to allow map to be passed to function(s)
type NPmap [][]string

func main() {

	data, err := ioutil.ReadFile("puzzle.tst")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("DAY18, Part 1: Settlers of The North Pole")

	var northpolemap NPmap

	for row := range fileContent {
		var mapRow []string
		for col := range fileContent[row] {
			mapRow = append(mapRow, fileContent[row][col:col+1])
		}
		northpolemap = append(northpolemap, mapRow)
	}

	printMap(northpolemap)
	runRound(&northpolemap)
	printMap(northpolemap)
	//	test(&northpolemap)
}

func printMap(npmap NPmap) {
	fmt.Println()
	for row := range npmap {
		for col := range npmap[row] {
			fmt.Print(npmap[row][col])
		}
		fmt.Println()
	}
}

func test(npmap *NPmap) {
	for x := range *npmap {
		fmt.Println((*npmap)[x])

	}
}

func runRound(npmap *NPmap) {
	for row := range *npmap {
		for col := range (*npmap)[row] {
			switch (*npmap)[row][col] {
			case ".":
				treeCount := 0
				// An open acre will become filled with trees if three or more adjacent acres contained trees. Otherwise, nothing happens.
				if row-1 >= 0 { // If no row above
					if col-1 >= 0 && (*npmap)[row-1][col-1] == "|" { // Not beyond map border left.
						treeCount++
					}
					if (*npmap)[row-1][col] == "|" {
						treeCount++
					}
					if col+1 < len((*npmap)[row]) && (*npmap)[row-1][col+1] == "|" { // Not beyond map border right.
						treeCount++
					}
				} // row less than 0, nothing here...
				if col-1 >= 0 && (*npmap)[row][col-1] == "|" {
					treeCount++
				}
				if col+1 > len((*npmap)[row]) && (*npmap)[row][col+1] == "|" {
					treeCount++
				}
				if row+1 < len((*npmap)) {
					if col-1 >= 0 && (*npmap)[row+1][col-1] == "|" { // Not beyond map border left.
						treeCount++
					}
					if (*npmap)[row+1][col] == "|" {
						treeCount++
					}
					if col+1 < len((*npmap)[row]) && (*npmap)[row+1][col+1] == "|" { // Not beyond map border right.
						treeCount++
					}
				}
				if treeCount > 2 {
					(*npmap)[row][col] = "|"
				}
			case "|":
				// An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards. Otherwise, nothing happens.

			case "#":
				// An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.
			}
		}
		fmt.Println()
	}
}
