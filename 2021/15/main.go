package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xyPoint struct {
	x int
	y int
}

type queueItem struct {
	pos       xyPoint
	riskLevel int
	index     int // The index of the item in the heap.
}

// A RiskQueue implements heap.Interface and holds queueItems.
type RiskQueue []queueItem

func (rq RiskQueue) Len() int { return len(rq) }
func (rq RiskQueue) Less(i, j int) bool {
	return rq[i].riskLevel < rq[j].riskLevel
}
func (rq RiskQueue) Swap(i, j int) {
	rq[i], rq[j] = rq[j], rq[i]
	rq[i].index = i
	rq[j].index = j
}
func (rq *RiskQueue) Push(x interface{}) {
	n := len(*rq)
	item := x.(queueItem)
	item.index = n
	*rq = append(*rq, item)
}
func (rq *RiskQueue) Pop() interface{} {
	old := *rq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*rq = old[0 : n-1]
	return item
}

func main() {
	fmt.Println()
	fileContent, _ := os.ReadFile("puzzle.txt")
	fileRows := string(fileContent)
	fileLines := strings.Split(strings.TrimSpace(fileRows), "\n")

	// Setup
	dx := [4]int{0, 0, -1, 1}
	dy := [4]int{-1, 1, 0, 0}
	grid := make(map[xyPoint]int)

	maxX, maxY := 0, 0
	for x, line := range fileLines {
		var numberList []int
		for _, word := range strings.Split(line, "") {
			num, _ := strconv.Atoi(word)
			numberList = append(numberList, num)
		}
		for y, v := range numberList {
			grid[xyPoint{x, y}] = v
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	start := xyPoint{0, 0}
	target := xyPoint{maxX, maxY}
	risk := func(pos xyPoint) int {
		og := xyPoint{pos.x % (maxX + 1), pos.y % (maxY + 1)}
		risk := grid[og] +
			(pos.x)/(maxX+1) + (pos.y)/(maxY+1)
		if risk > 9 {
			return risk - 9
		}
		return risk
	}

	shortestAt := make(map[xyPoint]int)
	rq := make(RiskQueue, 0)
	heap.Init(&rq)
	rq.Push(queueItem{pos: start, riskLevel: 0})
	for rq.Len() > 0 {
		head := heap.Pop(&rq).(queueItem)
		for i := 0; i < 4; i++ {
			next := xyPoint{head.pos.x + dx[i], head.pos.y + dy[i]}
			if next.x > target.x || next.x < 0 || next.y > target.y || next.y < 0 {
				continue
			}
			nextRisk := head.riskLevel + risk(next)
			if sAt, ok := shortestAt[next]; ok && sAt <= nextRisk {
				continue
			}
			shortestAt[next] = nextRisk
			rq.Push(queueItem{pos: next, riskLevel: nextRisk})
		}
	}
	fmt.Println()
	fmt.Println("Day 2021-15:")
	fmt.Println("Part 1/")
	fmt.Println("what is the lowest total risk of any path from the top left to the bottom right?")
	fmt.Println(shortestAt[target])

	// --------------------------------------------------
	// Same startposition as Part1, but target different.
	p2target := xyPoint{(maxX+1)*5 - 1, (maxY+1)*5 - 1}
	p2shortestAt := make(map[xyPoint]int)
	p2rq := make(RiskQueue, 0)
	heap.Init(&p2rq)
	p2rq.Push(queueItem{pos: start, riskLevel: 0})
	for p2rq.Len() > 0 {
		head := heap.Pop(&p2rq).(queueItem)
		for i := 0; i < 4; i++ {
			next := xyPoint{head.pos.x + dx[i], head.pos.y + dy[i]}
			if next.x > p2target.x || next.x < 0 || next.y > p2target.y || next.y < 0 {
				continue
			}
			nextRisk := head.riskLevel + risk(next)
			if sAt, ok := p2shortestAt[next]; ok && sAt <= nextRisk {
				continue
			}
			p2shortestAt[next] = nextRisk
			p2rq.Push(queueItem{pos: next, riskLevel: nextRisk})
		}
	}

	fmt.Println()
	fmt.Println("Part 2/")
	fmt.Println("what is the lowest total risk of any path from the top left to the bottom right?")
	fmt.Println(p2shortestAt[p2target])
}
