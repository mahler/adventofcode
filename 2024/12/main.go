package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}

func parseInput(scanner *bufio.Scanner) [][]rune {
	var input [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []rune(line))
	}
	return input
}

func getNeighbors(pos [2]int, grid [][]rune, plant rune) [][2]int {
	var neighbors [][2]int
	for _, dir := range directions {
		neighbor := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if neighbor[0] >= 0 && neighbor[0] < len(grid) &&
			neighbor[1] >= 0 && neighbor[1] < len(grid[0]) &&
			grid[neighbor[0]][neighbor[1]] == plant {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func bfs(pos [2]int, visited map[[2]int]bool, grid [][]rune) (map[[2]int]bool, int) {
	plant := grid[pos[0]][pos[1]]
	area := make(map[[2]int]bool)
	perimeter := 0
	queue := [][2]int{pos}
	visited[pos] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		area[curr] = true

		neighbors := getNeighbors(curr, grid, plant)
		perimeter += 4 - len(neighbors)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return area, perimeter
}

func countRegionSides(region map[[2]int]bool) int {
	sideCount := 0
	for _, dir := range directions {
		sides := make(map[[2]int]bool)
		for pos := range region {
			neighbor := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			if !region[neighbor] {
				sides[neighbor] = true
			}
		}

		remove := make(map[[2]int]bool)
		for side := range sides {
			next := [2]int{side[0] + dir[1], side[1] + dir[0]}
			for sides[next] {
				remove[next] = true
				next = [2]int{next[0] + dir[1], next[1] + dir[0]}
			}
		}
		sideCount += len(sides) - len(remove)
	}
	return sideCount
}

func part1(grid [][]rune) int {
	visited := make(map[[2]int]bool)
	price := 0

	for i := range grid {
		for j := range grid[i] {
			if !visited[[2]int{i, j}] {
				regionArea, regionPerimeter := bfs([2]int{i, j}, visited, grid)
				price += len(regionArea) * regionPerimeter
			}
		}
	}
	return price
}

func part2(grid [][]rune) int {
	visited := make(map[[2]int]bool)
	price := 0

	for i := range grid {
		for j := range grid[i] {
			if !visited[[2]int{i, j}] {
				regionArea, _ := bfs([2]int{i, j}, visited, grid)
				sides := countRegionSides(regionArea)
				price += len(regionArea) * sides
			}
		}
	}
	return price
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	inputScanner := bufio.NewScanner(inputFile)
	inputGrid := parseInput(inputScanner)

	result1 := part1(inputGrid)
	fmt.Println("Part 1: What is the total price of fencing all regions on your map?")
	fmt.Println(result1)

	result2 := part2(inputGrid)
	fmt.Println()
	fmt.Println("Part 2: What is the new total price of fencing all regions on your map?")
	fmt.Println(result2)
}
