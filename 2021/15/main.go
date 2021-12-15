package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type pos struct {
	i, j int
	val  uint16
}
type minHeap []pos

func main() {
	fmt.Println()
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	// Setup
	fileRows := string(fileContent)
	fileLines := strings.Split(strings.TrimSpace(fileRows), "\n")
	// Use Dijkstra's
	m := len(fileLines)
	n := len(fileLines[0])
	var grid [100][100]byte
	var seen [100][100]bool
	var dist [100][100]uint16
	for i, row := range fileLines {
		for j := range row {
			grid[i][j] = fileLines[i][j] - '0'
			dist[i][j] = math.MaxUint16
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	// Keep a min heap of distances
	h := make(minHeap, 1, 10000)
	h[0] = pos{0, 0, 0}

	// While there are entries in the min heap (always true for this)
	var lowRiskPath int
	for {
		x := heap.Pop(&h).(pos)
		seen[x.i][x.j] = true
		if x.i == m-1 && x.j == n-1 {
			lowRiskPath = int(x.val)
			break
		}
		for _, nei := range [][2]int{
			{x.i + 1, x.j}, {x.i - 1, x.j}, {x.i, x.j - 1}, {x.i, x.j + 1},
		} {
			ii, jj := nei[0], nei[1]
			if !ok(ii, jj) || seen[ii][jj] {
				continue
			}
			risk := x.val + uint16(grid[ii][jj])
			if risk >= dist[ii][jj] {
				continue
			}
			dist[ii][jj] = risk
			heap.Push(&h, pos{ii, jj, risk})
		}
	}

	fmt.Println("Low Risk path:", lowRiskPath)

	// ---------------------------------------
	fmt.Println(fileLines)
}

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pos))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
