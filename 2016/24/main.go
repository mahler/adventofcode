package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

// createMaze converts input strings to 2D array and finds marked points
func createMaze(input []string) ([][]int, []Point) {
	maze := make([][]int, len(input))
	markedPoints := make(map[int]Point)

	// Convert maze to 2D array
	for i, row := range input {
		maze[i] = make([]int, len(row))
		for j, c := range row {
			if c == '#' {
				maze[i][j] = -1
			} else {
				maze[i][j] = 0
			}

			// Store marked points
			if c >= '0' && c <= '9' {
				markedPoints[int(c-'0')] = Point{i, j}
			}
		}
	}

	// Convert map to sorted slice
	points := make([]Point, len(markedPoints))
	for i := 0; i < len(markedPoints); i++ {
		points[i] = markedPoints[i]
	}

	return maze, points
}

// step implements BFS to find shortest path
func step(maze [][]int, ends []Point, p int, to Point) int {
	if len(ends) == 0 {
		return -1 // No path found
	}

	// Copy maze for modification
	mazeCopy := make([][]int, len(maze))
	for i := range maze {
		mazeCopy[i] = make([]int, len(maze[i]))
		copy(mazeCopy[i], maze[i])
	}

	// Directions for adjacent cells
	dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	next := make([]Point, 0)
	for _, end := range ends {
		mazeCopy[end.x][end.y] = p

		for _, dir := range dirs {
			nx, ny := end.x+dir.x, end.y+dir.y

			// Check if we reached target
			if nx == to.x && ny == to.y {
				return p
			}

			// Check bounds and if cell is free
			if nx < 0 || ny < 0 || nx >= len(maze) || ny >= len(maze[0]) || mazeCopy[nx][ny] != 0 {
				continue
			}

			// Add to next points to check
			next = append(next, Point{nx, ny})
			mazeCopy[nx][ny] = p // Mark as visited
		}
	}

	return step(mazeCopy, next, p+1, to)
}

// findPath finds shortest path between two points
func findPath(maze [][]int, from, to Point) int {
	return step(maze, []Point{from}, 1, to)
}

// permutations generates all permutations of numbers from 1 to n-1
func permutations(n int) [][]int {
	nums := make([]int, n-1)
	for i := range nums {
		nums[i] = i + 1
	}

	var result [][]int
	var generate func([]int, int)
	generate = func(arr []int, n int) {
		if n == 1 {
			temp := make([]int, len(arr))
			copy(temp, arr)
			result = append(result, temp)
			return
		}

		for i := 0; i < n; i++ {
			generate(arr, n-1)
			if n%2 == 1 {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			} else {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			}
		}
	}

	generate(nums, len(nums))
	return result
}

// minSum finds minimum sum of paths
func minSum(n int, distances map[string]int, andBack bool) int {
	minSum := 999999
	perms := permutations(n)

	for _, perm := range perms {
		sum := distances[fmt.Sprintf("0,%d", perm[0])]

		for i := 0; i < len(perm)-1; i++ {
			sum += distances[fmt.Sprintf("%d,%d", perm[i], perm[i+1])]
		}

		if andBack {
			sum += distances[fmt.Sprintf("%d,0", perm[len(perm)-1])]
		}

		if sum < minSum {
			minSum = sum
		}
	}

	return minSum
}

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	maze, points := createMaze(input)

	// Calculate distances between all points
	distances := make(map[string]int)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			dist := findPath(maze, points[i], points[j])
			distances[fmt.Sprintf("%d,%d", i, j)] = dist
			distances[fmt.Sprintf("%d,%d", j, i)] = dist
		}
	}

	fmt.Println("Part 1: Given your actual map, and starting from location 0,")
	fmt.Println("what is the fewest number of steps required to visit every non-0 number marked on the map at least once?")
	fmt.Println(minSum(len(points), distances, false))

	fmt.Println()
	fmt.Println("Part 2: What is the fewest number of steps required to start at 0,")
	fmt.Println("visit every non-0 number marked on the map at least once, and then return to 0?")
	fmt.Println(minSum(len(points), distances, true))
}
