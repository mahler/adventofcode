package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Gate represents a logic gate configuration
type Gate struct {
	wireA, wireB, gateType, outputWire string
}

func readInputFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines, scanner.Err()
}

func findGate(xWire, yWire, gateType string, configurations []string) string {
	subStrA := fmt.Sprintf("%s %s %s -> ", xWire, gateType, yWire)
	subStrB := fmt.Sprintf("%s %s %s -> ", yWire, gateType, xWire)

	for _, config := range configurations {
		if strings.Contains(config, subStrA) || strings.Contains(config, subStrB) {
			parts := strings.Split(config, " -> ")
			return parts[len(parts)-1]
		}
	}
	return ""
}

func swapOutputWires(wireA, wireB string, configurations []string) []string {
	newConfigurations := make([]string, 0, len(configurations))

	for _, config := range configurations {
		parts := strings.Split(config, " -> ")
		inputWires, outputWire := parts[0], parts[1]

		switch outputWire {
		case wireA:
			newConfigurations = append(newConfigurations, inputWires+" -> "+wireB)
		case wireB:
			newConfigurations = append(newConfigurations, inputWires+" -> "+wireA)
		default:
			newConfigurations = append(newConfigurations, config)
		}
	}
	return newConfigurations
}

func checkParallelAdders(configurations []string) []string {
	var currentCarryWire string
	var swaps []string

	for bit := 0; bit < 45; bit++ {
		xWire := fmt.Sprintf("x%02d", bit)
		yWire := fmt.Sprintf("y%02d", bit)
		zWire := fmt.Sprintf("z%02d", bit)

		if bit == 0 {
			currentCarryWire = findGate(xWire, yWire, "AND", configurations)
			continue
		}

		abXorGate := findGate(xWire, yWire, "XOR", configurations)
		abAndGate := findGate(xWire, yWire, "AND", configurations)
		cinAbXorGate := findGate(abXorGate, currentCarryWire, "XOR", configurations)

		if cinAbXorGate == "" {
			swaps = append(swaps, abXorGate, abAndGate)
			configurations = swapOutputWires(abXorGate, abAndGate, configurations)
			bit = -1 // Will become 0 after increment
			continue
		}

		if cinAbXorGate != zWire {
			swaps = append(swaps, cinAbXorGate, zWire)
			configurations = swapOutputWires(cinAbXorGate, zWire, configurations)
			bit = -1 // Will become 0 after increment
			continue
		}

		cinAbAndGate := findGate(abXorGate, currentCarryWire, "AND", configurations)
		carryWire := findGate(abAndGate, cinAbAndGate, "OR", configurations)
		currentCarryWire = carryWire
	}

	return swaps
}

func main() {
	lines, err := readInputFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	//	part1solution(lines)
	// Find divider index
	dividerIndex := -1
	for i, line := range lines {
		if line == "" {
			dividerIndex = i
			break
		}
	}

	initialWires := lines[:dividerIndex]
	configurations := lines[dividerIndex+1:]

	// Initialize maps and sets
	wiresMap := make(map[string]bool)
	unprocessedGates := make(map[Gate]struct{})
	readyGates := make(map[Gate]struct{})
	zOutputs := make(map[string]bool)

	// Process initial wires
	for _, wires := range initialWires {
		parts := strings.Split(wires, ":")
		wireName := strings.TrimSpace(parts[0])
		wireValue, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		wiresMap[wireName] = wireValue != 0
	}

	// Process configurations
	for _, gates := range configurations {
		parts := strings.Split(gates, " -> ")
		inputConfig := strings.Split(parts[0], " ")
		wireA := inputConfig[0]
		gateType := inputConfig[1]
		wireB := inputConfig[2]
		outputWire := parts[1]

		gate := Gate{wireA, wireB, gateType, outputWire}

		if _, hasA := wiresMap[wireA]; hasA {
			if _, hasB := wiresMap[wireB]; hasB {
				readyGates[gate] = struct{}{}
				continue
			}
		}
		unprocessedGates[gate] = struct{}{}
	}

	// Process gates
	for {
		// Process ready gates
		for gate := range readyGates {
			delete(readyGates, gate)

			var outputWireValue bool
			switch gate.gateType {
			case "AND":
				outputWireValue = wiresMap[gate.wireA] && wiresMap[gate.wireB]
			case "OR":
				outputWireValue = wiresMap[gate.wireA] || wiresMap[gate.wireB]
			case "XOR":
				outputWireValue = wiresMap[gate.wireA] != wiresMap[gate.wireB]
			}

			wiresMap[gate.outputWire] = outputWireValue
			if strings.HasPrefix(gate.outputWire, "z") {
				zOutputs[gate.outputWire] = outputWireValue
			}
		}

		if len(unprocessedGates) == 0 {
			break
		}

		// Check unprocessed gates
		for gate := range unprocessedGates {
			_, hasA := wiresMap[gate.wireA]
			_, hasB := wiresMap[gate.wireB]
			if hasA && hasB {
				readyGates[gate] = struct{}{}
				delete(unprocessedGates, gate)
			}
		}
	}

	// Sort z outputs and build binary number
	var zKeys []string
	for k := range zOutputs {
		zKeys = append(zKeys, k)
	}
	sort.Strings(zKeys)

	var binaryNum string
	for i := len(zKeys) - 1; i >= 0; i-- {
		if zOutputs[zKeys[i]] {
			binaryNum += "1"
		} else {
			binaryNum += "0"
		}
	}

	// Convert binary to decimal
	result, _ := strconv.ParseInt(binaryNum, 2, 64)

	fmt.Println("Part 1: What decimal number does it output on the wires starting with z?")
	fmt.Println(result)

	// Part 2...
	// p2solution(lines)
	dividerIndex = 0 // reset var.
	for i, line := range lines {
		if line == "" {
			dividerIndex = i
			break
		}
	}

	p2configurations := lines[dividerIndex+1:]
	swaps := checkParallelAdders(p2configurations)

	// Sort and remove duplicates
	sort.Strings(swaps)
	uniqueSwaps := make([]string, 0, len(swaps))
	seen := make(map[string]bool)

	for _, swap := range swaps {
		if !seen[swap] {
			seen[swap] = true
			uniqueSwaps = append(uniqueSwaps, swap)
		}
	}

	fmt.Println()
	fmt.Println("Part 2: what do you get if you sort the names of the eight wires involved")
	fmt.Println("in a swap and then join those names with commas?")
	fmt.Println(strings.Join(uniqueSwaps, ","))
}
