package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}

	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	field := [100][100]bool{}

	for y, line := range fileLines {
		for x := 0; x < 100; x++ {
			if line[x] == '#' {
				field[x][y] = true
			}
		}
	}
	// Backup for part 2
	p2field := field

	for i := 0; i < 100; i++ {
		grid := [100][100]bool{}
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				neighbors := 0
				for dx := x - 1; dx <= x+1; dx++ {
					for dy := y - 1; dy <= y+1; dy++ {
						// Skip self
						if dx == x && dy == y {
							continue
						}
						// Skip border
						if dx < 0 || dx >= 100 {
							continue
						}
						if dy < 0 || dy >= 100 {
							continue
						}
						// Check position in field
						if field[dx][dy] {
							neighbors++
						}
					}
				}
				// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
				if field[x][y] {
					if neighbors == 2 || neighbors == 3 {
						grid[x][y] = true
					}
				} else {
					// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
					if neighbors == 3 {
						grid[x][y] = true
					}
				}
			}
		}
		field = grid
	}
	//	printGrid(field)

	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 18, part 1: Like a GIF For Your Yard")
	fmt.Println(countLightsOn(field))

	p2field[0][0] = true
	p2field[99][0] = true
	p2field[0][99] = true
	p2field[99][99] = true

	for i := 0; i < 100; i++ {
		grid := [100][100]bool{}
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				neighbors := 0
				for dx := x - 1; dx <= x+1; dx++ {
					for dy := y - 1; dy <= y+1; dy++ {
						// Skip self
						if dx == x && dy == y {
							continue
						}
						// Skip borders
						if dx < 0 || dx >= 100 {
							continue
						}
						if dy < 0 || dy >= 100 {
							continue
						}
						// Check position in field
						if p2field[dx][dy] {
							neighbors++
						}
					}
				}
				if p2field[x][y] {
					if neighbors == 2 || neighbors == 3 {
						grid[x][y] = true
					}
				} else {
					if neighbors == 3 {
						grid[x][y] = true
					}
				}
			}
		}
		p2field = grid
		// Refreeze the 4 cornors
		p2field[0][0] = true
		p2field[99][0] = true
		p2field[0][99] = true
		p2field[99][99] = true
	}

	// ------------ PART 2 ------------------------
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("How many lights are on after 100 steps with 4 cornor lights always on?")
	fmt.Println(countLightsOn(p2field))

}

func countLightsOn(grid [100][100]bool) int {
	lightsOn := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] {
				lightsOn++
			}
		}
	}
	return lightsOn
}

func printGrid(grid [100][100]bool) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
