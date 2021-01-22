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

type pointCount struct {
	ground int
	tree   int
	lumber int
}

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

	fmt.Println("Round 0 /")
	printMap(northpolemap)
	for round := 1; round <= 10; round++ {
		runRound(&northpolemap)
		fmt.Println("Round", round, "/")
		printMap(northpolemap)
	}

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

func runRound(npmap *NPmap) {

	duplicate := make(NPmap, len((*npmap)))
	for i := range *npmap {
		duplicate[i] = make([]string, len((*npmap)[i]))
		copy(duplicate[i], (*npmap)[i])
	}

	for row := range duplicate {
		for col := range duplicate[row] {
			p := mapPoint(duplicate, row, col)
			switch (*npmap)[row][col] {
			case "#": // An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard and at least one acre containing trees.
				// Otherwise, it becomes open.
				if p.lumber > 0 && p.tree > 0 {
					(*npmap)[row][col] = "#"
				} else {
					(*npmap)[row][col] = "."
				}

			case "|": // An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards. Otherwise, nothing happens.
				if p.lumber > 2 {
					(*npmap)[row][col] = "#"
				}
			case ".": // An open acre will become filled with trees if three or more adjacent acres contained trees. Otherwise, nothing happens.
				if p.tree > 2 {
					(*npmap)[row][col] = "|"
				}
			}
		}
	}
}

func mapPoint(myMap NPmap, myRow, myCol int) pointCount {
	var p pointCount
	for row := -1; row < 2; row++ {
		for col := -1; col < 2; col++ {
			thisRow := myRow + row
			thisCol := myCol + col
			//fmt.Println("row:", thisRow, "- col:", thisCol)
			// If not the mappoint itself.
			if !(col == 0 && row == 0) {
				// If not above the top row or below the buttom row.
				if thisRow >= 0 && thisRow <= len(myMap)-1 {
					// If not too far to the left nor to the right on the map...
					if thisCol >= 0 && thisCol <= len(myMap[thisRow])-1 {
						//fmt.Println("Found:", myMap[thisRow][thisCol])
						switch myMap[thisRow][thisCol] {
						case "#":
							p.lumber++
						case "|":
							p.tree++
						case ".":
							p.ground++
						}
					}

				}
			}
		}
	}
	return p
}
