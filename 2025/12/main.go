package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shape struct {
	ID   int
	Rows []ShapeRow
}

type ShapeRow struct {
	Value int
	Width int
}

type Grid struct {
	Width  int
	Height int
	Counts []int
}

type Variation struct {
	Height int
	Width  int
	Rows   []int
}

type Item struct {
	ID      int
	Area    int
	PlacedR int
	PlacedC int
}

func parseInput(lines []string) (map[int][]ShapeRow, []Grid) {
	shapes := make(map[int][]ShapeRow)
	var grids []Grid
	currentID := -1

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			dims := strings.Split(parts[0], "x")
			width, _ := strconv.Atoi(dims[0])
			height, _ := strconv.Atoi(dims[1])

			countStrs := strings.Fields(parts[1])
			counts := make([]int, len(countStrs))
			for i, s := range countStrs {
				counts[i], _ = strconv.Atoi(s)
			}

			grids = append(grids, Grid{Width: width, Height: height, Counts: counts})
		} else if strings.HasSuffix(line, ":") {
			currentID, _ = strconv.Atoi(line[:len(line)-1])
			shapes[currentID] = []ShapeRow{}
		} else {
			val := 0
			for i, c := range line {
				if c == '#' {
					val |= 1 << (len(line) - 1 - i)
				}
			}
			shapes[currentID] = append(shapes[currentID], ShapeRow{Value: val, Width: len(line)})
		}
	}

	return shapes, grids
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func generateVariations(baseShapes map[int][]ShapeRow) map[int][]Variation {
	variations := make(map[int][]Variation)

	for sid, rows := range baseShapes {
		maxW := 0
		for _, r := range rows {
			maxW = max(maxW, r.Width)
		}

		// Convert to bit matrix
		matrix := make([][]int, len(rows))
		for i, r := range rows {
			row := make([]int, maxW)
			for j := 0; j < maxW; j++ {
				row[j] = (r.Value >> (r.Width - 1 - j)) & 1
			}
			matrix[i] = row
		}

		seenHashes := make(map[string]bool)
		var shapeVars []Variation

		current := matrix
		// 2 flips * 4 rotations
		for flip := 0; flip < 2; flip++ {
			for rot := 0; rot < 4; rot++ {
				// Trim bounding box
				minR, maxR := len(current), -1
				minC, maxC := len(current[0]), -1
				hasBits := false

				for r := 0; r < len(current); r++ {
					for c := 0; c < len(current[0]); c++ {
						if current[r][c] != 0 {
							hasBits = true
							minR = min(minR, r)
							maxR = max(maxR, r)
							minC = min(minC, c)
							maxC = max(maxC, c)
						}
					}
				}

				if hasBits {
					trimmedH := maxR - minR + 1
					trimmedW := maxC - minC + 1
					intRows := make([]int, trimmedH)

					for r := minR; r <= maxR; r++ {
						rVal := 0
						for c := minC; c <= maxC; c++ {
							rVal = (rVal << 1) | current[r][c]
						}
						intRows[r-minR] = rVal
					}

					// Create signature
					sig := strings.Builder{}
					for _, v := range intRows {
						sig.WriteString(strconv.Itoa(v))
						sig.WriteString(",")
					}

					if !seenHashes[sig.String()] {
						seenHashes[sig.String()] = true
						shapeVars = append(shapeVars, Variation{
							Height: trimmedH,
							Width:  trimmedW,
							Rows:   intRows,
						})
					}
				}

				// Rotate 90 degrees clockwise
				h, w := len(current), len(current[0])
				next := make([][]int, w)
				for c := 0; c < w; c++ {
					next[c] = make([]int, h)
					for r := 0; r < h; r++ {
						next[c][r] = current[h-1-r][c]
					}
				}
				current = next
			}

			// Flip vertically
			for i, j := 0, len(current)-1; i < j; i, j = i+1, j-1 {
				current[i], current[j] = current[j], current[i]
			}
		}

		variations[sid] = shapeVars
	}

	return variations
}

func popCount(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func isSpaceSufficient(grid []int, width, height, requiredArea, minItemArea, slack int) bool {
	usedArea := 0
	for _, row := range grid {
		usedArea += popCount(row)
	}

	freeArea := (width * height) - usedArea
	if freeArea < requiredArea {
		return false
	}

	if freeArea > requiredArea+slack {
		return true
	}

	// Build collision set
	collision := make(map[[2]int]bool)
	for r := 0; r < height; r++ {
		rowVal := grid[r]
		if rowVal == 0 {
			continue
		}
		for c := 0; c < width; c++ {
			if (rowVal>>(width-1-c))&1 != 0 {
				collision[[2]int{r, c}] = true
			}
		}
	}

	visited := make(map[[2]int]bool)
	usableFreeArea := 0

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			pos := [2]int{r, c}
			if !collision[pos] && !visited[pos] {
				// BFS for island
				queue := [][2]int{pos}
				visited[pos] = true
				islandSize := 0

				for len(queue) > 0 {
					curr := queue[0]
					queue = queue[1:]
					islandSize++

					neighbors := [][2]int{
						{curr[0] + 1, curr[1]},
						{curr[0] - 1, curr[1]},
						{curr[0], curr[1] + 1},
						{curr[0], curr[1] - 1},
					}

					for _, next := range neighbors {
						nr, nc := next[0], next[1]
						if nr >= 0 && nr < height && nc >= 0 && nc < width {
							if !collision[next] && !visited[next] {
								visited[next] = true
								queue = append(queue, next)
							}
						}
					}
				}

				if islandSize >= minItemArea {
					usableFreeArea += islandSize
				}

				if usableFreeArea >= requiredArea {
					return true
				}
			}
		}
	}

	return usableFreeArea >= requiredArea
}

func solveRecursive(grid []int, items []Item, itemIdx, width, height int, variations map[int][]Variation, minGlobalArea int) bool {
	if itemIdx == len(items) {
		return true
	}

	item := &items[itemIdx]
	sid := item.ID

	// Check remaining space
	remainingArea := 0
	for i := itemIdx; i < len(items); i++ {
		remainingArea += items[i].Area
	}

	if !isSpaceSufficient(grid, width, height, remainingArea, minGlobalArea, 20) {
		return false
	}

	// Symmetry breaking
	startR, startC := 0, 0
	if itemIdx > 0 && items[itemIdx-1].ID == sid {
		startR = items[itemIdx-1].PlacedR
		startC = items[itemIdx-1].PlacedC
	}

	for _, variant := range variations[sid] {
		hVar, wVar, rows := variant.Height, variant.Width, variant.Rows

		for r := startR; r <= height-hVar; r++ {
			cBegin := 0
			if r == startR {
				cBegin = startC
			}
			cLimit := width - wVar + 1

			for c := cBegin; c < cLimit; c++ {
				// Collision check
				fits := true
				shift := width - c - wVar

				for i := 0; i < hVar; i++ {
					if grid[r+i]&(rows[i]<<shift) != 0 {
						fits = false
						break
					}
				}

				if fits {
					// Place shape
					for i := 0; i < hVar; i++ {
						grid[r+i] |= rows[i] << shift
					}

					item.PlacedR = r
					item.PlacedC = c

					// Recurse
					if solveRecursive(grid, items, itemIdx+1, width, height, variations, minGlobalArea) {
						return true
					}

					// Backtrack
					for i := 0; i < hVar; i++ {
						grid[r+i] ^= rows[i] << shift
					}
				}
			}
		}
	}

	return false
}

func part1(lines []string) int {
	shapes, queries := parseInput(lines)
	variations := generateVariations(shapes)

	// Pre-calculate areas
	shapeAreas := make(map[int]int)
	for sid, vars := range variations {
		if len(vars) > 0 {
			area := 0
			for _, row := range vars[0].Rows {
				area += popCount(row)
			}
			shapeAreas[sid] = area
		}
	}

	validCount := 0

	for _, query := range queries {
		width, height, counts := query.Width, query.Height, query.Counts

		// Build items list
		var items []Item
		for sid, count := range counts {
			if count > 0 {
				area := shapeAreas[sid]
				for i := 0; i < count; i++ {
					items = append(items, Item{ID: sid, Area: area})
				}
			}
		}

		if len(items) == 0 {
			validCount++
			continue
		}

		// Sort by area descending
		for i := 0; i < len(items); i++ {
			for j := i + 1; j < len(items); j++ {
				if items[i].Area < items[j].Area || (items[i].Area == items[j].Area && items[i].ID > items[j].ID) {
					items[i], items[j] = items[j], items[i]
				}
			}
		}

		totalArea := 0
		for _, item := range items {
			totalArea += item.Area
		}

		if totalArea > width*height {
			continue
		}

		minGlobalArea := items[len(items)-1].Area
		grid := make([]int, height)

		if solveRecursive(grid, items, 0, width, height, variations, minGlobalArea) {
			validCount++
		}
	}

	return validCount
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("How many of the regions can fit all of the presents listed?")
	fmt.Println(part1(lines))
}
