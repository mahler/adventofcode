package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph struct {
	next  map[string]map[string]bool
	cache map[string]int
}

func NewGraph() *Graph {
	return &Graph{
		next:  make(map[string]map[string]bool),
		cache: make(map[string]int),
	}
}

func (g *Graph) AddEdges(u string, vs []string) {
	if g.next[u] == nil {
		g.next[u] = make(map[string]bool)
	}
	for _, v := range vs {
		g.next[u][v] = true
	}
}

func (g *Graph) PathCount(u, dest string) int {
	// Create cache key
	key := u + ":" + dest

	// Check cache
	if count, ok := g.cache[key]; ok {
		return count
	}

	// Base case
	if u == dest {
		g.cache[key] = 1
		return 1
	}

	// Recursive case: sum paths through neighbors
	total := 0
	for v := range g.next[u] {
		total += g.PathCount(v, dest)
	}

	g.cache[key] = total
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	graph := NewGraph()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		u := strings.TrimSpace(parts[0])
		vs := strings.Fields(parts[1])

		graph.AddEdges(u, vs)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	// Part 1
	p1 := graph.PathCount("you", "out")

	// Part 2
	m1 := graph.PathCount("fft", "dac")
	m2 := graph.PathCount("dac", "fft")

	var p2 int
	if m1 > 0 {
		p2 = graph.PathCount("svr", "fft") * m1 * graph.PathCount("dac", "out")
	} else {
		p2 = graph.PathCount("svr", "dac") * m2 * graph.PathCount("fft", "out")
	}

	fmt.Println(p1, p2)
}
