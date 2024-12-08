package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func antennaPositions(grid [][]rune) map[rune][]Point {
	antennas := make(map[rune][]Point)
	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], Point{x, y})
			}
		}
	}
	return antennas
}

func findAntinodes(antennas map[rune][]Point, xRange, yRange []int, resonants []int) map[Point]bool {
	antinodes := make(map[Point]bool)

	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				x1, y1 := locations[i].x, locations[i].y
				x2, y2 := locations[j].x, locations[j].y

				for _, resonant := range resonants {
					a1x := x1 + resonant*(x1-x2)
					a1y := y1 + resonant*(y1-y2)

					if !contains(xRange, a1x) || !contains(yRange, a1y) {
						break
					}
					antinodes[Point{a1x, a1y}] = true
				}

				for _, resonant := range resonants {
					a2x := x2 + resonant*(x2-x1)
					a2y := y2 + resonant*(y2-y1)

					if !contains(xRange, a2x) || !contains(yRange, a2y) {
						break
					}
					antinodes[Point{a2x, a2y}] = true
				}
			}
		}
	}
	return antinodes
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func generateRange(start, end int) []int {
	result := []int{}
	for i := start; i < end; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	data, _ := os.Open("input.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var grid [][]rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	antennas := antennaPositions(grid)

	// Part 1: resonants = 1..1
	antinodes := findAntinodes(antennas, generateRange(0, len(grid[0])), generateRange(0, len(grid)), []int{1})
	fmt.Println("Part 1: How many unique locations within the bounds of the map contain an antinode?")
	fmt.Println(len(antinodes))

	// Part 2: resonants = 0..
	var resonants []int
	for i := 0; ; i++ {
		resonants = append(resonants, i)
		if i > len(grid[0]) && i > len(grid) {
			break
		}
	}

	antinodes = findAntinodes(antennas, generateRange(0, len(grid[0])), generateRange(0, len(grid)), resonants)
	fmt.Println()
	fmt.Println("Part 2: How many unique locations within the bounds of the map contain an antinode?")
	fmt.Println(len(antinodes))
}
