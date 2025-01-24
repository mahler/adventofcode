package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var rules2, rules3 [][2]string

	// Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input and parse rules
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, " => ") {
			parts := strings.Split(line, " => ")
			if len(parts[0]) == 5 { // Rules for 2x2 blocks
				rules2 = append(rules2, [2]string{
					strings.ReplaceAll(parts[0], "/", ""),
					strings.ReplaceAll(parts[1], "/", ""),
				})
			} else if len(parts[0]) == 11 { // Rules for 3x3 blocks
				rules3 = append(rules3, [2]string{
					strings.ReplaceAll(parts[0], "/", ""),
					strings.ReplaceAll(parts[1], "/", ""),
				})
			}
		}
	}

	// Define the initial grid
	grid := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}

	for iteration := 1; iteration <= 18; iteration++ {
		blockSize := determineBlockSize(len(grid))
		grid = processGrid(grid, rules2, rules3, blockSize)

		if iteration == 5 {
			fmt.Println("Part 1: How many pixels stay on after 5 iterations?")
			fmt.Println(countOn(grid))
		}
	}

	fmt.Println()
	fmt.Println("Part 2: How many pixels stay on after 18 iterations?")
	fmt.Println(countOn(grid))
}

// determineBlockSize determines the block size for the grid
func determineBlockSize(size int) int {
	if size%2 == 0 {
		return 2
	}
	return 3
}

// processGrid processes the grid based on the rules and block size
func processGrid(grid [][]int, rules2, rules3 [][2]string, blockSize int) [][]int {
	newBlockSize := blockSize + 1
	numBlocks := len(grid) / blockSize
	newGridSize := numBlocks * newBlockSize
	newGrid := make([][]int, newGridSize)
	for i := range newGrid {
		newGrid[i] = make([]int, newGridSize)
	}

	for bi := 0; bi < numBlocks; bi++ {
		for bj := 0; bj < numBlocks; bj++ {
			block := extractBlock(grid, bi, bj, blockSize)
			var transformedBlock [][]int
			if blockSize == 2 {
				transformedBlock = matchAndTransform(block, rules2)
			} else {
				transformedBlock = matchAndTransform(block, rules3)
			}
			insertBlock(newGrid, transformedBlock, bi, bj, newBlockSize)
		}
	}

	return newGrid
}

// extractBlock extracts a block from the grid
func extractBlock(grid [][]int, bi, bj, blockSize int) [][]int {
	block := make([][]int, blockSize)
	for i := 0; i < blockSize; i++ {
		block[i] = grid[bi*blockSize+i][bj*blockSize : (bj+1)*blockSize]
	}
	return block
}

// insertBlock inserts a block into the new grid
func insertBlock(grid, block [][]int, bi, bj, blockSize int) {
	for i := range block {
		copy(grid[bi*blockSize+i][bj*blockSize:], block[i])
	}
}

// matchAndTransform matches a block with rules and transforms it
func matchAndTransform(block [][]int, rules [][2]string) [][]int {
	strBlock := blockToString(block)
	for _, rule := range rules {
		if matchBlock(strBlock, rule[0]) {
			return stringToBlock(rule[1])
		}
	}
	panic("No matching rule found!")
}

// blockToString converts a block to a string
func blockToString(block [][]int) string {
	var sb strings.Builder
	for _, row := range block {
		for _, cell := range row {
			if cell == 1 {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
	}
	return sb.String()
}

// stringToBlock converts a string to a block
func stringToBlock(s string) [][]int {
	size := int(len(s))
	side := int(0)
	for side*side < size {
		side++
	}

	block := make([][]int, side)
	for i := 0; i < side; i++ {
		block[i] = make([]int, side)
		for j := 0; j < side; j++ {
			if s[i*side+j] == '#' {
				block[i][j] = 1
			}
		}
	}
	return block
}

// matchBlock checks if a block matches a rule (including rotations and flips)
func matchBlock(block, rule string) bool {
	for i := 0; i < 4; i++ { // Try all rotations
		if block == rule {
			return true
		}
		block = rotateBlock(block)
	}
	block = flipBlock(block) // Try flipped versions
	for i := 0; i < 4; i++ {
		if block == rule {
			return true
		}
		block = rotateBlock(block)
	}
	return false
}

// rotateBlock rotates a block (string representation) 90 degrees clockwise
func rotateBlock(block string) string {
	size := int(len(block))
	side := 0
	for side*side < size {
		side++
	}
	rotated := make([]rune, size)
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			rotated[j*side+(side-1-i)] = rune(block[i*side+j])
		}
	}
	return string(rotated)
}

// flipBlock flips a block (string representation) horizontally
func flipBlock(block string) string {
	size := int(len(block))
	side := 0
	for side*side < size {
		side++
	}
	flipped := make([]rune, size)
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			flipped[i*side+(side-1-j)] = rune(block[i*side+j])
		}
	}
	return string(flipped)
}

// countOn counts the number of "on" cells (1s) in the grid
func countOn(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				count++
			}
		}
	}
	return count
}
