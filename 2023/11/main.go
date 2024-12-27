package main

import (
	"bufio"
	"fmt"
	"os"
)

func readImage(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var image [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		image = append(image, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return image
}

func findEmptyRows(image [][]rune) []int {
	var emptyRows []int
	for y, row := range image {
		allEmpty := true
		for _, char := range row {
			if char != '.' {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			emptyRows = append(emptyRows, y)
		}
	}
	return emptyRows
}

func findEmptyCols(image [][]rune) []int {
	var emptyCols []int
	width := len(image[0])
	for x := 0; x < width; x++ {
		allEmpty := true
		for y := 0; y < len(image); y++ {
			if image[y][x] != '.' {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			emptyCols = append(emptyCols, x)
		}
	}
	return emptyCols
}

func insertEmptyRows(image [][]rune, emptyRows []int) [][]rune {
	width := len(image[0])
	for i := len(emptyRows) - 1; i >= 0; i-- {
		row := emptyRows[i]
		newRow := make([]rune, width)
		for j := range newRow {
			newRow[j] = '.'
		}
		image = append(image[:row], append([][]rune{newRow}, image[row:]...)...)
	}
	return image
}

func insertEmptyCols(image [][]rune, emptyCols []int) [][]rune {
	for i := len(emptyCols) - 1; i >= 0; i-- {
		col := emptyCols[i]
		for row := range image {
			image[row] = append(image[row][:col], append([]rune{'.'}, image[row][col:]...)...)
		}
	}
	return image
}

func findGalaxyCoords(image [][]rune) [][2]int {
	var galaxyCoords [][2]int
	for y, row := range image {
		for x, char := range row {
			if char == '#' {
				galaxyCoords = append(galaxyCoords, [2]int{x, y})
			}
		}
	}
	return galaxyCoords
}

func coordsAfterExpansion(coords [2]int, multiplier int, emptyCols, emptyRows []int) [2]int {
	emptyColsBefore := 0
	for _, col := range emptyCols {
		if col < coords[0] {
			emptyColsBefore++
		}
	}

	emptyRowsBefore := 0
	for _, row := range emptyRows {
		if row < coords[1] {
			emptyRowsBefore++
		}
	}

	newX := coords[0] + emptyColsBefore*(multiplier-1)
	newY := coords[1] + emptyRowsBefore*(multiplier-1)
	return [2]int{newX, newY}
}

func manhattanDistance(coord1, coord2 [2]int) int {
	return abs(coord1[0]-coord2[0]) + abs(coord1[1]-coord2[1])
}

func calculateShortestPaths(galaxyCoords [][2]int) int {
	sum := 0
	for i := 0; i < len(galaxyCoords); i++ {
		for j := i + 1; j < len(galaxyCoords); j++ {
			sum += manhattanDistance(galaxyCoords[i], galaxyCoords[j])
		}
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	image := readImage("input.txt")

	emptyRows := findEmptyRows(image)
	emptyCols := findEmptyCols(image)

	image = insertEmptyRows(image, emptyRows)
	image = insertEmptyCols(image, emptyCols)

	galaxyCoords := findGalaxyCoords(image)
	shortestPathsSum := calculateShortestPaths(galaxyCoords)

	fmt.Println("Part 1: Expand the universe, then find the length of the shortest path between")
	fmt.Println("every pair of galaxies. What is the sum of these lengths?")
	fmt.Println(shortestPathsSum)

	// --- Part 2
	image = readImage("input.txt")

	emptyRows = findEmptyRows(image)
	emptyCols = findEmptyCols(image)

	var galaxyCoordsp2 [][2]int
	expansionFactor := 1_000_000

	for y, row := range image {
		for x, char := range row {
			if char == '#' {
				newCoords := coordsAfterExpansion([2]int{x, y}, expansionFactor, emptyCols, emptyRows)
				galaxyCoordsp2 = append(galaxyCoordsp2, newCoords)
			}
		}
	}

	shortestPathsSum = calculateShortestPaths(galaxyCoordsp2)
	fmt.Println()
	fmt.Println("Part 2: Starting with the same initial image, expand the universe according to these new rules,")
	fmt.Println("then find the length of the shortest path between every pair of galaxies. What is the sum of these lengths?")
	fmt.Println(shortestPathsSum)
}
