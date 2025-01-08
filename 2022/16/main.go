package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Valve struct {
	name        string
	rate        uint32
	connections []string
}

type SimpleValve struct {
	name  string
	rate  uint32
	links []int
}

func parseValve(s string) (Valve, error) {
	parts := strings.Split(strings.TrimSpace(s), ";")
	if len(parts) != 2 {
		return Valve{}, fmt.Errorf("invalid valve format")
	}

	// Parse name and rate
	firstPart := strings.TrimSpace(parts[0])
	if !strings.HasPrefix(firstPart, "Valve ") {
		return Valve{}, fmt.Errorf("invalid valve prefix")
	}
	nameAndRate := strings.Split(firstPart[6:], " has flow rate=")
	if len(nameAndRate) != 2 {
		return Valve{}, fmt.Errorf("invalid name/rate format")
	}

	name := nameAndRate[0]
	var rate uint32
	_, err := fmt.Sscanf(nameAndRate[1], "%d", &rate)
	if err != nil {
		return Valve{}, fmt.Errorf("invalid rate: %v", err)
	}

	// Parse connections
	secondPart := strings.TrimSpace(parts[1])
	secondPart = strings.TrimPrefix(secondPart, "tunnel leads to valve ")
	secondPart = strings.TrimPrefix(secondPart, "tunnels lead to valves ")
	connections := strings.Split(secondPart, ", ")

	return Valve{
		name:        name,
		rate:        rate,
		connections: connections,
	}, nil
}

func parse(input string) []SimpleValve {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	valves := make([]Valve, 0)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		valve, err := parseValve(line)
		if err != nil {
			panic(err)
		}
		valves = append(valves, valve)
	}

	// Create index map
	idxMap := make(map[string]int)
	for i, v := range valves {
		idxMap[v.name] = i
	}

	// Convert to SimpleValve
	simpleValves := make([]SimpleValve, len(valves))
	for i, v := range valves {
		links := make([]int, len(v.connections))
		for j, conn := range v.connections {
			links[j] = idxMap[conn]
		}
		simpleValves[i] = SimpleValve{
			name:  v.name,
			rate:  v.rate,
			links: links,
		}
	}

	return simpleValves
}

func initGraph(valves []SimpleValve) [][]uint32 {
	l := len(valves)
	graph := make([][]uint32, l)
	for i := range graph {
		graph[i] = make([]uint32, l)
		for j := range graph[i] {
			graph[i][j] = math.MaxUint32 / 4
		}
	}

	for i, v := range valves {
		for _, j := range v.links {
			graph[i][j] = 1
		}
	}

	return graph
}

func floydWarshall(graph [][]uint32) [][]uint32 {
	l := len(graph)
	dist := make([][]uint32, l)
	for i := range dist {
		dist[i] = make([]uint32, l)
		copy(dist[i], graph[i])
	}

	for k := 0; k < l; k++ {
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func travelingSalesman(
	valves []SimpleValve,
	memo map[uint64]uint32,
	nonZeroValves []int,
	dist [][]uint32,
	mask uint64,
	minutes uint32,
	flow uint32,
	i int,
	depth uint32,
) uint32 {
	maxFlow := flow

	if current, exists := memo[mask]; !exists || flow > current {
		memo[mask] = flow
	}

	for _, j := range nonZeroValves {
		var curMinutes uint32
		if dist[i][j]+1 <= minutes {
			curMinutes = minutes - dist[i][j] - 1
		}

		if (mask&(1<<j)) == 0 || curMinutes <= 0 {
			continue
		}

		curMask := mask &^ (1 << j)
		curFlow := flow + (curMinutes * valves[j].rate)

		newFlow := travelingSalesman(valves, memo, nonZeroValves, dist, curMask, curMinutes, curFlow, j, depth+1)
		if newFlow > maxFlow {
			maxFlow = newFlow
		}
	}

	return maxFlow
}

func simulate(valves []SimpleValve, dist [][]uint32, initMask uint64, startIdx int, minutes uint32) (uint32, map[uint64]uint32) {
	nonZeroValves := make([]int, 0)
	for i, v := range valves {
		if v.rate > 0 {
			nonZeroValves = append(nonZeroValves, i)
		}
	}

	maskFlow := make(map[uint64]uint32)
	flow := travelingSalesman(valves, maskFlow, nonZeroValves, dist, initMask, minutes, 0, startIdx, 0)

	return flow, maskFlow
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("File not found: %v", err))
	}

	valves := parse(string(contents))

	graph := initGraph(valves)
	dist := floydWarshall(graph)

	var startIdx int
	for i, v := range valves {
		if v.name == "AA" {
			startIdx = i
			break
		}
	}

	initMask := uint64(1<<len(dist)) - 1
	flow, _ := simulate(valves, dist, initMask, startIdx, 30)

	fmt.Println("part 1: What is the most pressure you can release?")
	fmt.Println(flow)

	// Part 2
	for i, v := range valves {
		if v.name == "AA" {
			startIdx = i
			break
		}
	}

	_, elfMemo := simulate(valves, dist, initMask, startIdx, 26)
	_, elephantMemo := simulate(valves, dist, initMask, startIdx, 26)

	maxFlow := uint32(0)
	for elfMask, elfFlow := range elfMemo {
		for mask, elephantFlow := range elephantMemo {
			if (^mask)&(^elfMask)&initMask == 0 {
				if sum := elephantFlow + elfFlow; sum > maxFlow {
					maxFlow = sum
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("Part 2: With you and an elephant working together for 26 minutes, what is the most pressure you could release?")
	fmt.Println(maxFlow)
}
