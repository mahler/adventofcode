package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// parse takes an input string and returns two maps:
// 1. map of node names to weights
// 2. map of node names to their successor nodes
func parse(input string) (map[string]int, map[string][]string) {
	weights := make(map[string]int)
	succs := make(map[string][]string)

	// Remove all non-alphanumeric characters except whitespace
	reg := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	cleanInput := reg.ReplaceAllString(input, "")

	scanner := bufio.NewScanner(strings.NewReader(cleanInput))
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) < 2 {
			continue
		}

		name := words[0]
		weight, _ := strconv.Atoi(words[1])
		weights[name] = weight

		// Collect successors (remaining words after name and weight)
		successors := words[2:]
		succs[name] = successors
	}

	return weights, succs
}

// topologicalSort performs a topological sort on the graph
func topologicalSort(nodes []string, getSuccs func(string) []string) ([]string, error) {
	result := make([]string, 0, len(nodes))
	visited := make(map[string]bool)
	temp := make(map[string]bool)

	var visit func(string) error
	visit = func(node string) error {
		if temp[node] {
			return fmt.Errorf("cycle detected")
		}
		if visited[node] {
			return nil
		}

		temp[node] = true
		for _, succ := range getSuccs(node) {
			if err := visit(succ); err != nil {
				return err
			}
		}
		temp[node] = false
		visited[node] = true
		result = append([]string{node}, result...)
		return nil
	}

	for _, node := range nodes {
		if !visited[node] {
			if err := visit(node); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

// allEqual checks if all elements in a slice are equal
func allEqual(slice []int) bool {
	if len(slice) == 0 {
		return true
	}
	first := slice[0]
	for _, value := range slice[1:] {
		if value != first {
			return false
		}
	}
	return true
}

func main() {
	// Read input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	weights, succs := parse(string(input))

	// Get all nodes for topological sort
	nodes := make([]string, 0, len(succs))
	for node := range succs {
		nodes = append(nodes, node)
	}

	// Perform topological sort
	topo, err := topologicalSort(nodes, func(n string) []string {
		return succs[n]
	})
	if err != nil {
		fmt.Printf("Error in topological sort: %v\n", err)
		return
	}

	fmt.Println("Part 1: Before you're ready to help them, you need to make sure")
	fmt.Println("your information is correct. What is the name of the bottom program?")
	fmt.Println(topo[0])

	// Part 2
	// Create a copy of original weights
	oweights := make(map[string]int)
	for k, v := range weights {
		oweights[k] = v
	}

	// Process nodes in reverse topological order
	for i := len(topo) - 1; i >= 0; i-- {
		n := topo[i]
		sn := succs[n]
		ws := make([]int, len(sn))
		for j, s := range sn {
			ws[j] = weights[s]
		}

		if allEqual(ws) {
			sum := 0
			for _, w := range ws {
				sum += w
			}
			weights[n] += sum
		} else {
			var m int
			if len(ws) >= 3 {
				if ws[0] == ws[1] {
					m = ws[0]
				} else {
					m = ws[2]
				}
			}

			// Find the node with different weight
			var c string
			for i, w := range ws {
				if w != m {
					c = sn[i]
					break
				}
			}

			fmt.Println()
			fmt.Println("Part 2: Given that exactly one program is the wrong weight,")
			fmt.Println("what would its weight need to be to balance the entire tower?")
			fmt.Println(oweights[c] + m - weights[c])
			break
		}
	}
}
