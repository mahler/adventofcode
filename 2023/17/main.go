package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	DIRS = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

type State struct {
	cost, x, y, dir int
	priority        int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func inRange(pos []int, arr [][]int) bool {
	return pos[0] >= 0 && pos[0] < len(arr) && pos[1] >= 0 && pos[1] < len(arr[0])
}

func run(ll [][]int, minDist, maxDist int) int {
	q := make(PriorityQueue, 0)
	heap.Init(&q)
	heap.Push(&q, &State{cost: 0, x: 0, y: 0, dir: -1, priority: 0})

	seen := make(map[string]bool)
	costs := make(map[string]int)

	for q.Len() > 0 {
		curr := heap.Pop(&q).(*State)
		cost, x, y, dd := curr.cost, curr.x, curr.y, curr.dir

		if x == len(ll)-1 && y == len(ll[0])-1 {
			return cost
		}

		key := fmt.Sprintf("%d,%d,%d", x, y, dd)
		if seen[key] {
			continue
		}
		seen[key] = true

		for direction := 0; direction < 4; direction++ {
			if direction == dd || (direction+2)%4 == dd {
				continue
			}

			costIncrease := 0
			for distance := 1; distance <= maxDist; distance++ {
				xx := x + DIRS[direction][0]*distance
				yy := y + DIRS[direction][1]*distance

				if !inRange([]int{xx, yy}, ll) {
					break
				}

				costIncrease += ll[xx][yy]
				if distance < minDist {
					continue
				}

				newCost := cost + costIncrease
				newKey := fmt.Sprintf("%d,%d,%d", xx, yy, direction)

				if existingCost, exists := costs[newKey]; exists && existingCost <= newCost {
					continue
				}

				costs[newKey] = newCost
				heap.Push(&q, &State{
					cost:     newCost,
					x:        xx,
					y:        yy,
					dir:      direction,
					priority: newCost,
				})
			}
		}
	}
	return -1
}

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	ll := make([][]int, len(lines))
	for i, line := range lines {
		ll[i] = make([]int, len(line))
		for j, ch := range line {
			ll[i][j], _ = strconv.Atoi(string(ch))
		}
	}

	fmt.Println("Part 1: What is the least heat loss it can incur?")
	fmt.Println(run(ll, 1, 3))
	fmt.Println()
	fmt.Println("Part 2: What is the least heat loss it can incur?")
	fmt.Println(run(ll, 4, 10))
}
