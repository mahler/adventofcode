package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Node represents a celestial object in the orbit map
type Node struct {
	code     string
	root     bool
	children []*Node
}

// NewNode creates a new Node with the given code
func NewNode(code string) *Node {
	return &Node{
		code:     code,
		root:     true,
		children: make([]*Node, 0),
	}
}

// buildOrbitGraph constructs the graph from the orbit map input
func buildOrbitGraph(orbitMap []string) map[string]*Node {
	orbitGraph := make(map[string]*Node)

	for _, orbit := range orbitMap {
		parts := strings.Split(orbit, ")")
		left, right := parts[0], parts[1]

		if _, exists := orbitGraph[left]; !exists {
			orbitGraph[left] = NewNode(left)
		}
		if _, exists := orbitGraph[right]; !exists {
			orbitGraph[right] = NewNode(right)
		}

		orbitGraph[right].root = false
		orbitGraph[left].children = append(orbitGraph[left].children, orbitGraph[right])
		orbitGraph[right].children = append(orbitGraph[right].children, orbitGraph[left])
	}

	return orbitGraph
}

// countChecksums calculates the total number of direct and indirect orbits
func countChecksums(graph map[string]*Node) int {
	seen := make(map[string]bool)

	var traverse func(*Node, int) int
	traverse = func(node *Node, depth int) int {
		if seen[node.code] {
			return 0
		}
		seen[node.code] = true

		sum := depth
		for _, child := range node.children {
			sum += traverse(child, depth+1)
		}
		return sum
	}

	total := 0
	for _, node := range graph {
		if node.root {
			// Reset seen map for each root node
			seen = make(map[string]bool)
			total += traverse(node, 0)
		}
	}
	return total
}

// minimumTransfers finds the minimum number of orbital transfers required
func minimumTransfers(graph map[string]*Node, source, destination string) int {
	type QueueItem struct {
		node  *Node
		depth int
	}

	seen := make(map[string]int)
	queue := []QueueItem{{graph[source], 0}}
	seen[source] = 0

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, child := range current.node.children {
			if _, exists := seen[child.code]; !exists {
				seen[child.code] = current.depth + 1
				queue = append([]QueueItem{{child, current.depth + 1}}, queue...)
			}
		}

		if _, exists := seen[destination]; exists {
			break
		}
	}

	return seen[destination] - 2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var orbitMap []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		orbitMap = append(orbitMap, scanner.Text())
	}

	orbitGraph := buildOrbitGraph(orbitMap)
	part1 := countChecksums(orbitGraph)

	fmt.Println("Part 1: What is the total number of direct and indirect orbits in your map data?")
	fmt.Println(part1)

	part2 := minimumTransfers(orbitGraph, "YOU", "SAN")
	fmt.Println()
	fmt.Println("Part 2: What is the minimum number of orbital transfers required to move from the object YOU are orbiting to the object SAN is orbiting?")
	fmt.Println(part2)
}
