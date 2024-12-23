package main

import (
	_ "embed"
	"fmt"
	"slices"
)

type Input struct {
	nodes map[int][]int
	edges [676][676]bool
}

//go:embed input.txt
var input string

func main() {
	nodes := make(map[int][]int, 1000)
	edges := [676][676]bool{}

	bytes := []byte(input)
	for i := 0; i < len(bytes); i += 6 {
		from := 26*int(bytes[i]-'a') + int(bytes[i+1]-'a')
		to := 26*int(bytes[i+3]-'a') + int(bytes[i+4]-'a')

		if nodes[from] == nil {
			nodes[from] = make([]int, 0, 16)
		}
		if nodes[to] == nil {
			nodes[to] = make([]int, 0, 16)
		}

		nodes[from] = append(nodes[from], to)
		nodes[to] = append(nodes[to], from)
		edges[from][to] = true
		edges[to][from] = true
	}

	parsedData := Input{nodes, edges}

	// Part 1
	seen := [676]bool{}
	triangles := 0

	for n1 := 494; n1 < 520; n1++ {
		if neighbours, ok := parsedData.nodes[n1]; ok {
			seen[n1] = true
			for i, n2 := range neighbours {
				for _, n3 := range neighbours[i:] {
					if !seen[n2] && !seen[n3] && parsedData.edges[n2][n3] {
						triangles++
					}
				}
			}
		}
	}

	fmt.Println("Part 1: How many contain at least one computer with a name that starts with t?")
	fmt.Println(triangles)

	// Part 2
	seen = [676]bool{}
	clique := make([]int, 0, 676)
	largest := make([]int, 0, 676)

	for n1, neighbours := range parsedData.nodes {
		if !seen[n1] {
			clique = clique[:0]
			clique = append(clique, n1)

			for _, n2 := range neighbours {
				allConnected := true
				for _, c := range clique {
					if !parsedData.edges[n2][c] {
						allConnected = false
						break
					}
				}
				if allConnected {
					seen[n2] = true
					clique = append(clique, n2)
				}
			}

			if len(clique) > len(largest) {
				largest = largest[:0]
				largest = append(largest, clique...)
			}
		}
	}

	slices.Sort(largest)

	var result []byte
	for _, n := range largest {
		result = append(result, byte(n/26)+'a', byte(n%26)+'a', ',')
	}

	fmt.Println()
	fmt.Println("Part 2: What is the password to get into the LAN party?")
	fmt.Println(string(result[:len(result)-1]))

}
